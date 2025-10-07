package usecase

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"coffee-chat-service/modules/entity"
	"coffee-chat-service/modules/model"

	// "coffee-chat-service/modules/repository"
	"gorm.io/gorm"
)

type FloorPlanUseCase struct {
	DB *gorm.DB
}

func (uc *FloorPlanUseCase) CreateFloorPlan(req model.CreateFloorPlanRequest) (*model.FloorPlanResponse, error) {
	if req.ImageURL == "" {
		return nil, errors.New("image_url is required")
	}

	var createdFloor entity.Floor
	err := uc.DB.Transaction(func(tx *gorm.DB) error {
		floor := entity.Floor{
			FloorNumber: req.FloorNumber,
			ImageURL:    req.ImageURL,
		}
		if err := tx.Create(&floor).Error; err != nil {
			return err
		}

		for _, td := range req.Tables {
			table := entity.Table{
				TableNumber: td.TableNumber,
				TableName:   td.TableName,
				XCoordinate: td.XCoordinate,
				YCoordinate: td.YCoordinate,
				FloorID:     floor.ID,
			}
			if err := tx.Create(&table).Error; err != nil {
				return err
			}
		}
		createdFloor = floor
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to save data to database: %w", err)
	}

	return uc.GetFloorPlanByNumber(createdFloor.FloorNumber)
}

func (uc *FloorPlanUseCase) GetFloorPlanByNumber(floorNumber int) (*model.FloorPlanResponse, error) {
	var floor entity.Floor
	if err := uc.DB.Preload("Tables").First(&floor, "floor_number = ?", floorNumber).Error; err != nil {
		return nil, err
	}

	tables := make([]model.TableData, 0, len(floor.Tables))
	for _, t := range floor.Tables {
		tables = append(tables, model.TableData{
			TableNumber: t.TableNumber,
			TableName:   t.TableName,
			XCoordinate: t.XCoordinate,
			YCoordinate: t.YCoordinate,
		})
	}

	response := &model.FloorPlanResponse{
		ID:          floor.ID,
		FloorNumber: floor.FloorNumber,
		ImageURL:    floor.ImageURL,
		Tables:      tables,
	}

	return response, nil
}

func (uc *FloorPlanUseCase) GetAllFloors() ([]model.FloorInfoResponse, error) {
	var floors []entity.Floor
	if err := uc.DB.Order("floor_number asc").Find(&floors).Error; err != nil {
		return nil, err
	}

	response := make([]model.FloorInfoResponse, 0, len(floors))
	for _, floor := range floors {
		response = append(response, model.FloorInfoResponse{
			ID:          floor.ID,
			FloorNumber: floor.FloorNumber,
		})
	}
	return response, nil
}

func (uc *FloorPlanUseCase) UpdateTable(tableID uint, req model.UpdateTableRequest) (*entity.Table, error) {
	var table entity.Table
	if err := uc.DB.First(&table, tableID).Error; err != nil {
		return nil, fmt.Errorf("table with ID %d not found", tableID)
	}

	table.TableName = req.TableName
	table.XCoordinate = req.XCoordinate
	table.YCoordinate = req.YCoordinate

	if err := uc.DB.Save(&table).Error; err != nil {
		return nil, err
	}

	return &table, nil
}

func (uc *FloorPlanUseCase) DeleteTable(tableID uint) error {
	result := uc.DB.Delete(&entity.Table{}, tableID)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("table with ID %d not found", tableID)
	}

	return nil
}

func (uc *FloorPlanUseCase) DeleteFloor(floorID uint) error {
	var floor entity.Floor
	if err := uc.DB.First(&floor, floorID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("floor with ID %d not found", floorID)
		}
		return err
	}

	imagePath := strings.TrimPrefix(floor.ImageURL, "/")
	log.Printf("Attempting to delete image file at path: ./%s", imagePath)

	err := uc.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Unscoped().Where("floor_id = ?", floorID).Delete(&entity.Table{}).Error; err != nil {
			return err
		}

		if err := tx.Unscoped().Delete(&entity.Floor{}, floorID).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to delete floor data from database: %w", err)
	}

	if err := os.Remove(imagePath); err != nil {
		log.Printf("Warning: could not delete image file %s: %v", imagePath, err)
	} else {
		log.Printf("Successfully deleted image file: %s", imagePath)
	}

	return nil
}
