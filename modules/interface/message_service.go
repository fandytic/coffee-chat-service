package interfaces

import (
	"coffee-chat-service/modules/entity"
	"coffee-chat-service/modules/model"
)

type MessageServiceInterface interface {
	SaveAndBroadcastMessage(req model.SendMessageRequest) (*entity.Message, error)
	GetAllMessages() ([]entity.Message, error)
}
