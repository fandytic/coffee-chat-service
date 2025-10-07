package interfaces

import "coffee-chat-service/modules/model"

type DashboardServiceInterface interface {
	GetStats() (*model.DashboardStatsResponse, error)
}
