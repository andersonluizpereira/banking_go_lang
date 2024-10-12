package models

import "time"

type Transfer struct {
	ID             int       `json:"id"`
	FromAccountNum string    `json:"from_account_num"`
	ToAccountNum   string    `json:"to_account_num"`
	Amount         float64   `json:"amount"`
	Status         string    `json:"status"` // "success" ou "failed"
	CreatedAt      time.Time `json:"created_at"`
}
