package model

type CreateGroupRequest struct {
	Name      string `json:"name"`
	MemberIDs []uint `json:"member_ids"` // IDs customer yang pertama kali di-invite (selain pembuat)
}

type InviteToGroupRequest struct {
	CustomerIDs []uint `json:"customer_ids"`
}

type GroupResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	CreatorID uint   `json:"creator_id"`
}

type GroupMemberResponse struct {
	CustomerID  uint   `json:"customer_id"` // Customer ID
	Name        string `json:"name"`
	PhotoURL    string `json:"photo_url"`
	TableNumber string `json:"table_number"`
	FloorNumber int    `json:"floor_number"`
}
