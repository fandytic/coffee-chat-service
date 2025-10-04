package model

// SendMessageRequest adalah model untuk body request pengiriman pesan.
type SendMessageRequest struct {
	User string `json:"user"`
	Text string `json:"text"`
}
