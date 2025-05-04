package handlers

import (
	"backend/db"
	"backend/models"
	"encoding/json"
	"net/http"
)

func GetMaintenanceReport(w http.ResponseWriter, r *http.Request) {
	query := `
		SELECT m.id, vm.vehicle_id, m.maintenance_date, m.description, m.cost, v.status_id
		FROM Maintenance m
		JOIN Vehicle_Maintenance vm ON m.id = vm.maintenance_id
		JOIN Vehicle v ON vm.vehicle_id = v.id
		WHERE ($1::DATE IS NULL OR m.maintenance_date >= $1)
		  AND ($2::DATE IS NULL OR m.maintenance_date <= $2)
		  AND ($3::NUMERIC IS NULL OR m.cost >= $3)
		  AND ($4::NUMERIC IS NULL OR m.cost <= $4)
		  AND ($5::INT IS NULL OR vm.vehicle_id = $5)
		  AND ($6::TEXT IS NULL OR m.description ILIKE '%' || $6 || '%')
		  AND ($7::INT IS NULL OR v.status_id = $7)
	`

	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")
	minCost := r.URL.Query().Get("min_cost")
	maxCost := r.URL.Query().Get("max_cost")
	vehicleID := r.URL.Query().Get("vehicle_id")
	description := r.URL.Query().Get("description")
	statusID := r.URL.Query().Get("status_id")

	rows, err := db.Conn.Query(r.Context(), query,
		parseDate(startDate), parseDate(endDate),
		parseFloat(minCost), parseFloat(maxCost),
		parseInt(vehicleID), parseString(description), parseInt(statusID),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var report []models.MaintenanceReport
	for rows.Next() {
		var m models.MaintenanceReport
		if err := rows.Scan(&m.ID, &m.VehicleID, &m.Date, &m.Description, &m.Cost, &m.StatusID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		report = append(report, m)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}
