package interfaces

import "coffee-chat-service/modules/entity"

type OrderRepositoryInterface interface {
	FindMenusByIDs(menuIDs []uint) (map[uint]entity.Menu, error)
	FindCustomerWithTable(customerID uint) (*entity.Customer, error)
	CreateOrder(order *entity.Order) error
	FindAll() ([]entity.Order, error)
	FindByID(id uint) (*entity.Order, error)
}
