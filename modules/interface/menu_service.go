package interfaces

import (
	"coffee-chat-service/modules/entity"
	"coffee-chat-service/modules/model"
)

type MenuServiceInterface interface {
	CreateMenu(req model.MenuRequest) (*entity.Menu, error)
	GetAllMenus(search string) ([]entity.Menu, error)
	GetMenuByID(id uint) (*entity.Menu, error)
	UpdateMenu(id uint, req model.MenuRequest) (*entity.Menu, error)
	DeleteMenu(id uint) error
}
