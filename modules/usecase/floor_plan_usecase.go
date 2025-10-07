package usecase

import (
	"fmt"
	"strings"

	"coffee-chat-service/modules/entity"
	"coffee-chat-service/modules/model"

	interfaces "coffee-chat-service/modules/interface"
)

type FloorPlanUseCase struct {
	FloorPlanRepo interfaces.FloorPlanRepositoryInterface
}

func (uc *FloorPlanUseCase) CreateFloorPlan(req model.CreateFloorPlanRequest) (*model.FloorPlanResponse, error) {
	floor := &entity.Floor{
		FloorNumber: req.FloorNumber,
		ImageURL:    req.ImageURL,
	}

	var tables []entity.Table
	for _, td := range req.Tables {
		tables = append(tables, entity.Table{
			TableNumber: td.TableNumber,
			TableName:   td.TableName,
			XCoordinate: td.XCoordinate,
			YCoordinate: td.YCoordinate,
		})
	}

	if err := uc.FloorPlanRepo.CreateFloorPlan(floor, tables); err != nil {
		return nil, err
	}

	return uc.GetFloorPlanByNumber(req.FloorNumber)
}

func (uc *FloorPlanUseCase) GetFloorPlanByNumber(floorNumber int) (*model.FloorPlanResponse, error) {
	floor, err := uc.FloorPlanRepo.FindFloorByNumber(floorNumber)
	if err != nil {
		return nil, err
	}

	userCounts, err := uc.FloorPlanRepo.CountUsersPerTable()
	if err != nil {
		return nil, err
	}

	tables := make([]model.TableData, 0, len(floor.Tables))
	for _, t := range floor.Tables {
		tables = append(tables, model.TableData{
			ID:               t.ID,
			TableNumber:      t.TableNumber,
			TableName:        t.TableName,
			XCoordinate:      t.XCoordinate,
			YCoordinate:      t.YCoordinate,
			ActiveUsersCount: userCounts[t.ID],
		})
	}

	return &model.FloorPlanResponse{
		ID:          floor.ID,
		FloorNumber: floor.FloorNumber,
		ImageURL:    floor.ImageURL,
		Tables:      tables,
	}, nil
}

func (uc *FloorPlanUseCase) GetAllFloors() ([]model.FloorInfoResponse, error) {
	floors, err := uc.FloorPlanRepo.FindAllFloors()
	if err != nil {
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
	return uc.FloorPlanRepo.UpdateTable(tableID, req)
}

func (uc *FloorPlanUseCase) DeleteTable(tableID uint) error {
	return uc.FloorPlanRepo.DeleteTable(tableID)
}

func (uc *FloorPlanUseCase) DeleteFloor(floorID uint) error {
	// 1. Ambil data lantai berdasarkan ID untuk mendapatkan path gambar
	floor, err := uc.FloorPlanRepo.FindFloorByID(floorID)
	if err != nil {
		// Jika tidak ditemukan atau ada error lain, kembalikan error
		return fmt.Errorf("floor with ID %d not found", floorID)
	}

	// 2. Siapkan path file untuk dihapus
	imagePath := strings.TrimPrefix(floor.ImageURL, "/")

	// 3. Panggil repository untuk menghapus data dari DB dan file dari server
	return uc.FloorPlanRepo.DeleteFloorAndTables(floorID, imagePath)
}
