package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"

	"coffee-chat-service/modules/entity"
	interfaces "coffee-chat-service/modules/interface"
)

type WebSocketMessage struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type DirectMessage struct {
	SenderID uint
	Message  []byte
}

type MessagePayload struct {
	RecipientID      *uint  `json:"recipient_id,omitempty"`
	GroupID          *uint  `json:"group_id,omitempty"`
	Text             string `json:"text"`
	ReplyToMessageID *uint  `json:"reply_to_message_id,omitempty"`
	MenuID           *uint  `json:"menu_id,omitempty"`
	OrderID          *uint  `json:"order_id,omitempty"`
}

type MenuInfo struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	ImageURL string  `json:"image_url"`
}

type OrderItemInfo struct {
	ID       uint      `json:"id"`
	MenuID   uint      `json:"menu_id"`
	Quantity int       `json:"quantity"`
	Price    float64   `json:"price"`
	Menu     *MenuInfo `json:"menu,omitempty"`
}

type OrderInfo struct {
	ID               uint            `json:"id"`
	CustomerID       uint            `json:"customer_id"`
	RecipientID      *uint           `json:"recipient_id,omitempty"`
	TableID          uint            `json:"table_id"`
	TableNumber      string          `json:"table_number"`
	TableName        string          `json:"table_name"`
	TableFloorNumber int             `json:"table_floor_number"`
	NeedType         string          `json:"need_type"`
	Notes            string          `json:"notes,omitempty"`
	SubTotal         float64         `json:"sub_total"`
	Tax              float64         `json:"tax"`
	Total            float64         `json:"total"`
	OrderItems       []OrderItemInfo `json:"order_items"`
}

type RepliedMessageInfo struct {
	ID         uint       `json:"id"`
	Text       string     `json:"text"`
	SenderName string     `json:"sender_name"`
	Menu       *MenuInfo  `json:"menu,omitempty"`
	Order      *OrderInfo `json:"order,omitempty"`
}

type IncomingMessagePayload struct {
	MessageID         uint                `json:"message_id"`
	SenderID          uint                `json:"sender_id"`
	SenderName        string              `json:"sender_name"`
	SenderPhotoURL    string              `json:"sender_photo_url"`
	SenderTableNumber string              `json:"sender_table_number"`
	SenderFloorNumber int                 `json:"sender_floor_number"`
	Text              string              `json:"text"`
	Timestamp         time.Time           `json:"timestamp"`
	ReplyTo           *RepliedMessageInfo `json:"reply_to,omitempty"`
	Menu              *MenuInfo           `json:"menu,omitempty"`
	Order             *OrderInfo          `json:"order,omitempty"`
}

type IncomingGroupMessagePayload struct {
	MessageID         uint                `json:"message_id"`
	GroupID           uint                `json:"group_id"`
	SenderID          uint                `json:"sender_id"`
	SenderName        string              `json:"sender_name"`
	SenderPhotoURL    string              `json:"sender_photo_url"`
	SenderTableNumber string              `json:"sender_table_number"`
	SenderFloorNumber int                 `json:"sender_floor_number"`
	Text              string              `json:"text"`
	Timestamp         time.Time           `json:"timestamp"`
	ReplyTo           *RepliedMessageInfo `json:"reply_to,omitempty"`
	Menu              *MenuInfo           `json:"menu,omitempty"`
	Order             *OrderInfo          `json:"order,omitempty"`
}
type Hub struct {
	clients         map[*Client]bool
	customers       map[uint]*Client
	admins          map[uint]*Client
	incoming        chan *DirectMessage
	Broadcast       chan []byte
	register        chan *Client
	unregister      chan *Client
	DB              *gorm.DB
	BroadcastAdmins chan []byte
	BlockRepo       interfaces.BlockRepositoryInterface
	GroupRepo       interfaces.GroupRepositoryInterface
}

