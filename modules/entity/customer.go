package entity

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name     string `gorm:"not null"`
	PhotoURL string
	Status   string `gorm:"default:'active';not null"` // Status bisa 'active' atau 'inactive'
	TableID  uint   `gorm:"not null"`
	Table    Table
}
