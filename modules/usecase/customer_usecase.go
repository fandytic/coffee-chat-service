package usecase

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"

	"coffee-chat-service/modules/entity"
	"coffee-chat-service/modules/model"
)

type CustomerUseCase struct {
	DB *gorm.DB
}

func (uc *CustomerUseCase) CheckIn(req model.CustomerCheckInRequest) (*model.CustomerCheckInResponse, error) {
	var table entity.Table
	if err := uc.DB.First(&table, req.TableID).Error; err != nil {
		return nil, errors.New("table not found")
	}

	customer := entity.Customer{
		Name:     req.Name,
		PhotoURL: req.PhotoURL,
		TableID:  req.TableID,
		Status:   "active",
	}
	if err := uc.DB.Create(&customer).Error; err != nil {
		return nil, fmt.Errorf("could not create customer: %w", err)
	}

	claims := jwt.MapClaims{
		"customer_id": customer.ID,
		"name":        customer.Name,
		"table_id":    customer.TableID,
		"exp":         time.Now().Add(time.Hour * 8).Unix(), // Sesi berlaku 8 jam
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET_KEY")
	authToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, errors.New("failed to generate session token")
	}

	return &model.CustomerCheckInResponse{
		ID:        customer.ID,
		Name:      customer.Name,
		PhotoURL:  customer.PhotoURL,
		TableID:   customer.TableID,
		AuthToken: authToken,
	}, nil
}

// GetActiveCustomers mengambil semua customer yang berstatus aktif
func (uc *CustomerUseCase) GetActiveCustomers() ([]model.ActiveCustomerResponse, error) {
	var customers []entity.Customer
	err := uc.DB.Preload("Table").Where("status = ?", "active").Find(&customers).Error
	if err != nil {
		return nil, err
	}

	response := make([]model.ActiveCustomerResponse, 0, len(customers))
	for _, cust := range customers {
		response = append(response, model.ActiveCustomerResponse{
			ID:          cust.ID,
			Name:        cust.Name,
			PhotoURL:    cust.PhotoURL,
			TableNumber: cust.Table.TableNumber,
		})
	}

	return response, nil
}
