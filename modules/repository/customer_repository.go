package repository

import (
	"gorm.io/gorm"

	"coffee-chat-service/modules/entity"
	interfaces "coffee-chat-service/modules/interface"
	"coffee-chat-service/modules/model"
)

type CustomerRepository struct {
	DB *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{DB: db}
}

func (r *CustomerRepository) FindAllActiveExcept(customerID uint, filter model.CustomerFilter) ([]entity.Customer, error) {
	var customers []entity.Customer
	query := r.DB.Joins("JOIN tables ON tables.id = customers.table_id").
		Joins("JOIN floors ON floors.id = tables.floor_id").
		Where("customers.status = ? AND customers.id != ?", "active", customerID)

	// Tambahkan filter secara dinamis jika ada
	if filter.Search != "" {
		query = query.Where("customers.name ILIKE ?", "%"+filter.Search+"%")
	}
	if filter.FloorNumber != 0 {
		query = query.Where("floors.floor_number = ?", filter.FloorNumber)
	}
	if filter.TableNumber != "" {
		query = query.Where("tables.table_number = ?", filter.TableNumber)
	}

	// Eksekusi query akhir
	err := query.Preload("Table.Floor").Order("customers.updated_at desc").Find(&customers).Error
	return customers, err
}

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

func (r *CustomerRepository) FindAll(search string) ([]entity.Customer, error) {
	var customers []entity.Customer
	query := r.DB.Preload("Table").Order("updated_at desc")
	if search != "" {
		query = query.Where("name ILIKE ?", "%"+search+"%")
	}
	err := query.Find(&customers).Error
	return customers, err
}

func (r *CustomerRepository) FindTableDetailsByID(tableID uint) (*entity.Table, error) {
	var table entity.Table
	err := r.DB.Preload("Floor").First(&table, tableID).Error
	return &table, err
}
