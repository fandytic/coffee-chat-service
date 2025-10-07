package repository

import (
	"coffee-chat-service/modules/entity"

	"gorm.io/gorm"
)

type DashboardRepository struct {
	DB *gorm.DB
}

func NewDashboardRepository(db *gorm.DB) *DashboardRepository {
	return &DashboardRepository{DB: db}
}

func (r *DashboardRepository) CountTotalTables() (int64, error) {
	var count int64
	err := r.DB.Model(&entity.Table{}).Count(&count).Error
	return count, err
}

func (r *DashboardRepository) CountActiveUsers() (int64, error) {
	var count int64
	err := r.DB.Model(&entity.Customer{}).Where("status = ?", "active").Count(&count).Error
	return count, err
}

func (r *DashboardRepository) CountOccupiedTables() (int64, error) {
	var count int64
	err := r.DB.Model(&entity.Customer{}).Where("status = ?", "active").Distinct("table_id").Count(&count).Error
	return count, err
}
