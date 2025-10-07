package model

// Model untuk request body saat customer check-in
type CustomerCheckInRequest struct {
	TableID  uint   `json:"table_id"`
	Name     string `json:"name"`
	PhotoURL string `json:"photo_url,omitempty"` // Opsional
}

// Model untuk respons setelah check-in berhasil
type CustomerCheckInResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	PhotoURL  string `json:"photo_url"`
	TableID   uint   `json:"table_id"`
	AuthToken string `json:"auth_token"` // Token JWT untuk sesi customer
}

// Model untuk respons daftar customer aktif
type ActiveCustomerResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	PhotoURL    string `json:"photo_url"`
	TableNumber string `json:"table_number"`
}
