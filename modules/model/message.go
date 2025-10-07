package model

type SendMessageRequest struct {
	User string `json:"user"`
	Text string `json:"text"`
}
