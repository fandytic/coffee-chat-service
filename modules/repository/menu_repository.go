package repository

import (
	"fmt"

	"gorm.io/gorm"

	"coffee-chat-service/modules/entity"
)

type MenuRepository struct {
	DB *gorm.DB
}

func NewMenuRepository(db *gorm.DB) *MenuRepository {
	return &MenuRepository{DB: db}
}

func (r *MenuRepository) Create(menu *entity.Menu) error {
	return r.DB.Create(menu).Error
}

func (r *MenuRepository) FindAll(search string) ([]entity.Menu, error) {
	var menus []entity.Menu
	query := r.DB.Order("created_at desc")
	if search != "" {
		query = query.Where("name ILIKE ?", "%"+search+"%")
	}
	err := query.Find(&menus).Error
	return menus, err
}

func (r *MenuRepository) FindByID(id uint) (*entity.Menu, error) {
	var menu entity.Menu
	err := r.DB.First(&menu, id).Error
	return &menu, err
}

func (r *MenuRepository) Update(menu *entity.Menu) error {
	return r.DB.Save(menu).Error
}

func (r *MenuRepository) Delete(id uint) error {
	result := r.DB.Delete(&entity.Menu{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("menu with ID %d not found", id)
	}
	return nil
}
