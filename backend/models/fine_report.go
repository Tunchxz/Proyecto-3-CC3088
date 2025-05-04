package models

import "time"

type FineReport struct {
	ID         int       `json:"id"`
	ContractID int       `json:"rental_contract_id"`
	Date       time.Time `json:"fine_date"`
	Amount     float64   `json:"amount"`
	Reason     string    `json:"reason"`
	StatusID   int       `json:"status_id"`
}
