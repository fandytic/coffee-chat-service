package model

import "time"

type SendMessageRequest struct {
	User string `json:"user"`
	Text string `json:"text"`
}

type ChatHistoryReply struct {
	ID         uint             `json:"id"`
	Text       string           `json:"text"`
	SenderName string           `json:"sender_name"`
	Menu       *ChatHistoryMenu `json:"menu,omitempty"`
}

type ChatHistoryMenu struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	ImageURL string  `json:"image_url"`
}

type ChatHistoryOrderItem struct {
	ID       uint             `json:"id"`
	MenuID   uint             `json:"menu_id"`
	Quantity int              `json:"quantity"`
	Price    float64          `json:"price"`
	Menu     *ChatHistoryMenu `json:"menu,omitempty"`
}

type ChatHistoryOrder struct {
	ID               uint                   `json:"id"`
	CustomerID       uint                   `json:"customer_id"`
	RecipientID      *uint                  `json:"recipient_id,omitempty"`
	TableID          uint                   `json:"table_id"`
	TableNumber      string                 `json:"table_number"`
	TableName        string                 `json:"table_name"`
	TableFloorNumber int                    `json:"table_floor_number"`
	NeedType         string                 `json:"need_type"`
	Notes            string                 `json:"notes,omitempty"`
	SubTotal         float64                `json:"sub_total"`
	Tax              float64                `json:"tax"`
	Total            float64                `json:"total"`
	OrderItems       []ChatHistoryOrderItem `json:"order_items"`
}

type ChatHistoryMessage struct {
	MessageID         uint              `json:"message_id"`
	SenderID          uint              `json:"sender_id"`
	SenderName        string            `json:"sender_name"`
	SenderPhotoURL    string            `json:"sender_photo_url"`
	SenderTableNumber string            `json:"sender_table_number"`
	SenderFloorNumber int               `json:"sender_floor_number"`
	Text              string            `json:"text"`
	Timestamp         time.Time         `json:"timestamp"`
	ReplyTo           *ChatHistoryReply `json:"reply_to,omitempty"`
	Menu              *ChatHistoryMenu  `json:"menu,omitempty"`
	Order             *ChatHistoryOrder `json:"order,omitempty"`
}
