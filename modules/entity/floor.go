package entity

import "gorm.io/gorm"

type Floor struct {
	gorm.Model
	FloorNumber int     `gorm:"unique;not null"`
	ImageURL    string  `gorm:"not null"`
	Tables      []Table // Relasi one-to-many
}
