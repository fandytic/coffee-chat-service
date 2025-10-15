package interfaces

import (
	"coffee-chat-service/modules/entity"
	"coffee-chat-service/modules/model"
	"time"
)

type UnreadResult struct {
	SenderID uint
	Count    int
}

type CustomerRepositoryInterface interface {
	FindAllActiveExcept(customerID uint, filter model.CustomerFilter) ([]entity.Customer, error)
	CountUnreadMessagesFor(recipientID uint) ([]UnreadResult, error)
	CheckTableExists(tableID uint) (bool, error)
	CreateCustomer(customer *entity.Customer) error
	FindAll(search string) ([]entity.Customer, error)
	FindTableDetailsByID(tableID uint) (*entity.Table, error)
	UpdateStatusForInactiveCustomers(timeout time.Duration) (int64, error)
}
