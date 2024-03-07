package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"primary_key"`
	Name     string    `gorm:"size:100"`
	Email    string    `gorm:"unique"`
	Password string    `gorm:"size:255"`
}

type Products struct {
	ID          int     `json:"primary_key"`
	Name        string  `json:"size:100"`
	Description string  `json:"unique"`
	Price       float64 `json:"price"`
}
