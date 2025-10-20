package entity

import "gorm.io/gorm"

type ChatGroup struct {
	gorm.Model
	Name      string `gorm:"not null"`
	CreatorID uint   `gorm:"not null"` // Customer ID
	Creator   Customer
	Members   []ChatGroupMember
	Messages  []GroupChatMessage
}

type ChatGroupMember struct {
	gorm.Model
	ChatGroupID uint `gorm:"uniqueIndex:idx_group_customer"`
	CustomerID  uint `gorm:"uniqueIndex:idx_group_customer"`
	ChatGroup   ChatGroup
	Customer    Customer
}

type GroupChatMessage struct {
	gorm.Model
	ChatGroupID      uint   `gorm:"not null"`
	SenderID         uint   `gorm:"not null"` // Customer ID
	Text             string `gorm:"not null"`
	ReplyToMessageID *uint
	MenuID           *uint
	OrderID          *uint

	ChatGroup      ChatGroup
	Sender         Customer `gorm:"foreignKey:SenderID"`
	Menu           *Menu
	Order          *Order
	ReplyToMessage *GroupChatMessage `gorm:"foreignKey:ReplyToMessageID;references:ID"`
}
