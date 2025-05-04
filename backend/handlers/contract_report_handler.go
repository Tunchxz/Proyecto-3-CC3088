package handlers

import (
	"backend/db"
	"backend/models"
	"encoding/json"
	"net/http"
)

func GetContractReport(w http.ResponseWriter, r *http.Request) {
	query := `
		SELECT rc.id, rc.reservation_id, rc.start_date, rc.end_date, rc.status_id
		FROM RentalContract rc
		JOIN Reservation r ON rc.reservation_id = r.id
		JOIN Vehicle v ON r.vehicle_id = v.id
		WHERE ($1::DATE IS NULL OR rc.start_date >= $1)
		  AND ($2::DATE IS NULL OR rc.end_date <= $2)
		  AND ($3::INT IS NULL OR rc.status_id = $3)
		  AND ($4::INT IS NULL OR v.facility_id = $4)
		  AND ($5::INT IS NULL OR r.customer_id = $5)
	`

	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")
	statusID := r.URL.Query().Get("status_id")
	facilityID := r.URL.Query().Get("facility_id")
	customerID := r.URL.Query().Get("customer_id")

	rows, err := db.Conn.Query(r.Context(), query, parseDate(startDate), parseDate(endDate), parseInt(statusID), parseInt(facilityID), parseInt(customerID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var report []models.ContractReport
	for rows.Next() {
		var c models.ContractReport
		if err := rows.Scan(&c.ID, &c.ReservationID, &c.StartDate, &c.EndDate, &c.StatusID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		report = append(report, c)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}
