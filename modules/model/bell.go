package model

import "time"

type BellNotificationPayload struct {
	CustomerName string    `json:"customer_name"`
	TableNumber  string    `json:"table_number"`
	FloorNumber  int       `json:"floor_number"`
	CallTime     time.Time `json:"call_time"`
}
