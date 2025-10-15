package model

import "time"

const (
	OrderNeedSelf         = "self_order"
	OrderNeedForOthers    = "order_for_other"
	OrderNeedRequestTreat = "request_treat"
)

type CreateOrderRequest struct {
	Notes               string             `json:"notes"`
	NeedType            string             `json:"need_type"`
	RecipientCustomerID *uint              `json:"recipient_customer_id,omitempty"`
	OrderItems          []OrderItemRequest `json:"items"`
}

type OrderItemRequest struct {
	MenuID   uint `json:"menu_id"`
	Quantity int  `json:"quantity"`
}

type CreateOrderResponse struct {
	OrderID      uint               `json:"order_id"`
	CustomerID   uint               `json:"customer_id"`
	CustomerName string             `json:"customer_name"`
	TableID      uint               `json:"table_id"`
	TableNumber  string             `json:"table_number"`
	TableName    string             `json:"table_name"`
	NeedType     string             `json:"need_type"`
	Recipient    *OrderRecipient    `json:"recipient,omitempty"`
	Notes        string             `json:"notes,omitempty"`
	SubTotal     float64            `json:"sub_total"`
	Tax          float64            `json:"tax"`
	Total        float64            `json:"total"`
	Items        []OrderItemSummary `json:"items"`
	CreatedAt    time.Time          `json:"created_at"`
}

type OrderRecipient struct {
	CustomerID  uint   `json:"customer_id"`
	Name        string `json:"name"`
	TableID     uint   `json:"table_id"`
	TableNumber string `json:"table_number"`
}

type OrderItemSummary struct {
	MenuID     uint    `json:"menu_id"`
	MenuName   string  `json:"menu_name"`
	Quantity   int     `json:"quantity"`
	UnitPrice  float64 `json:"unit_price"`
	TotalPrice float64 `json:"total_price"`
}
