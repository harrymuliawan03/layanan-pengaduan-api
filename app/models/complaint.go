package models

import (
	"gorm.io/gorm"
)

type Complaint struct {
	gorm.Model
	ID          uint
	UserID      uint
	Title       string
	Description string
	Status      string
	Photo       string
	Address     string
}
