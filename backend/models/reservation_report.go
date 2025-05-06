package models

import "time"

type ReservationReport struct {
	ID         int       `json:"id"`
	CustomerID int       `json:"customer_id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	CarPlate   string    `json:"car_plate"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
	StatusName string    `json:"status_name"`
}
