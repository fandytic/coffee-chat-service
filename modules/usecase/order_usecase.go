package usecase

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"strings"

	"coffee-chat-service/modules/entity"
	interfaces "coffee-chat-service/modules/interface"
	"coffee-chat-service/modules/model"
	"coffee-chat-service/modules/repository"
	"coffee-chat-service/modules/websocket"
	"gorm.io/gorm"
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

	switch needType {
	case model.OrderNeedSelf, model.OrderNeedForOthers, model.OrderNeedRequestTreat:
	default:
		return nil, &model.ValidationError{Message: "invalid need_type value"}
	}

	customer, err := uc.OrderRepo.FindCustomerWithTable(customerID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &model.ValidationError{Message: "customer not found"}
		}
		return nil, err
	}
	if customer.TableID == 0 {
		return nil, &model.ValidationError{Message: "customer does not have an assigned table"}
	}

	var recipient *entity.Customer
	if needType != model.OrderNeedSelf {
		if req.RecipientCustomerID == nil {
			return nil, &model.ValidationError{Message: "recipient_customer_id is required"}
		}
		if *req.RecipientCustomerID == customerID && needType == model.OrderNeedForOthers {
			return nil, &model.ValidationError{Message: "recipient must be different from the ordering customer"}
		}

		recipient, err = uc.OrderRepo.FindCustomerWithTable(*req.RecipientCustomerID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, &model.ValidationError{Message: "recipient not found"}
			}
			return nil, err
		}
		if recipient.Status != "active" {
			return nil, &model.ValidationError{Message: "recipient is not active"}
		}
	}

	uniqueMenuIDs := make(map[uint]struct{})
	menuIDs := make([]uint, 0, len(req.OrderItems))
	for _, item := range req.OrderItems {
		if item.MenuID == 0 {
			return nil, &model.ValidationError{Message: "menu_id is required for each item"}
		}
		if item.Quantity <= 0 {
			return nil, &model.ValidationError{Message: "quantity must be greater than zero"}
		}

		if _, exists := uniqueMenuIDs[item.MenuID]; !exists {
			uniqueMenuIDs[item.MenuID] = struct{}{}
			menuIDs = append(menuIDs, item.MenuID)
		}
	}

	menuMap, err := uc.OrderRepo.FindMenusByIDs(menuIDs)
	if err != nil {
		return nil, err
	}
	if len(menuMap) != len(menuIDs) {
		return nil, &model.ValidationError{Message: "one or more menu items not found"}
	}

	var subTotal float64
	orderItems := make([]entity.OrderItem, 0, len(req.OrderItems))
	responseItems := make([]model.OrderItemSummary, 0, len(req.OrderItems))
	summaryLines := make([]string, 0, len(req.OrderItems))
	for _, item := range req.OrderItems {
		menu := menuMap[item.MenuID]
		lineTotal := menu.Price * float64(item.Quantity)

		subTotal += lineTotal
		orderItems = append(orderItems, entity.OrderItem{
			MenuID:   item.MenuID,
			Quantity: item.Quantity,
			Price:    menu.Price,
		})

		responseItems = append(responseItems, model.OrderItemSummary{
			MenuID:     menu.ID,
			MenuName:   menu.Name,
			Quantity:   item.Quantity,
			UnitPrice:  menu.Price,
			TotalPrice: lineTotal,
		})

		summaryLines = append(summaryLines, fmt.Sprintf("- %s x%d (%s)", menu.Name, item.Quantity, formatCurrency(lineTotal)))
	}

	tax := subTotal * TAX_RATE
	total := subTotal + tax

	tableID := customer.TableID
	tableNumber := customer.Table.TableNumber
	tableName := customer.Table.TableName
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

	if recipient != nil {
		order.RecipientID = new(uint)
		*order.RecipientID = recipient.ID
	}

	if err := uc.OrderRepo.CreateOrder(order); err != nil {
		return nil, err
	}

	if fullOrder, err := uc.OrderRepo.FindByID(order.ID); err == nil {
		notification := map[string]interface{}{
			"type": "NEW_ORDER",
			"data": fullOrder,
		}

		if payload, err := json.Marshal(notification); err != nil {
			log.Printf("failed to marshal order notification: %v", err)
		} else {
			uc.Hub.BroadcastAdmins <- payload
		}
	}

	var recipientSummary *model.OrderRecipient
	if recipient != nil {
		messageText := buildChatMessage(needType, customer.Name, recipient.Name, summaryLines, subTotal, tax, total, req.Notes, tableNumber)
		chatMessage := &entity.ChatMessage{
			SenderID:    customerID,
			RecipientID: recipient.ID,
			Text:        messageText,
			OrderID:     &order.ID,
		}

		if needType == model.OrderNeedRequestTreat && len(req.OrderItems) > 0 {
			firstMenuID := req.OrderItems[0].MenuID
			if menu, ok := menuMap[firstMenuID]; ok {
				chatMessage.MenuID = new(uint)
				*chatMessage.MenuID = menu.ID
			}
		}
		if err := uc.ChatRepo.CreateMessage(chatMessage); err != nil {
			log.Printf("failed to create chat message notification: %v", err)
		} else {
			uc.Hub.SendChatMessage(chatMessage)
		}
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

func formatCurrency(amount float64) string {
	rounded := int64(math.Round(amount))
	if rounded == 0 {
		return "Rp 0"
	}

	negative := false
	if rounded < 0 {
		negative = true
		rounded = -rounded
	}

	digits := fmt.Sprintf("%d", rounded)
	var parts []string
	for len(digits) > 3 {
		parts = append([]string{digits[len(digits)-3:]}, parts...)
		digits = digits[:len(digits)-3]
	}
	parts = append([]string{digits}, parts...)
	formatted := strings.Join(parts, ".")
	if negative {
		formatted = "-" + formatted
	}
	return "Rp " + formatted
}

func buildChatMessage(needType, senderName, recipientName string, itemLines []string, subTotal, tax, total float64, notes, tableNumber string) string {
	var builder strings.Builder

	switch needType {
	case model.OrderNeedForOthers:
		builder.WriteString(fmt.Sprintf("Halo %s! %s baru saja memesankan untukmu.\n", recipientName, senderName))
	case model.OrderNeedRequestTreat:
		builder.WriteString(fmt.Sprintf("Halo %s! %s meminta traktiran. Berikut detail pesanannya.\n", recipientName, senderName))
	default:
		builder.WriteString(fmt.Sprintf("Halo %s! %s membuat pesanan baru.\n", recipientName, senderName))
	}

	if tableNumber != "" {
		builder.WriteString(fmt.Sprintf("Meja: %s\n", tableNumber))
	}

	builder.WriteString("\nDetail Pesanan:\n")
	builder.WriteString(strings.Join(itemLines, "\n"))
	builder.WriteString("\n")
	builder.WriteString(fmt.Sprintf("\nSub Total: %s\n", formatCurrency(subTotal)))
	builder.WriteString(fmt.Sprintf("Pajak (11%%): %s\n", formatCurrency(tax)))
	builder.WriteString(fmt.Sprintf("Total: %s", formatCurrency(total)))

	if strings.TrimSpace(notes) != "" {
		builder.WriteString(fmt.Sprintf("\nCatatan: %s", notes))
	}

	return builder.String()
}
