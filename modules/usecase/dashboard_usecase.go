package usecase

import (
	interfaces "coffee-chat-service/modules/interface"
	"coffee-chat-service/modules/model"
)

type DashboardUseCase struct {
	DashboardRepo interfaces.DashboardRepositoryInterface
}

func (uc *DashboardUseCase) GetStats() (*model.DashboardStatsResponse, error) {
	totalTables, err := uc.DashboardRepo.CountTotalTables()
	if err != nil {
		return nil, err
	}

	activeUsers, err := uc.DashboardRepo.CountActiveUsers()
	if err != nil {
		return nil, err
	}

	occupiedTables, err := uc.DashboardRepo.CountOccupiedTables()
	if err != nil {
		return nil, err
	}

	return &model.DashboardStatsResponse{
		TotalTables:    totalTables,
		OccupiedTables: occupiedTables,
		EmptyTables:    totalTables - occupiedTables,
		ActiveUsers:    activeUsers,
	}, nil
}
