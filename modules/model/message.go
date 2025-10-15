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
}
