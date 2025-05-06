package handlers

import (
	"backend/db"
	"backend/models"
	"encoding/json"
	"net/http"
)

func GetPaymentReport(w http.ResponseWriter, r *http.Request) {
	query := `
		SELECT
			p.id,
			p.rental_contract_id,
			p.fine_id,
			p.payment_date,
			p.amount,
			p.payment_method,
			os.status_name
		FROM Payment p
		JOIN OperationStatus os ON p.status_id = os.id
		WHERE ($1::DATE IS NULL OR p.payment_date >= $1)
		  AND ($2::DATE IS NULL OR p.payment_date <= $2)
		  AND ($3::TEXT IS NULL OR p.payment_method ILIKE $3)
		  AND ($4::NUMERIC IS NULL OR p.amount >= $4)
		  AND ($5::NUMERIC IS NULL OR p.amount <= $5)
		  AND ($6::TEXT IS NULL OR os.status_name ILIKE $6)
	`

	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")
	method := r.URL.Query().Get("payment_method")
	minAmount := r.URL.Query().Get("min_amount")
	maxAmount := r.URL.Query().Get("max_amount")
	statusName := r.URL.Query().Get("status_name")

	rows, err := db.Conn.Query(r.Context(), query,
		parseDate(startDate),
		parseDate(endDate),
		likeString(method),
		parseFloat(minAmount),
		parseFloat(maxAmount),
		likeString(statusName),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var report []models.PaymentReport
	for rows.Next() {
		var p models.PaymentReport
		if err := rows.Scan(
			&p.ID, &p.ContractID, &p.FineID, &p.Date, &p.Amount, &p.Method, &p.StatusName,
		); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		report = append(report, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}
