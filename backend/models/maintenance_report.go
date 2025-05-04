package models

import "time"

type MaintenanceReport struct {
	ID          int       `json:"id"`
	VehicleID   int       `json:"vehicle_id"`
	Date        time.Time `json:"maintenance_date"`
	Description string    `json:"description"`
	Cost        float64   `json:"cost"`
	StatusID    int       `json:"status_id"`
}
