package entity

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	Name        string  `gorm:"not null" json:"name"`
	Description string  `json:"description"`
	Price       float64 `gorm:"not null" json:"price"`
	ImageURL    string  `gorm:"not null" json:"image_url"`
}
