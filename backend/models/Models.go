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
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
	Amount      float64 `json:"amount"`
}

type Categories struct {
	ID          int    `json:"primary_key"`
	Name        string `json:"size:200"`
	Description string `json:"unique"`
}
