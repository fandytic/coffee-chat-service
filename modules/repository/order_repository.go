package repository

import (
	"gorm.io/gorm"

	"coffee-chat-service/modules/entity"
)

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (r *OrderRepository) FindWishlistByID(id uint) (*entity.Order, error) {
	var order entity.Order
	err := r.DB.
		Preload("Customer").
		Preload("OrderItems.Menu").
		Where("status = ?", "pending_wishlist").
		First(&order, id).Error
	return &order, err
}

func (r *OrderRepository) UpdateOrder(order *entity.Order) error {
	return r.DB.Save(order).Error
}

func (r *OrderRepository) FindMenusByIDs(menuIDs []uint) (map[uint]entity.Menu, error) {
	var menus []entity.Menu
	if err := r.DB.Where("id IN ?", menuIDs).Find(&menus).Error; err != nil {
		return nil, err
	}
	menuMap := make(map[uint]entity.Menu)
	for _, menu := range menus {
		menuMap[menu.ID] = menu
	}
	return menuMap, nil
}

func (r *OrderRepository) FindCustomerWithTable(customerID uint) (*entity.Customer, error) {
	var customer entity.Customer
	if err := r.DB.Preload("Table").First(&customer, customerID).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *OrderRepository) CreateOrder(order *entity.Order) error {
	return r.DB.Create(order).Error
}

func (r *OrderRepository) FindAll() ([]entity.Order, error) {
	var orders []entity.Order
	err := r.DB.
		Preload("Customer.Table").
		Preload("Recipient.Table").
		Preload("Table").
		Preload("OrderItems.Menu").
		Order("created_at desc").
		Find(&orders).Error
	return orders, err
}

func (r *OrderRepository) FindByID(id uint) (*entity.Order, error) {
	var order entity.Order
	err := r.DB.
		Preload("Customer.Table").
		Preload("Payer.Table"). // Muat juga data si pembayar
		Preload("Recipient.Table").
		Preload("Table").
		Preload("OrderItems.Menu").
		First(&order, id).Error
	return &order, err
}
