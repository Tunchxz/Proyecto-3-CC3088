package handlers

import (
	"backend/db"
	"backend/models"
	"encoding/json"
	"net/http"
)

func GetPaymentReport(w http.ResponseWriter, r *http.Request) {
	query := `
		SELECT id, rental_contract_id, fine_id, payment_date, amount, payment_method, status_id
		FROM Payment
		WHERE ($1::DATE IS NULL OR payment_date >= $1)
		  AND ($2::DATE IS NULL OR payment_date <= $2)
		  AND ($3::TEXT IS NULL OR payment_method = $3)
		  AND ($4::NUMERIC IS NULL OR amount >= $4)
		  AND ($5::NUMERIC IS NULL OR amount <= $5)
		  AND ($6::INT IS NULL OR status_id = $6)
		  AND ($7::INT IS NULL OR rental_contract_id = $7)
		  AND ($8::INT IS NULL OR fine_id = $8)
	`

	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")
	method := r.URL.Query().Get("payment_method")
	minAmount := r.URL.Query().Get("min_amount")
	maxAmount := r.URL.Query().Get("max_amount")
	statusID := r.URL.Query().Get("status_id")
	contractID := r.URL.Query().Get("rental_contract_id")
	fineID := r.URL.Query().Get("fine_id")

	rows, err := db.Conn.Query(r.Context(), query,
		parseDate(startDate), parseDate(endDate),
		parseString(method),
		parseFloat(minAmount), parseFloat(maxAmount),
		parseInt(statusID), parseInt(contractID), parseInt(fineID),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var report []models.PaymentReport
	for rows.Next() {
		var p models.PaymentReport
		if err := rows.Scan(&p.ID, &p.ContractID, &p.FineID, &p.Date, &p.Amount, &p.Method, &p.StatusID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		report = append(report, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}
