package interfaces

type DashboardRepositoryInterface interface {
	CountTotalTables() (int64, error)
	CountActiveUsers() (int64, error)
	CountOccupiedTables() (int64, error)
}
