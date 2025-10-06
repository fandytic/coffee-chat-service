package model

// Digunakan untuk unmarshal data meja dari request multipart
type TableData struct {
	TableNumber string `json:"table_number"`
	TableName   string `json:"table_name"`
	XCoordinate int    `json:"x"`
	YCoordinate int    `json:"y"`
}

// Model untuk respons saat mengambil data denah
type FloorPlanResponse struct {
	ID          uint        `json:"id"`
	FloorNumber int         `json:"floor_number"`
	ImageURL    string      `json:"image_url"`
	Tables      []TableData `json:"tables"`
}

// UpdateTableRequest adalah model untuk body request pembaruan meja.
type UpdateTableRequest struct {
	TableName   string `json:"table_name"`
	XCoordinate int    `json:"x"`
	YCoordinate int    `json:"y"`
}

// FloorInfoResponse adalah model respons sederhana untuk daftar lantai.
type FloorInfoResponse struct {
	ID          uint `json:"id"`
	FloorNumber int  `json:"floor_number"`
}
