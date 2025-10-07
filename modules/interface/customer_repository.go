package interfaces

import "coffee-chat-service/modules/entity"

type UnreadResult struct {
	SenderID uint
	Count    int
}

type CustomerRepositoryInterface interface {
	FindAllActiveExcept(customerID uint) ([]entity.Customer, error)
	CountUnreadMessagesFor(recipientID uint) ([]UnreadResult, error)
	CheckTableExists(tableID uint) (bool, error)    // <-- TAMBAHKAN INI
	CreateCustomer(customer *entity.Customer) error // <-- TAMBAHKAN INI
}
