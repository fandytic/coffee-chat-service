package usecase

import (
	"coffee-chat-service/modules/entity"
	interfaces "coffee-chat-service/modules/interface"
	"coffee-chat-service/modules/model"
)

type MenuUseCase struct {
	MenuRepo interfaces.MenuRepositoryInterface
}

func (uc *MenuUseCase) CreateMenu(req model.MenuRequest) (*entity.Menu, error) {
	menu := &entity.Menu{
		Name:     req.Name,
		Price:    req.Price,
		ImageURL: req.ImageURL,
	}
	err := uc.MenuRepo.Create(menu)
	return menu, err
}

func (uc *MenuUseCase) GetAllMenus(search string) ([]entity.Menu, error) {
	return uc.MenuRepo.FindAll(search)
}

func (uc *MenuUseCase) GetMenuByID(id uint) (*entity.Menu, error) {
	return uc.MenuRepo.FindByID(id)
}

func (uc *MenuUseCase) UpdateMenu(id uint, req model.MenuRequest) (*entity.Menu, error) {
	menu, err := uc.MenuRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	menu.Name = req.Name
	menu.Price = req.Price
	menu.ImageURL = req.ImageURL
	err = uc.MenuRepo.Update(menu)
	return menu, err
}

func (uc *MenuUseCase) DeleteMenu(id uint) error {
	return uc.MenuRepo.Delete(id)
}
