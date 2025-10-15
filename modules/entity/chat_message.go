package entity

import (
	"gorm.io/gorm"
)

type ChatMessage struct {
	gorm.Model
	SenderID         uint
	RecipientID      uint
	Text             string
	ReplyToMessageID *uint
	IsRead           bool `gorm:"default:false;not null"`
	MenuID           *uint

	Sender         Customer `gorm:"foreignKey:SenderID"`
	Recipient      Customer `gorm:"foreignKey:RecipientID"`
	Menu           *Menu
	ReplyToMessage *ChatMessage `gorm:"foreignKey:ReplyToMessageID;references:ID"`
}
