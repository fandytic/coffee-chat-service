package usecase

import (
	"coffee-chat-service/modules/entity"
	interfaces "coffee-chat-service/modules/interface"
	"coffee-chat-service/modules/model"
)

type ChatUseCase struct {
	ChatRepo interfaces.ChatRepositoryInterface
}

func (uc *ChatUseCase) MarkMessagesAsRead(senderID, recipientID uint) error {
	return uc.ChatRepo.MarkMessagesAsRead(senderID, recipientID)
}

func (uc *ChatUseCase) GetMessageHistory(user1ID, user2ID uint) ([]model.ChatHistoryMessage, error) {
	messages, err := uc.ChatRepo.GetMessageHistory(user1ID, user2ID)
	if err != nil {
		return nil, err
	}

	history := make([]model.ChatHistoryMessage, 0, len(messages))
	for _, message := range messages {
		history = append(history, buildChatHistoryMessage(&message))
	}

	return history, nil
}

func buildChatHistoryMessage(message *entity.ChatMessage) model.ChatHistoryMessage {
	history := model.ChatHistoryMessage{
		MessageID: message.ID,
		SenderID:  message.SenderID,
		Text:      message.Text,
		Timestamp: message.CreatedAt,
	}

	if message.Sender.ID != 0 {
		history.SenderName = message.Sender.Name
		history.SenderPhotoURL = message.Sender.PhotoURL

		if message.Sender.Table.ID != 0 {
			history.SenderTableNumber = message.Sender.Table.TableNumber

			if message.Sender.Table.Floor.ID != 0 {
				history.SenderFloorNumber = message.Sender.Table.Floor.FloorNumber
			}
		}
	}

	if message.ReplyToMessageID != nil && message.ReplyToMessage != nil {
		repliedSenderName := ""
		if message.ReplyToMessage.Sender.ID != 0 {
			repliedSenderName = message.ReplyToMessage.Sender.Name
		}

		reply := &model.ChatHistoryReply{
			ID:         message.ReplyToMessage.ID,
			Text:       message.ReplyToMessage.Text,
			SenderName: repliedSenderName,
		}

		if message.ReplyToMessage.Menu != nil {
			reply.Menu = &model.ChatHistoryMenu{
				ID:       message.ReplyToMessage.Menu.ID,
				Name:     message.ReplyToMessage.Menu.Name,
				Price:    message.ReplyToMessage.Menu.Price,
				ImageURL: message.ReplyToMessage.Menu.ImageURL,
			}
		}

		history.ReplyTo = reply
	}

	if message.Menu != nil {
		history.Menu = &model.ChatHistoryMenu{
			ID:       message.Menu.ID,
			Name:     message.Menu.Name,
			Price:    message.Menu.Price,
			ImageURL: message.Menu.ImageURL,
		}
	}

	return history
}
