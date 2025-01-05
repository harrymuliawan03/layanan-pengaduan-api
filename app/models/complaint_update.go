package models

import (
	"gorm.io/gorm"
)

type ComplaintUpdate struct {
	gorm.Model
	ID          uint
	ComplaintID uint
	AdminID     uint
	Status      string
	Note        string
	UpdatedAt   string
}
