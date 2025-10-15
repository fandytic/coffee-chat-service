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

func (r *ChatRepository) CreateMessage(message *entity.ChatMessage) error {
	return r.DB.Create(message).Error
}

func (r *ChatRepository) FindLastMessages(userID uint) (map[uint]*entity.ChatMessage, error) {
	var messages []entity.ChatMessage

	err := r.DB.Raw(`
		SELECT * FROM chat_messages WHERE id IN (
			SELECT MAX(id) FROM chat_messages
			WHERE sender_id = ? OR recipient_id = ?
			GROUP BY (
				CASE
					WHEN sender_id = ? THEN recipient_id
					ELSE sender_id
				END
			)
		)
	`, userID, userID, userID).Scan(&messages).Error

	if err != nil {
		return nil, err
	}

	lastMessagesMap := make(map[uint]*entity.ChatMessage)
	for i, msg := range messages {
		if msg.SenderID == userID {
			lastMessagesMap[msg.RecipientID] = &messages[i]
		} else {
			lastMessagesMap[msg.SenderID] = &messages[i]
		}
	}

	return lastMessagesMap, nil
}

func (r *ChatRepository) GetMessageHistory(user1ID, user2ID uint) ([]entity.ChatMessage, error) {
	var messages []entity.ChatMessage
	err := r.DB.Preload("Sender.Table.Floor").
		Preload("ReplyToMessage.Sender").
		Preload("ReplyToMessage.Menu").
		Preload("Menu").
		Preload("Order.Table.Floor").
		Preload("Order.OrderItems.Menu").
		Where("(sender_id = ? AND recipient_id = ?) OR (sender_id = ? AND recipient_id = ?)",
			user1ID, user2ID, user2ID, user1ID).
		Order("created_at asc").
		Find(&messages).Error
	return messages, err
}
