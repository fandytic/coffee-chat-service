package model

// DashboardStatsResponse adalah model untuk data statistik dasbor.
type DashboardStatsResponse struct {
	TotalTables    int64 `json:"total_tables"`
	OccupiedTables int64 `json:"occupied_tables"`
	EmptyTables    int64 `json:"empty_tables"`
	ActiveUsers    int64 `json:"active_users"`
}
