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

func (r *OrderRepository) FindMenuPrices(menuIDs []uint) (map[uint]float64, error) {
	var menus []entity.Menu
	if err := r.DB.Where("id IN ?", menuIDs).Find(&menus).Error; err != nil {
		return nil, err
	}
	priceMap := make(map[uint]float64)
	for _, menu := range menus {
		priceMap[menu.ID] = menu.Price
	}
	return priceMap, nil
}

func (r *OrderRepository) CreateOrder(order *entity.Order) error {
	return r.DB.Create(order).Error
}

func (r *OrderRepository) FindAll() ([]entity.Order, error) {
	var orders []entity.Order
	err := r.DB.Preload("Customer").Preload("OrderItems.Menu").Order("created_at desc").Find(&orders).Error
	return orders, err
}

func (r *OrderRepository) FindByID(id uint) (*entity.Order, error) {
	var order entity.Order
	err := r.DB.Preload("Customer").Preload("OrderItems.Menu").First(&order, id).Error
	return &order, err
}
