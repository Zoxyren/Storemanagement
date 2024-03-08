package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"primary_key"`
	UserID   string    `json:"user_id"`
	Name     string    `json:"size:100"`
	Email    string    `json:"unique"`
	Password string    `json:"size:255"`
}

type Products struct {
	ID          int     `json:"primary_key"`
	Name        string  `json:"size:100"`
	Description string  `json:"unique"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
}

type Categories struct {
	ID          int    `json:"primary_key"`
	Name        string `json:"size:200"`
	Description string `json:"unique"`
}
