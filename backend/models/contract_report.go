package models

import "time"

type ContractReport struct {
	ID            int       `json:"id"`
	ReservationID int       `json:"reservation_id"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
	StatusName    string    `json:"status_name"`
}
