package repository

import (
	"gorm.io/gorm"

	"coffee-chat-service/modules/entity"
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

func (r *AdminRepository) FindByID(id uint) (*entity.Admin, error) {
	var admin entity.Admin
	err := r.db.First(&admin, id).Error
	return &admin, err
}

func (r *AdminRepository) Create(admin *entity.Admin) error {
	return r.db.Create(admin).Error
}

func (r *AdminRepository) Update(admin *entity.Admin) error {
	return r.db.Save(admin).Error
}
