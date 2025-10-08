package model

import "time"

type CustomerCheckInRequest struct {
	TableID  uint   `json:"table_id"`
	Name     string `json:"name"`
	PhotoURL string `json:"photo_url,omitempty"` // Opsional
}

type CustomerCheckInResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	PhotoURL  string `json:"photo_url"`
	TableID   uint   `json:"table_id"`
	AuthToken string `json:"auth_token"`
}

type ActiveCustomerResponse struct {
	ID                  uint   `json:"id"`
	Name                string `json:"name"`
	PhotoURL            string `json:"photo_url"`
	TableNumber         string `json:"table_number"`
	UnreadMessagesCount int    `json:"unread_messages_count"`
}

type AllCustomersResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	PhotoURL    string    `json:"photo_url"`
	TableNumber string    `json:"table_number"`
	Status      string    `json:"status"`
	LastLogin   time.Time `json:"last_login"`
}
