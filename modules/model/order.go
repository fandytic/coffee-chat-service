package model

type CreateOrderRequest struct {
	Notes      string             `json:"notes"`
	OrderItems []OrderItemRequest `json:"items"`
}

type OrderItemRequest struct {
	MenuID   uint `json:"menu_id"`
	Quantity int  `json:"quantity"`
}
