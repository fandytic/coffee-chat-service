package interfaces

import "coffee-chat-service/modules/entity"

type ChatRepositoryInterface interface {
	MarkMessagesAsRead(senderID, recipientID uint) error
	FindLastMessages(userID uint) (map[uint]*entity.ChatMessage, error)
	GetMessageHistory(user1ID, user2ID uint) ([]entity.ChatMessage, error)
	CreateMessage(message *entity.ChatMessage) error
}
