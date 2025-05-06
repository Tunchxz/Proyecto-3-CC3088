package models

import "time"

type MaintenanceReport struct {
	ID          int       `json:"id"`
	CarPlate    string    `json:"car_plate"`
	Date        time.Time `json:"maintenance_date"`
	Description string    `json:"description"`
	Cost        float64   `json:"cost"`
	StatusName  string    `json:"status_name"`
}
