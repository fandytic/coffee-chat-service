package interfaces

import (
	"coffee-chat-service/modules/model"
	"mime/multipart"
)

type FloorPlanServiceInterface interface {
	CreateFloorPlan(floorNumber int, fileHeader *multipart.FileHeader, tablesData []model.TableData) (*model.FloorPlanResponse, error)
	GetFloorPlanByNumber(floorNumber int) (*model.FloorPlanResponse, error)
}
