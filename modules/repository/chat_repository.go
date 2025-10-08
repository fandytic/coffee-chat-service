package repository

import (
	"gorm.io/gorm"

	"coffee-chat-service/modules/entity"
)

type ChatRepository struct {
	DB *gorm.DB
}

func NewChatRepository(db *gorm.DB) *ChatRepository {
	return &ChatRepository{DB: db}
}

func (r *ChatRepository) MarkMessagesAsRead(senderID, recipientID uint) error {
	return r.DB.Model(&entity.ChatMessage{}).
		Where("sender_id = ? AND recipient_id = ? AND is_read = ?", senderID, recipientID, false).
		Update("is_read", true).Error
}
