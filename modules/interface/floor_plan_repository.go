package interfaces

import (
	"coffee-chat-service/modules/entity"
	"coffee-chat-service/modules/model"
)

type UserCountResult struct {
	TableID uint
	Count   int
}

type FloorPlanRepositoryInterface interface {
	CreateFloorPlan(floor *entity.Floor, tables []entity.Table) error
	FindFloorByNumber(floorNumber int) (*entity.Floor, error)
	FindFloorByID(floorID uint) (*entity.Floor, error)
	FindAllFloors() ([]entity.Floor, error)
	UpdateTable(tableID uint, req model.UpdateTableRequest) (*entity.Table, error)
	DeleteTable(tableID uint) error
	DeleteFloorAndTables(floorID uint, imagePath string) error
	CountUsersPerTable() (map[uint]int, error)
}
