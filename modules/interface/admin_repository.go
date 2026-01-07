package interfaces

import "coffee-chat-service/modules/entity"

type AdminRepositoryInterface interface {
	FindByUsername(username string) (*entity.Admin, error)
	FindByID(id uint) (*entity.Admin, error)
	Create(admin *entity.Admin) error
	Update(admin *entity.Admin) error
}
