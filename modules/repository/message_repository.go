package repository

import (
	"coffee-chat-service/modules/entity"

	"gorm.io/gorm"
)

type MessageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *MessageRepository {
	return &MessageRepository{db: db}
}

func (r *MessageRepository) Create(message *entity.Message) error {
	return r.db.Create(message).Error
}

func (r *MessageRepository) GetAll() ([]entity.Message, error) {
	var messages []entity.Message
	err := r.db.Order("timestamp asc").Find(&messages).Error
	return messages, err
}
