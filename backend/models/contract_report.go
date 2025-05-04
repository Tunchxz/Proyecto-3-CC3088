package models

import "time"

type ContractReport struct {
	ID            int       `json:"id"`
	ReservationID int       `json:"reservation_id"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
	StatusID      int       `json:"status_id"`
}
