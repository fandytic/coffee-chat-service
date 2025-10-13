package model

import "time"

type CustomerCheckInRequest struct {
	TableID  uint   `json:"table_id"`
	Name     string `json:"name"`
	PhotoURL string `json:"photo_url,omitempty"` // Opsional
}

type CustomerCheckInResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	PhotoURL    string `json:"photo_url"`
	TableID     uint   `json:"table_id"`
	TableNumber string `json:"table_number"`
	FloorNumber int    `json:"floor_number"`
	AuthToken   string `json:"auth_token"`
}

type LastMessage struct {
	Text      string    `json:"text"`
	Timestamp time.Time `json:"timestamp"`
}

type ActiveCustomerResponse struct {
	ID                  uint         `json:"id"`
	Name                string       `json:"name"`
	PhotoURL            string       `json:"photo_url"`
	TableNumber         string       `json:"table_number"`
	FloorNumber         int          `json:"floor_number"`
	UnreadMessagesCount int          `json:"unread_messages_count"`
	LastMessage         *LastMessage `json:"last_message,omitempty"`
}

type PaginatedActiveCustomersResponse struct {
	Total     int                      `json:"total"`
	Customers []ActiveCustomerResponse `json:"customers"`
}

type AllCustomersResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	PhotoURL    string    `json:"photo_url"`
	TableNumber string    `json:"table_number"`
	Status      string    `json:"status"`
	LastLogin   time.Time `json:"last_login"`
}
