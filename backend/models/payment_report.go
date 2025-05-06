package models

import "time"

type PaymentReport struct {
	ID         int       `json:"id"`
	ContractID *int      `json:"rental_contract_id,omitempty"`
	FineID     *int      `json:"fine_id,omitempty"`
	Date       time.Time `json:"payment_date"`
	Amount     float64   `json:"amount"`
	Method     string    `json:"payment_method"`
	StatusName string    `json:"status_name"`
}
