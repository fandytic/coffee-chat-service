package entity

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	CustomerID      uint
	PayerCustomerID *uint
	TableID         uint
	NeedType        string
	RecipientID     *uint
	Total           float64 `gorm:"not null"`
	Tax             float64 `gorm:"not null"`
	SubTotal        float64 `gorm:"not null"`
	Status          string  `gorm:"default:'pending';not null"` // Contoh: pending, pending_wishlist, processing, completed
	Notes           string
	OrderItems      []OrderItem
	Customer        Customer
	Table           Table
	Recipient       *Customer `gorm:"foreignKey:RecipientID"`
	Payer           *Customer `gorm:"foreignKey:PayerCustomerID"`
}

type OrderItem struct {
	gorm.Model
	OrderID  uint
	MenuID   uint
	Quantity int     `gorm:"not null"`
	Price    float64 `gorm:"not null"`
	Menu     Menu
}
