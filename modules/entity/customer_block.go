package entity

import "gorm.io/gorm"

type CustomerBlock struct {
	gorm.Model
	BlockerID uint `gorm:"uniqueIndex:idx_blocker_blocked"`
	BlockedID uint `gorm:"uniqueIndex:idx_blocker_blocked"`

	Blocker Customer `gorm:"foreignKey:BlockerID"`
	Blocked Customer `gorm:"foreignKey:BlockedID"`
}
