package usecase

import (
	interfaces "coffee-chat-service/modules/interface"
)

type ChatUseCase struct {
	ChatRepo interfaces.ChatRepositoryInterface
}

func (uc *ChatUseCase) MarkMessagesAsRead(senderID, recipientID uint) error {
	return uc.ChatRepo.MarkMessagesAsRead(senderID, recipientID)
}
