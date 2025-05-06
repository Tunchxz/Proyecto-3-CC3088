package handlers

import (
	"backend/db"
	"backend/models"
	"encoding/json"
	"net/http"
)

func GetFineReport(w http.ResponseWriter, r *http.Request) {
	query := `
		SELECT
			f.id,
			f.rental_contract_id,
			f.fine_date,
			f.amount,
			f.reason,
			os.status_name
		FROM Fine f
		JOIN RentalContract rc ON f.rental_contract_id = rc.id
		JOIN Reservation r ON rc.reservation_id = r.id
		JOIN OperationStatus os ON f.status_id = os.id
		WHERE ($1::DATE IS NULL OR f.fine_date >= $1)
		  AND ($2::DATE IS NULL OR f.fine_date <= $2)
		  AND ($3::NUMERIC IS NULL OR f.amount >= $3)
		  AND ($4::NUMERIC IS NULL OR f.amount <= $4)
		  AND ($5::TEXT IS NULL OR os.status_name ILIKE $5)
		  AND ($6::TEXT IS NULL OR f.reason ILIKE $6)
	`

	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")
	minAmount := r.URL.Query().Get("min_amount")
	maxAmount := r.URL.Query().Get("max_amount")
	statusName := r.URL.Query().Get("status_name")
	keyword := r.URL.Query().Get("keyword")

	rows, err := db.Conn.Query(r.Context(), query,
		parseDate(startDate),
		parseDate(endDate),
		parseFloat(minAmount),
		parseFloat(maxAmount),
		likeString(statusName),
		likeString(keyword),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var report []models.FineReport
	for rows.Next() {
		var f models.FineReport
		if err := rows.Scan(
			&f.ID, &f.ContractID, &f.Date, &f.Amount, &f.Reason, &f.StatusName,
		); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		report = append(report, f)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}
