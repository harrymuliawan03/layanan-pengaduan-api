package models

import (
	"time"

	"gorm.io/gorm"
)

type AuthVerif struct {
	gorm.Model
	UserID     uint      `gorm:"not null"`
	UniqueCode string    `gorm:"not null"`
	IsUsed     bool      `gorm:"not null;default:false"`
	ExpiryAt   time.Time `gorm:"not null"`
}

func (*AuthVerif) TableName() string {
	return "auth_verifs"
}
