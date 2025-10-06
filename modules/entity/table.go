package entity

import "gorm.io/gorm"

type Table struct {
	gorm.Model
	TableNumber string `gorm:"not null"`
	TableName   string
	XCoordinate int  `gorm:"not null"`
	YCoordinate int  `gorm:"not null"`
	FloorID     uint // Foreign Key untuk tabel Floor
}
