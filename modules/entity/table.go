package entity

import "gorm.io/gorm"

type Table struct {
	gorm.Model
	TableNumber string `gorm:"not null"`
	TableName   string
	XCoordinate float64 `gorm:"not null"`
	YCoordinate float64 `gorm:"not null"`
	FloorID     uint
}
