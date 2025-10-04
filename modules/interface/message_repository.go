package interfaces

import "coffee-chat-service/modules/entity"

type MessageRepositoryInterface interface {
	Create(message *entity.Message) error
	GetAll() ([]entity.Message, error)
}
