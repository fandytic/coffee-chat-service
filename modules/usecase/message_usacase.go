package usecase

import (
	"time"

	"coffee-chat-service/modules/entity"
	interfaces "coffee-chat-service/modules/interface"
	"coffee-chat-service/modules/model"
	"coffee-chat-service/modules/websocket"
)

type MessageUseCase struct {
	Repo interfaces.MessageRepositoryInterface
	Hub  *websocket.Hub
}

func (uc *MessageUseCase) SaveAndBroadcastMessage(req model.SendMessageRequest) (*entity.Message, error) {
	message := &entity.Message{
		User:      req.User,
		Text:      req.Text,
		Timestamp: time.Now(),
	}

	if err := uc.Repo.Create(message); err != nil {
		return nil, err
	}

	uc.Hub.Broadcast <- message

	return message, nil
}

func (uc *MessageUseCase) GetAllMessages() ([]entity.Message, error) {
	return uc.Repo.GetAll()
}
