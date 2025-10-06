package interfaces

import "coffee-chat-service/modules/model"

type AuthServiceInterface interface {
	Login(req model.LoginRequest) (*model.LoginResponse, error)
}
