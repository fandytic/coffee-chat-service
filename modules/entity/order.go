package entity

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	CustomerID uint
	Total      float64 `gorm:"not null"`
	Tax        float64 `gorm:"not null"`
	SubTotal   float64 `gorm:"not null"`
	Status     string  `gorm:"default:'pending';not null"` // pending, processing, completed, cancelled
	Notes      string
	OrderItems []OrderItem
	Customer   Customer
}

type OrderItem struct {
	gorm.Model
	OrderID  uint
	MenuID   uint
	Quantity int     `gorm:"not null"`
	Price    float64 `gorm:"not null"` // Harga menu saat dipesan
	Menu     Menu
}
