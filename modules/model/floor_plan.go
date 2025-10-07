package model

type TableData struct {
	ID               uint   `json:"table_id"`
	TableNumber      string `json:"table_number"`
	TableName        string `json:"table_name"`
	XCoordinate      int    `json:"x"`
	YCoordinate      int    `json:"y"`
	ActiveUsersCount int    `json:"active_users_count"`
}

type FloorPlanResponse struct {
	ID          uint        `json:"id"`
	FloorNumber int         `json:"floor_number"`
	ImageURL    string      `json:"image_url"`
	Tables      []TableData `json:"tables"`
}

type UpdateTableRequest struct {
	TableName   string `json:"table_name"`
	XCoordinate int    `json:"x"`
	YCoordinate int    `json:"y"`
}

type FloorInfoResponse struct {
	ID          uint `json:"id"`
	FloorNumber int  `json:"floor_number"`
}

type CreateFloorPlanRequest struct {
	FloorNumber int         `json:"floor_number"`
	ImageURL    string      `json:"image_url"`
	Tables      []TableData `json:"tables"`
}
