package usecase

import (
	"coffee-chat-service/modules/entity"
	interfaces "coffee-chat-service/modules/interface"
)

type ChatUseCase struct {
	ChatRepo interfaces.ChatRepositoryInterface
}

func (uc *ChatUseCase) MarkMessagesAsRead(senderID, recipientID uint) error {
	return uc.ChatRepo.MarkMessagesAsRead(senderID, recipientID)
}

func (uc *ChatUseCase) GetMessageHistory(user1ID, user2ID uint) ([]entity.ChatMessage, error) {
	return uc.ChatRepo.GetMessageHistory(user1ID, user2ID)
}
