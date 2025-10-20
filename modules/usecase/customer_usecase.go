package usecase

import (
	"errors"
	"fmt"
	"log"
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
	OrderRepo    interfaces.OrderRepositoryInterface
}

func (uc *CustomerUseCase) CheckIn(req model.CustomerCheckInRequest) (*model.CustomerCheckInResponse, error) {
	table, err := uc.CustomerRepo.FindTableDetailsByID(req.TableID)
	if err != nil {
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
		"exp":         time.Now().Add(time.Hour * 8).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET_KEY")
	authToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, errors.New("failed to generate session token")
	}

	return &model.CustomerCheckInResponse{
		ID:          customer.ID,
		Name:        customer.Name,
		PhotoURL:    customer.PhotoURL,
		TableID:     customer.TableID,
		TableNumber: table.TableNumber,
		FloorNumber: table.Floor.FloorNumber,
		AuthToken:   authToken,
	}, nil
}

func (uc *CustomerUseCase) GetActiveCustomers(loggedInCustomerID uint, filter model.CustomerFilter) (*model.PaginatedActiveCustomersResponse, error) {
	customers, err := uc.CustomerRepo.FindAllActiveExcept(loggedInCustomerID, filter)
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

	activeWishlists, err := uc.OrderRepo.FindActiveWishlistsByCustomerID()
	if err != nil {
		// Log error tapi jangan hentikan proses, anggap saja tidak ada wishlist
		log.Printf("Warning: could not retrieve active wishlists: %v", err)
		activeWishlists = make(map[uint]uint)
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

		var wishlistID *uint
		if id, ok := activeWishlists[cust.ID]; ok {
			wishlistID = &id
		}

		customerResponses = append(customerResponses, model.ActiveCustomerResponse{
			ID:                  cust.ID,
			Name:                cust.Name,
			PhotoURL:            cust.PhotoURL,
			TableNumber:         cust.Table.TableNumber,
			FloorNumber:         cust.Table.Floor.FloorNumber,
			UnreadMessagesCount: unreadMap[cust.ID],
			LastMessage:         lastMsg,
			WishlistID:          wishlistID,
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

func (uc *CustomerUseCase) CleanUpInactiveCustomers() {
	timeout := 8 * time.Hour
	rowsAffected, err := uc.CustomerRepo.UpdateStatusForInactiveCustomers(timeout)
	if err != nil {
		log.Printf("ERROR: Failed to run cleanup for inactive customers: %v", err)
		return
	}
	if rowsAffected > 0 {
		log.Printf("INFO: Cleaned up %d inactive customer session(s)", rowsAffected)
	}
}

func (uc *CustomerUseCase) RevokeCustomerAccess(customerID uint) error {
	return uc.CustomerRepo.UpdateStatus(customerID, "revoked")
}
