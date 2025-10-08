package interfaces

import "coffee-chat-service/modules/entity"

type MenuRepositoryInterface interface {
	Create(menu *entity.Menu) error
	FindAll(search string) ([]entity.Menu, error)
	FindByID(id uint) (*entity.Menu, error)
	Update(menu *entity.Menu) error
	Delete(id uint) error
}
