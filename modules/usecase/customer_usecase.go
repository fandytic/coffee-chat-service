package usecase

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"coffee-chat-service/modules/entity"
	interfaces "coffee-chat-service/modules/interface"
	"coffee-chat-service/modules/model"
)

type CustomerUseCase struct {
	CustomerRepo interfaces.CustomerRepositoryInterface
	ChatRepo     interfaces.ChatRepositoryInterface
}

func (uc *CustomerUseCase) CheckIn(req model.CustomerCheckInRequest) (*model.CustomerCheckInResponse, error) {
	tableExists, err := uc.CustomerRepo.CheckTableExists(req.TableID)
	if err != nil || !tableExists {
		return nil, errors.New("table not found")
	}

	customer := &entity.Customer{
		Name:     req.Name,
		PhotoURL: req.PhotoURL,
		TableID:  req.TableID,
		Status:   "active",
	}

	if err := uc.CustomerRepo.CreateCustomer(customer); err != nil {
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

func (uc *CustomerUseCase) GetActiveCustomers(loggedInCustomerID uint) (*model.PaginatedActiveCustomersResponse, error) {
	customers, err := uc.CustomerRepo.FindAllActiveExcept(loggedInCustomerID)
	if err != nil {
		return nil, err
	}

	unreadCounts, err := uc.CustomerRepo.CountUnreadMessagesFor(loggedInCustomerID)
	if err != nil {
		return nil, err
	}
	unreadMap := make(map[uint]int)
	for _, result := range unreadCounts {
		unreadMap[result.SenderID] = result.Count
	}

	lastMessages, err := uc.ChatRepo.FindLastMessages(loggedInCustomerID)
	if err != nil {
		return nil, err
	}

	customerResponses := make([]model.ActiveCustomerResponse, 0, len(customers))
	for _, cust := range customers {
		var lastMsg *model.LastMessage
		if msg, ok := lastMessages[cust.ID]; ok {
			lastMsg = &model.LastMessage{
				Text:      msg.Text,
				Timestamp: msg.CreatedAt,
			}
		}

		customerResponses = append(customerResponses, model.ActiveCustomerResponse{
			ID:                  cust.ID,
			Name:                cust.Name,
			PhotoURL:            cust.PhotoURL,
			TableNumber:         cust.Table.TableNumber,
			UnreadMessagesCount: unreadMap[cust.ID],
			LastMessage:         lastMsg,
		})
	}

	finalResponse := &model.PaginatedActiveCustomersResponse{
		Total:     len(customers),
		Customers: customerResponses,
	}

	return finalResponse, nil
}

func (uc *CustomerUseCase) GetAllCustomers(search string) ([]model.AllCustomersResponse, error) {
	customers, err := uc.CustomerRepo.FindAll(search)
	if err != nil {
		return nil, err
	}

	response := make([]model.AllCustomersResponse, 0, len(customers))
	for _, cust := range customers {
		tableNumber := ""
		if cust.Table.TableNumber != "" {
			tableNumber = cust.Table.TableNumber
		}

		response = append(response, model.AllCustomersResponse{
			ID:          cust.ID,
			Name:        cust.Name,
			PhotoURL:    cust.PhotoURL,
			TableNumber: tableNumber,
			Status:      cust.Status,
			LastLogin:   cust.UpdatedAt,
		})
	}

	return response, nil
}
