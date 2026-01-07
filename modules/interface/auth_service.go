package interfaces

import "coffee-chat-service/modules/model"

type AuthServiceInterface interface {
	Login(req model.LoginRequest) (*model.LoginResponse, error)
	ResetPassword(req model.ResetPasswordRequest) error
	UpdatePassword(adminID uint, req model.UpdatePasswordRequest) error
	UpdateUsername(adminID uint, req model.UpdateUsernameRequest) error
}
