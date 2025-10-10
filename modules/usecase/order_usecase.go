package usecase

import (
	"encoding/json"
	"errors"

	"coffee-chat-service/modules/entity"
	"coffee-chat-service/modules/model"
	"coffee-chat-service/modules/repository"
	"coffee-chat-service/modules/websocket"
)

const TAX_RATE = 0.11 // Pajak 11%

type OrderUseCase struct {
	OrderRepo *repository.OrderRepository
	Hub       *websocket.Hub
}

func (uc *OrderUseCase) CreateOrder(customerID uint, req model.CreateOrderRequest) (*entity.Order, error) {
	if len(req.OrderItems) == 0 {
		return nil, errors.New("order must have at least one item")
	}

	var menuIDs []uint
	for _, item := range req.OrderItems {
		menuIDs = append(menuIDs, item.MenuID)
	}

	priceMap, err := uc.OrderRepo.FindMenuPrices(menuIDs)
	if err != nil {
		return nil, err
	}
	if len(priceMap) != len(menuIDs) {
		return nil, errors.New("one or more menu items not found")
	}

	var subTotal float64
	orderItems := make([]entity.OrderItem, 0)
	for _, item := range req.OrderItems {
		price := priceMap[item.MenuID]
		subTotal += price * float64(item.Quantity)
		orderItems = append(orderItems, entity.OrderItem{
			MenuID:   item.MenuID,
			Quantity: item.Quantity,
			Price:    price,
		})
	}

	tax := subTotal * TAX_RATE
	total := subTotal + tax

	order := &entity.Order{
		CustomerID: customerID,
		SubTotal:   subTotal,
		Tax:        tax,
		Total:      total,
		Status:     "pending",
		Notes:      req.Notes,
		OrderItems: orderItems,
	}

	if err := uc.OrderRepo.CreateOrder(order); err != nil {
		return nil, err
	}
	fullOrder, err := uc.OrderRepo.FindByID(order.ID)
	if err == nil {
		notification := map[string]interface{}{
			"type": "NEW_ORDER",
			"data": fullOrder,
		}
		notificationJSON, _ := json.Marshal(notification)

		uc.Hub.BroadcastAdmins <- notificationJSON
	}

	return order, nil
}

func (uc *OrderUseCase) GetAllOrders() ([]entity.Order, error) {
	return uc.OrderRepo.FindAll()
}
