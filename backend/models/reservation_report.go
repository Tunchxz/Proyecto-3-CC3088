package models

import "time"

type ReservationReport struct {
	ID         int       `json:"id"`
	CustomerID int       `json:"customer_id"`
	VehicleID  int       `json:"vehicle_id"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
	StatusID   int       `json:"status_id"`
}
