package repository

import (
	"coffee-chat-service/modules/entity"
	interfaces "coffee-chat-service/modules/interface"

	"gorm.io/gorm"
)

type CustomerRepository struct {
	DB *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{DB: db}
}

// FindAllActiveExcept mengambil semua customer aktif KECUALI diri sendiri.
func (r *CustomerRepository) FindAllActiveExcept(customerID uint) ([]entity.Customer, error) {
	var customers []entity.Customer
	err := r.DB.Preload("Table").
		Where("status = ? AND id != ?", "active", customerID).
		Find(&customers).Error
	return customers, err
}

// CountUnreadMessagesFor menghitung pesan belum dibaca yang ditujukan ke user tertentu.
func (r *CustomerRepository) CountUnreadMessagesFor(recipientID uint) ([]interfaces.UnreadResult, error) {
	var unreadCounts []interfaces.UnreadResult
	err := r.DB.Model(&entity.ChatMessage{}).
		Select("sender_id, count(*) as count").
		Where("recipient_id = ? AND is_read = ?", recipientID, false).
		Group("sender_id").
		Find(&unreadCounts).Error
	return unreadCounts, err
}

func (r *CustomerRepository) CheckTableExists(tableID uint) (bool, error) {
	var count int64
	err := r.DB.Model(&entity.Table{}).Where("id = ?", tableID).Count(&count).Error
	return count > 0, err
}

func (r *CustomerRepository) CreateCustomer(customer *entity.Customer) error {
	return r.DB.Create(customer).Error
}
