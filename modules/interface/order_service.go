package interfaces

import (
	"coffee-chat-service/modules/entity"
	"coffee-chat-service/modules/model"
)

type OrderServiceInterface interface {
	CreateOrder(customerID uint, req model.CreateOrderRequest) (*model.CreateOrderResponse, error)
	GetAllOrders() ([]entity.Order, error)
	GetWishlistDetails(wishlistID uint) (*entity.Order, error)
	AcceptWishlist(wishlistID, payerID uint) (*entity.Order, error)
	GetCustomerOrders(customerID uint) ([]model.OrderHistoryResponse, error)
}
