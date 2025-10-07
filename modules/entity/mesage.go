package entity

import "time"

type Message struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	User      string    `json:"user" gorm:"not null"`
	Text      string    `json:"text" gorm:"not null"`
	Timestamp time.Time `json:"timestamp"`
}
