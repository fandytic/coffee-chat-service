package repository

import (
	"coffee-chat-service/modules/entity"

	"gorm.io/gorm"
)

type AdminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *AdminRepository {
	return &AdminRepository{db: db}
}

func (r *AdminRepository) FindByUsername(username string) (*entity.Admin, error) {
	var admin entity.Admin
	err := r.db.Where("username = ?", username).First(&admin).Error
	return &admin, err
}

func (r *AdminRepository) Create(admin *entity.Admin) error {
	return r.db.Create(admin).Error
}