func NewHub(db *gorm.DB, blockRepo interfaces.BlockRepositoryInterface,
	groupRepo interfaces.GroupRepositoryInterface) *Hub {
	return &Hub{
		clients:         make(map[*Client]bool),
		customers:       make(map[uint]*Client),
		admins:          make(map[uint]*Client),
		incoming:        make(chan *DirectMessage),
		Broadcast:       make(chan []byte),
		register:        make(chan *Client),
		unregister:      make(chan *Client),
		BroadcastAdmins: make(chan []byte),
		DB:              db,
		BlockRepo:       blockRepo,
		GroupRepo:       groupRepo,
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			if client.CustomerID != 0 {
				h.customers[client.CustomerID] = client
				log.Printf("Customer connected: ID %d", client.CustomerID)
			}
			if client.AdminID != 0 {
				h.admins[client.AdminID] = client
				log.Printf("Admin connected: ID %d", client.AdminID)
			}

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)

				if client.CustomerID != 0 {
					delete(h.customers, client.CustomerID)
					log.Printf("Customer disconnected: ID %d", client.CustomerID)
				}
				if client.AdminID != 0 {
					delete(h.admins, client.AdminID)
					log.Printf("Admin disconnected: ID %d", client.AdminID)
				}
			}

		case directMsg := <-h.incoming:
			var payload MessagePayload
			if err := json.Unmarshal(directMsg.Message, &payload); err != nil {
				log.Printf("Error unmarshalling message: %v", err)
				continue
			}

			if payload.GroupID != nil {
				h.handleGroupMessage(directMsg.SenderID, payload)
			} else if payload.RecipientID != nil {
				h.handleDirectMessage(directMsg.SenderID, *payload.RecipientID, payload)
			} else {
				log.Printf("Invalid message payload: missing group_id and recipient_id")
			}

		case message := <-h.Broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
					delete(h.customers, client.CustomerID)
				}
			}
		case message := <-h.BroadcastAdmins:
			for _, adminClient := range h.admins {
				select {
				case adminClient.send <- message:
				default:
					close(adminClient.send)
					delete(h.clients, adminClient)
					delete(h.admins, adminClient.AdminID)
					log.Printf("Admin channel full or closed. Disconnecting Admin ID %d", adminClient.AdminID)
				}
			}

		}
	}
}

func (h *Hub) SendChatMessage(chatMessage *entity.ChatMessage) {
	if chatMessage == nil {
		return
	}

	isBlocked, err := h.BlockRepo.IsBlocked(chatMessage.SenderID, chatMessage.RecipientID)
	if err != nil {
		log.Printf("Error checking block status for chat message %d: %v", chatMessage.ID, err)
		return
	}
	if isBlocked {
		log.Printf("Blocked chat notification from %d to %d", chatMessage.SenderID, chatMessage.RecipientID)
		return
	}

	var sender entity.Customer
	if err := h.DB.Preload("Table.Floor").First(&sender, chatMessage.SenderID).Error; err != nil {
		log.Printf("failed to load sender for chat message %d: %v", chatMessage.ID, err)
		return
	}

	var repliedToInfo *RepliedMessageInfo
	if chatMessage.ReplyToMessageID != nil {
		var originalMsg entity.ChatMessage
		if err := h.DB.Preload("Sender").First(&originalMsg, *chatMessage.ReplyToMessageID).Error; err == nil {
			repliedToInfo = &RepliedMessageInfo{
				ID:         originalMsg.ID,
				Text:       originalMsg.Text,
				SenderName: originalMsg.Sender.Name,
			}
		}
	}

	var menuInfo *MenuInfo
	if chatMessage.MenuID != nil {
		var menu entity.Menu
		if err := h.DB.First(&menu, *chatMessage.MenuID).Error; err == nil {
			menuInfo = &MenuInfo{
				ID:       menu.ID,
				Name:     menu.Name,
				Price:    menu.Price,
				ImageURL: menu.ImageURL,
			}
		}
	}

	var orderInfo *OrderInfo
	if chatMessage.OrderID != nil {
		if order, err := h.loadOrder(*chatMessage.OrderID); err == nil {
			orderInfo = buildOrderInfo(order)
		} else {
			log.Printf("failed to load order for chat message %d: %v", chatMessage.ID, err)
		}
	}

	payload, err := h.buildIncomingPayload(chatMessage, &sender, repliedToInfo, menuInfo, orderInfo)
	if err != nil {
		return
	}

	if recipient, ok := h.customers[chatMessage.RecipientID]; ok {
		recipient.send <- payload
	}
}

