package handlers

import (
	"backend/db"
	"backend/models"
	"encoding/json"
	"net/http"
)

func GetMaintenanceReport(w http.ResponseWriter, r *http.Request) {
	query := `
		SELECT
			m.id,
			v.car_plate,
			m.maintenance_date,
			m.description,
			m.cost,
			os.status_name
		FROM Maintenance m
		JOIN Vehicle_Maintenance vm ON m.id = vm.maintenance_id
		JOIN Vehicle v ON vm.vehicle_id = v.id
		JOIN OperationStatus os ON v.status_id = os.id
		WHERE ($1::DATE IS NULL OR m.maintenance_date >= $1)
		  AND ($2::DATE IS NULL OR m.maintenance_date <= $2)
		  AND ($3::NUMERIC IS NULL OR m.cost >= $3)
		  AND ($4::NUMERIC IS NULL OR m.cost <= $4)
		  AND ($5::TEXT IS NULL OR v.car_plate ILIKE $5)
		  AND ($6::TEXT IS NULL OR m.description ILIKE $6)
		  AND ($7::TEXT IS NULL OR os.status_name ILIKE $7)
	`

	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")
	minCost := r.URL.Query().Get("min_cost")
	maxCost := r.URL.Query().Get("max_cost")
	carPlate := r.URL.Query().Get("car_plate")
	description := r.URL.Query().Get("description")
	statusName := r.URL.Query().Get("status_name")

	rows, err := db.Conn.Query(r.Context(), query,
		parseDate(startDate),
		parseDate(endDate),
		parseFloat(minCost),
		parseFloat(maxCost),
		likeString(carPlate),
		likeString(description),
		likeString(statusName),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var report []models.MaintenanceReport
	for rows.Next() {
		var m models.MaintenanceReport
		if err := rows.Scan(
			&m.ID, &m.CarPlate, &m.Date, &m.Description, &m.Cost, &m.StatusName,
		); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		report = append(report, m)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}
