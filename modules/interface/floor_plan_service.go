package interfaces

import (
	"mime/multipart"

	"coffee-chat-service/modules/model"
)

type FloorPlanServiceInterface interface {
	CreateFloorPlan(floorNumber int, fileHeader *multipart.FileHeader, tablesData []model.TableData) (*model.FloorPlanResponse, error)
	GetFloorPlanByNumber(floorNumber int) (*model.FloorPlanResponse, error)
}
