package usecase

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	interfaces "coffee-chat-service/modules/interface"
	"coffee-chat-service/modules/model"
)

type AuthUseCase struct {
	AdminRepo interfaces.AdminRepositoryInterface
}

func (uc *AuthUseCase) Login(req model.LoginRequest) (*model.LoginResponse, error) {
	admin, err := uc.AdminRepo.FindByUsername(req.Username)
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	claims := jwt.MapClaims{
		"username": admin.Username,
		"user_id":  admin.ID,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET_KEY")
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return &model.LoginResponse{Token: t}, nil
}
