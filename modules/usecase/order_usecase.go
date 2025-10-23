package usecase

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"coffee-chat-service/modules/entity"
	interfaces "coffee-chat-service/modules/interface"
	"coffee-chat-service/modules/model"
	"coffee-chat-service/modules/repository"
	"coffee-chat-service/modules/websocket"
)

const TAX_RATE = 0.11 // Pajak 11%

type OrderUseCase struct {
	OrderRepo *repository.OrderRepository
	ChatRepo  interfaces.ChatRepositoryInterface
	Hub       *websocket.Hub
}

func (uc *OrderUseCase) CreateOrder(customerID uint, req model.CreateOrderRequest) (*model.CreateOrderResponse, error) {
	if len(req.OrderItems) == 0 {
		return nil, &model.ValidationError{Message: "order must have at least one item"}
	}
	needType := strings.TrimSpace(req.NeedType)
	if needType == "" {
		return nil, &model.ValidationError{Message: "need_type is required"}
	}

	isWishlist := needType == model.OrderNeedRequestPublic
	if isWishlist {
		req.RecipientCustomerID = nil
	}

	customer, err := uc.OrderRepo.FindCustomerWithTable(customerID)
	if err != nil {
		return nil, &model.ValidationError{Message: "customer not found"}
	}
	if customer.Table.ID == 0 {
		return nil, &model.ValidationError{Message: "customer is not associated with a valid table"}
	}

	var recipient *entity.Customer
	if needType != model.OrderNeedSelf && !isWishlist {
		if req.RecipientCustomerID == nil {
			return nil, &model.ValidationError{Message: "recipient_customer_id is required"}
		}
		recipient, err = uc.OrderRepo.FindCustomerWithTable(*req.RecipientCustomerID)
		if err != nil {
			return nil, &model.ValidationError{Message: "recipient not found"}
		}
	}

	var menuIDs []uint
	for _, item := range req.OrderItems {
		menuIDs = append(menuIDs, item.MenuID)
	}
	menuMap, err := uc.OrderRepo.FindMenusByIDs(menuIDs)
	if err != nil || len(menuMap) != len(menuIDs) {
		return nil, &model.ValidationError{Message: "one or more menu items not found"}
	}

	var subTotal float64
	var orderItems []entity.OrderItem
	var responseItems []model.OrderItemSummary
	for _, item := range req.OrderItems {
		menu := menuMap[item.MenuID]
		lineTotal := menu.Price * float64(item.Quantity)
		subTotal += lineTotal
		orderItems = append(orderItems, entity.OrderItem{MenuID: item.MenuID, Quantity: item.Quantity, Price: menu.Price})
		responseItems = append(responseItems, model.OrderItemSummary{MenuID: menu.ID, MenuName: menu.Name, Quantity: item.Quantity, UnitPrice: menu.Price, TotalPrice: lineTotal})
	}
	tax := subTotal * TAX_RATE
	total := subTotal + tax

	tableID := customer.TableID
	tableNumber := customer.Table.TableNumber
	tableName := customer.Table.TableName
	var recipientSummary *model.OrderRecipient

	if needType == model.OrderNeedForOthers && recipient != nil {
		tableID = recipient.TableID
		tableNumber = recipient.Table.TableNumber
		tableName = recipient.Table.TableName
	}

	order := &entity.Order{
		CustomerID: customerID,
		TableID:    tableID,
		NeedType:   needType,
		SubTotal:   subTotal,
		Tax:        tax,
		Total:      total,
		Status:     "pending",
		Notes:      req.Notes,
		OrderItems: orderItems,
	}

	if isWishlist {
		order.Status = "pending_wishlist"
	}
	if recipient != nil {
		order.RecipientID = &recipient.ID
	}

	if err := uc.OrderRepo.CreateOrder(order); err != nil {
		return nil, err
	}

	if isWishlist {
		notification := map[string]interface{}{"type": "WISHLIST_UPDATED"}
		if payload, err := json.Marshal(notification); err == nil {
			uc.Hub.BroadcastToCustomers(payload)
		}
	} else {
		if fullOrder, err := uc.OrderRepo.FindByID(order.ID); err == nil {
			notification := map[string]interface{}{"type": "NEW_ORDER", "data": fullOrder}
			if payload, err := json.Marshal(notification); err == nil {
				uc.Hub.BroadcastAdmins <- payload
			}
		}
		if recipient != nil {
			var messageText string
			if needType == model.OrderNeedForOthers {
				messageText = fmt.Sprintf("%s mentraktirmu!", customer.Name)
			} else if needType == model.OrderNeedRequestTreat {
				messageText = fmt.Sprintf("%s minta ditraktir.", customer.Name)
			}
			chatMessage := &entity.ChatMessage{SenderID: customerID, RecipientID: recipient.ID, Text: messageText, OrderID: &order.ID}
			if err := uc.ChatRepo.CreateMessage(chatMessage); err == nil {
				uc.Hub.SendChatMessage(chatMessage)
			}
		}
	}

	if recipient != nil {
		recipientSummary = &model.OrderRecipient{
			CustomerID:  recipient.ID,
			Name:        recipient.Name,
			TableID:     recipient.TableID,
			TableNumber: recipient.Table.TableNumber,
		}
	}

	return &model.CreateOrderResponse{
		OrderID:      order.ID,
		CustomerID:   customer.ID,
		CustomerName: customer.Name,
		TableID:      tableID,
		TableNumber:  tableNumber,
		TableName:    tableName,
		NeedType:     needType,
		Recipient:    recipientSummary,
		Notes:        req.Notes,
		SubTotal:     subTotal,
		Tax:          tax,
		Total:        total,
		Items:        responseItems,
		CreatedAt:    order.CreatedAt,
	}, nil
}

func (uc *OrderUseCase) GetAllOrders() ([]entity.Order, error) {
	return uc.OrderRepo.FindAll()
}

func (uc *OrderUseCase) GetWishlistDetails(wishlistID uint) (*entity.Order, error) {
	return uc.OrderRepo.FindWishlistByID(wishlistID)
}

func (uc *OrderUseCase) AcceptWishlist(wishlistID, payerID uint) (*entity.Order, error) {
	wishlist, err := uc.OrderRepo.FindWishlistByID(wishlistID)
	if err != nil {
		return nil, errors.New("wishlist not found or has already been taken")
	}

	wishlist.Status = "pending"
	wishlist.PayerCustomerID = &payerID

	if err := uc.OrderRepo.UpdateOrder(wishlist); err != nil {
		return nil, err
	}

	if fullOrder, err := uc.OrderRepo.FindByID(wishlist.ID); err == nil {
		notification := map[string]interface{}{"type": "NEW_ORDER", "data": fullOrder}
		if payload, err := json.Marshal(notification); err == nil {
			uc.Hub.BroadcastAdmins <- payload
		}
	}

	notification := map[string]interface{}{"type": "WISHLIST_UPDATED"}
	if payload, err := json.Marshal(notification); err == nil {
		uc.Hub.BroadcastToCustomers(payload)
	}

	thankYouMessage := &entity.ChatMessage{
		SenderID:    wishlist.CustomerID,
		RecipientID: payerID,
		Text:        "Terima kasih orang baik...",
	}
	if err := uc.ChatRepo.CreateMessage(thankYouMessage); err == nil {
		uc.Hub.SendChatMessage(thankYouMessage)
	}

	return wishlist, nil
}
