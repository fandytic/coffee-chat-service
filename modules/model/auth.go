package model

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type ResetPasswordRequest struct {
	Username    string `json:"username" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=6"`
}

type UpdatePasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=6"`
}

type UpdateUsernameRequest struct {
	NewUsername string `json:"new_username" validate:"required"`
}
