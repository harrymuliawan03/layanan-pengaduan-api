package models

import (
	"time"

	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/utils"

	"gorm.io/gorm"
)

type User struct {
	ID            uint    `gorm:"primaryKey;autoIncrement"`
	Name          string  `gorm:"size:255;not null"`
	Email         string  `gorm:"size:255;not null;uniqueIndex"`
	Password      string  `gorm:"size:255;not null"`
	Role          string  `gorm:"size:255;not null"`
	RememberToken *string `gorm:"size:100"`
	Status        int16   `gorm:"not null;default:1"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

func (m *User) BeforeCreate(tx *gorm.DB) error {
	ctx := tx.Statement.Context

	hashedPass, _ := utils.HashPassword(m.Password, ctx)

	m.Password = hashedPass

	return nil

}

func (m *User) BeforeUpdate(tx *gorm.DB) error {
	ctx := tx.Statement.Context

	if m.Password != "" {
		hashedPass, _ := utils.HashPassword(m.Password, ctx)
		tx.Statement.SetColumn("password", hashedPass)
	}

	return nil

}
