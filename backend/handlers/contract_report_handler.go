package handlers

import (
	"backend/db"
	"backend/models"
	"encoding/json"
	"net/http"
)

func GetContractReport(w http.ResponseWriter, r *http.Request) {
	query := `
		SELECT
			rc.id,
			rc.reservation_id,
			rc.start_date,
			rc.end_date,
			os.status_name
		FROM RentalContract rc
		JOIN Reservation r ON rc.reservation_id = r.id
		JOIN OperationStatus os ON rc.status_id = os.id
		WHERE ($1::DATE IS NULL OR rc.start_date >= $1)
		  AND ($2::DATE IS NULL OR rc.end_date <= $2)
		  AND ($3::INT IS NULL OR rc.reservation_id = $3)
		  AND ($4::TEXT IS NULL OR os.status_name ILIKE $4)
	`

	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")
	reservationID := r.URL.Query().Get("reservation_id")
	statusName := r.URL.Query().Get("status_name")

	rows, err := db.Conn.Query(r.Context(), query,
		parseDate(startDate),
		parseDate(endDate),
		parseInt(reservationID),
		likeString(statusName),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var report []models.ContractReport
	for rows.Next() {
		var c models.ContractReport
		if err := rows.Scan(
			&c.ID, &c.ReservationID, &c.StartDate, &c.EndDate, &c.StatusName,
		); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		report = append(report, c)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}
