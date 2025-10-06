package interfaces

import "coffee-chat-service/modules/entity"

type AdminRepositoryInterface interface {
	FindByUsername(username string) (*entity.Admin, error)
	Create(admin *entity.Admin) error
}
