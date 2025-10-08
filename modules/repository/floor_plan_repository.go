package repository

import (
	"fmt"
	"log"
	"os"

	"gorm.io/gorm"

	"coffee-chat-service/modules/entity"
	interfaces "coffee-chat-service/modules/interface"
	"coffee-chat-service/modules/model"
)

type FloorPlanRepository struct {
	DB *gorm.DB
}

func NewFloorPlanRepository(db *gorm.DB) *FloorPlanRepository {
	return &FloorPlanRepository{DB: db}
}

func (r *FloorPlanRepository) CreateFloorPlan(floor *entity.Floor, tables []entity.Table) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(floor).Error; err != nil {
			return err
		}
		for i := range tables {
			tables[i].FloorID = floor.ID
		}
		if len(tables) > 0 {
			if err := tx.Create(&tables).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *FloorPlanRepository) FindFloorByNumber(floorNumber int) (*entity.Floor, error) {
	var floor entity.Floor
	err := r.DB.Preload("Tables").First(&floor, "floor_number = ?", floorNumber).Error
	return &floor, err
}

func (r *FloorPlanRepository) FindAllFloors() ([]entity.Floor, error) {
	var floors []entity.Floor
	err := r.DB.Order("floor_number asc").Find(&floors).Error
	return floors, err
}

func (r *FloorPlanRepository) UpdateTable(tableID uint, req model.UpdateTableRequest) (*entity.Table, error) {
	var table entity.Table
	if err := r.DB.First(&table, tableID).Error; err != nil {
		return nil, fmt.Errorf("table with ID %d not found", tableID)
	}
	table.TableName = req.TableName
	table.XCoordinate = req.XCoordinate
	table.YCoordinate = req.YCoordinate
	err := r.DB.Save(&table).Error
	return &table, err
}

func (r *FloorPlanRepository) DeleteTable(tableID uint) error {
	result := r.DB.Delete(&entity.Table{}, tableID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("table with ID %d not found", tableID)
	}
	return nil
}

func (r *FloorPlanRepository) DeleteFloorAndTables(floorID uint, imagePath string) error {
	err := r.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Unscoped().Where("floor_id = ?", floorID).Delete(&entity.Table{}).Error; err != nil {
			return err
		}
		if err := tx.Unscoped().Delete(&entity.Floor{}, floorID).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	if err := os.Remove(imagePath); err != nil {
		log.Printf("Warning: could not delete image file %s: %v", imagePath, err)
	} else {
		log.Printf("Successfully deleted image file: %s", imagePath)
	}
	return nil
}

func (r *FloorPlanRepository) CountUsersPerTable() (map[uint]int, error) {
	var results []interfaces.UserCountResult
	err := r.DB.Model(&entity.Customer{}).
		Select("table_id, count(*) as count").
		Where("status = ?", "active").
		Group("table_id").
		Find(&results).Error
	if err != nil {
		return nil, err
	}

	userCounts := make(map[uint]int)
	for _, result := range results {
		userCounts[result.TableID] = result.Count
	}
	return userCounts, nil
}

func (r *FloorPlanRepository) FindFloorByID(floorID uint) (*entity.Floor, error) {
	var floor entity.Floor
	err := r.DB.First(&floor, floorID).Error
	return &floor, err
}
