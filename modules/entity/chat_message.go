package entity

import (
	"gorm.io/gorm"
)

type ChatMessage struct {
	gorm.Model
	SenderID         uint
	RecipientID      uint
	Text             string
	ReplyToMessageID *uint // Pointer agar bisa bernilai null jika bukan balasan
	IsRead           bool  `gorm:"default:false;not null"`

	Sender    Customer `gorm:"foreignKey:SenderID"`
	Recipient Customer `gorm:"foreignKey:RecipientID"`
}
