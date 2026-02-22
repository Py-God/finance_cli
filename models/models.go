package models

import (
	"github.com/google/uuid"
	"time"
)

type Expense struct {
	ID          uuid.UUID `json:"id"`
	Category    string    `json:"category"`
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`
	Time        time.Time `json:"time"`
}

type Category struct {
	Short string `json:"short"`
	Name  string `json:"name"`
}