func (h *Hub) buildIncomingPayload(chatMessage *entity.ChatMessage, sender *entity.Customer, repliedToInfo *RepliedMessageInfo, menuInfo *MenuInfo, orderInfo *OrderInfo) ([]byte, error) {
	if chatMessage == nil || sender == nil {
		return nil, fmt.Errorf("invalid chat payload")
	}

	responsePayload := IncomingMessagePayload{
		MessageID:         chatMessage.ID,
		SenderID:          chatMessage.SenderID, // FIX: Use the actual sender ID from the message
		SenderName:        sender.Name,
		SenderPhotoURL:    sender.PhotoURL,
		SenderTableNumber: sender.Table.TableNumber,
		SenderFloorNumber: sender.Table.Floor.FloorNumber,
		Text:              chatMessage.Text,
		Timestamp:         chatMessage.CreatedAt,
		ReplyTo:           repliedToInfo,
		Menu:              menuInfo,
		Order:             orderInfo,
	}

	wrapper := WebSocketMessage{
		Type: "CHAT_MESSAGE",
		Data: responsePayload,
	}

	responseJSON, err := json.Marshal(wrapper)
	if err != nil {
		return nil, err
	}
	return responseJSON, nil
}

func (h *Hub) loadOrder(orderID uint) (*entity.Order, error) {
	var order entity.Order
	if err := h.DB.Preload("Table.Floor").Preload("OrderItems.Menu").First(&order, orderID).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func buildOrderInfo(order *entity.Order) *OrderInfo {
	if order == nil {
		return nil
	}

	orderInfo := &OrderInfo{
		ID:         order.ID,
		CustomerID: order.CustomerID,
		TableID:    order.TableID,
		NeedType:   order.NeedType,
		Notes:      order.Notes,
		SubTotal:   order.SubTotal,
		Tax:        order.Tax,
		Total:      order.Total,
	}

	if order.RecipientID != nil {
		orderInfo.RecipientID = order.RecipientID
	}

	if order.Table.ID != 0 {
		orderInfo.TableNumber = order.Table.TableNumber
		orderInfo.TableName = order.Table.TableName
		if order.Table.Floor.ID != 0 {
			orderInfo.TableFloorNumber = order.Table.Floor.FloorNumber
		}
	}

	if len(order.OrderItems) > 0 {
		items := make([]OrderItemInfo, 0, len(order.OrderItems))
		for _, item := range order.OrderItems {
			orderItemInfo := OrderItemInfo{
				ID:       item.ID,
				MenuID:   item.MenuID,
				Quantity: item.Quantity,
				Price:    item.Price,
			}

			if item.Menu.ID != 0 {
				orderItemInfo.Menu = &MenuInfo{
					ID:       item.Menu.ID,
					Name:     item.Menu.Name,
					Price:    item.Menu.Price,
					ImageURL: item.Menu.ImageURL,
				}
			}

			items = append(items, orderItemInfo)
		}
		orderInfo.OrderItems = items
	}

	return orderInfo
}

func (h *Hub) BroadcastToCustomers(message []byte) {
	for _, customerClient := range h.customers {
		select {
		case customerClient.send <- message:
		default:
			close(customerClient.send)
			delete(h.clients, customerClient)
			delete(h.customers, customerClient.CustomerID)
		}
	}
}

func (h *Hub) BroadcastToAdmins(message []byte) {
	for _, adminClient := range h.admins {
		select {
		case adminClient.send <- message:
		default:
			close(adminClient.send)
			delete(h.clients, adminClient)
			delete(h.admins, adminClient.AdminID)
			log.Printf("Admin channel full or closed. Disconnecting Admin ID %d", adminClient.AdminID)
		}
	}
}

func (h *Hub) handleDirectMessage(senderID, recipientID uint, payload MessagePayload) {
	isBlocked, err := h.BlockRepo.IsBlocked(senderID, recipientID)
	if err != nil {
		log.Printf("Error checking block status: %v", err)
		return
	}
	if isBlocked {
		log.Printf("Blocked message from %d to %d", senderID, payload.RecipientID)
		return
	}

	chatMessage := entity.ChatMessage{
		SenderID:         senderID,
		RecipientID:      recipientID,
		Text:             payload.Text,
		ReplyToMessageID: payload.ReplyToMessageID,
		MenuID:           payload.MenuID,
		OrderID:          payload.OrderID,
	}
	if err := h.DB.Create(&chatMessage).Error; err != nil {
		log.Printf("Failed to save chat message: %v", err)
		return
	}

	var sender entity.Customer
	if err := h.DB.Preload("Table.Floor").First(&sender, senderID).Error; err != nil {
		return
	}

	var repliedToInfo *RepliedMessageInfo
	if chatMessage.ReplyToMessageID != nil {
		var originalMsg entity.ChatMessage
		err := h.DB.Preload("Sender").
			Preload("Menu").
			Preload("Order.Table.Floor").
			Preload("Order.OrderItems.Menu").
			First(&originalMsg, *chatMessage.ReplyToMessageID).Error
		if err == nil {
			repliedToInfo = &RepliedMessageInfo{
				ID:         originalMsg.ID,
				Text:       originalMsg.Text,
				SenderName: originalMsg.Sender.Name,
			}

			if originalMsg.MenuID != nil && originalMsg.Menu != nil {
				repliedToInfo.Menu = buildMenuInfo(originalMsg.Menu)
			}

			if originalMsg.OrderID != nil && originalMsg.Order != nil {
				repliedToInfo.Order = buildOrderInfo(originalMsg.Order)
			}
		}
	}

	var menuInfo *MenuInfo
	if chatMessage.MenuID != nil {
		var menu entity.Menu
		if err := h.DB.First(&menu, *chatMessage.MenuID).Error; err == nil {
			menuInfo = buildMenuInfo(&menu)
		}
	}

	var orderInfo *OrderInfo
	if chatMessage.OrderID != nil {
		if order, err := h.loadOrder(*chatMessage.OrderID); err == nil {
			orderInfo = buildOrderInfo(order)
		} else {
			log.Printf("failed to load order for chat message %d: %v", chatMessage.ID, err)
		}
	}

	if recipient, ok := h.customers[recipientID]; ok {
		if payload, err := h.buildIncomingPayload(&chatMessage, &sender, repliedToInfo, menuInfo, orderInfo); err == nil {
			recipient.send <- payload
		}
	} else {
		log.Printf("Recipient not found or offline: CustomerID %d", payload.RecipientID)
	}
}

func (h *Hub) handleGroupMessage(senderID uint, payload MessagePayload) {
	groupID := *payload.GroupID

	isMember, err := h.GroupRepo.IsCustomerMember(groupID, senderID)
	if err != nil {
		log.Printf("Error checking group membership: %v", err)
		return
	}
	if !isMember {
		log.Printf("Blocked group message: Sender %d is not member of group %d", senderID, groupID)
		return
	}

	groupMessage := &entity.GroupChatMessage{
		ChatGroupID:      groupID,
		SenderID:         senderID,
		Text:             payload.Text,
		ReplyToMessageID: payload.ReplyToMessageID,
		MenuID:           payload.MenuID,
		OrderID:          payload.OrderID,
	}
	if err := h.GroupRepo.CreateGroupMessage(groupMessage); err != nil {
		log.Printf("Failed to save group message: %v", err)
		return
	}

	var sender entity.Customer
	if err := h.DB.Preload("Table.Floor").First(&sender, senderID).Error; err != nil {
		log.Printf("Failed to load sender for group message: %v", err)
		return
	}

	var repliedToInfo *RepliedMessageInfo
	if groupMessage.ReplyToMessageID != nil {
		var originalMsg entity.GroupChatMessage
		err := h.DB.Preload("Sender").
			Preload("Menu").
			Preload("Order.Table.Floor").
			Preload("Order.OrderItems.Menu").
			First(&originalMsg, *groupMessage.ReplyToMessageID).Error
		if err == nil {
			repliedToInfo = &RepliedMessageInfo{
				ID:         originalMsg.ID,
				Text:       originalMsg.Text,
				SenderName: originalMsg.Sender.Name,
			}
			if originalMsg.MenuID != nil && originalMsg.Menu != nil {
				repliedToInfo.Menu = buildMenuInfo(originalMsg.Menu)
			}
			if originalMsg.OrderID != nil && originalMsg.Order != nil {
				repliedToInfo.Order = buildOrderInfo(originalMsg.Order)
			}
		}
	}

	var menuInfo *MenuInfo
	if groupMessage.MenuID != nil {
		var menu entity.Menu
		if err := h.DB.First(&menu, *groupMessage.MenuID).Error; err == nil {
			menuInfo = buildMenuInfo(&menu)
		}
	}

	var orderInfo *OrderInfo
	if groupMessage.OrderID != nil {
		if order, err := h.loadOrder(*groupMessage.OrderID); err == nil {
			orderInfo = buildOrderInfo(order)
		} else {
			log.Printf("failed to load order for group chat message %d: %v", groupMessage.ID, err)
		}
	}

	responsePayload := IncomingGroupMessagePayload{
		MessageID:         groupMessage.ID,
		GroupID:           groupID,
		SenderID:          groupMessage.SenderID, // FIX: Use the actual sender ID from the message
		SenderName:        sender.Name,
		SenderPhotoURL:    sender.PhotoURL,
		SenderTableNumber: sender.Table.TableNumber,
		SenderFloorNumber: sender.Table.Floor.FloorNumber,
		Text:              groupMessage.Text,
		Timestamp:         groupMessage.CreatedAt,
		ReplyTo:           repliedToInfo,
		Menu:              menuInfo,
		Order:             orderInfo,
	}

	wrapper := WebSocketMessage{
		Type: "CHAT_MESSAGE",
		Data: responsePayload,
	}

	responseJSON, err := json.Marshal(wrapper)
	if err != nil {
		log.Printf("Failed to marshal group message: %v", err)
		return
	}

	h.BroadcastToGroup(groupID, senderID, responseJSON)
}

func (h *Hub) BroadcastToGroup(groupID, senderID uint, message []byte) {
	members, err := h.GroupRepo.GetGroupMembers(groupID)
	if err != nil {
		log.Printf("Failed to get group members for broadcast: %v", err)
		return
	}

	for _, member := range members {
		if member.CustomerID == senderID {
			continue
		}

		if client, ok := h.customers[member.CustomerID]; ok {
			select {
			case client.send <- message:
			default:
				close(client.send)
				delete(h.clients, client)
				delete(h.customers, client.CustomerID)
			}
		}
	}
}

func buildMenuInfo(menu *entity.Menu) *MenuInfo {
	if menu == nil {
		return nil
	}
	return &MenuInfo{
		ID:       menu.ID,
		Name:     menu.Name,
		Price:    menu.Price,
		ImageURL: menu.ImageURL,
	}
}
