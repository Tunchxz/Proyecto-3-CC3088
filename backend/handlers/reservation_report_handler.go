package handlers

import (
	"backend/db"
	"backend/models"
	"encoding/json"
	"net/http"
)

func GetReservationReport(w http.ResponseWriter, r *http.Request) {
	query := `
		SELECT r.id, r.customer_id, r.vehicle_id, r.start_date, r.end_date, r.status_id
		FROM Reservation r
		WHERE ($1::DATE IS NULL OR r.start_date >= $1)
		  AND ($2::DATE IS NULL OR r.end_date <= $2)
		  AND ($3::INT IS NULL OR r.status_id = $3)
		  AND ($4::INT IS NULL OR r.customer_id = $4)
		  AND ($5::INT IS NULL OR r.vehicle_id = $5)
	`

	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")
	statusID := r.URL.Query().Get("status_id")
	customerID := r.URL.Query().Get("customer_id")
	vehicleID := r.URL.Query().Get("vehicle_id")

	rows, err := db.Conn.Query(r.Context(), query, parseDate(startDate), parseDate(endDate), parseInt(statusID), parseInt(customerID), parseInt(vehicleID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var report []models.ReservationReport
	for rows.Next() {
		var res models.ReservationReport
		if err := rows.Scan(&res.ID, &res.CustomerID, &res.VehicleID, &res.StartDate, &res.EndDate, &res.StatusID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		report = append(report, res)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}
