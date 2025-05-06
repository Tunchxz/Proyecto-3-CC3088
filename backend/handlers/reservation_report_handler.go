package handlers

import (
	"backend/db"
	"backend/models"
	"encoding/json"
	"net/http"
)

func GetReservationReport(w http.ResponseWriter, r *http.Request) {
	query := `
		SELECT
			r.id,
			c.id,
			c.first_name,
			c.last_name,
			v.car_plate,
			r.start_date,
			r.end_date,
			os.status_name
		FROM Reservation r
		JOIN Customer c ON r.customer_id = c.id
		JOIN Vehicle v ON r.vehicle_id = v.id
		JOIN OperationStatus os ON r.status_id = os.id
		WHERE ($1::DATE IS NULL OR r.start_date >= $1)
		  AND ($2::DATE IS NULL OR r.end_date <= $2)
		  AND ($3::TEXT IS NULL OR os.status_name ILIKE '%' || $3 || '%')
		  AND ($4::TEXT IS NULL OR c.first_name || ' ' || c.last_name ILIKE '%' || $4 || '%')
		  AND ($5::TEXT IS NULL OR v.car_plate ILIKE '%' || $5 || '%')
	`

	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")
	statusName := r.URL.Query().Get("status_name")
	customerName := r.URL.Query().Get("customer_name")
	carPlate := r.URL.Query().Get("car_plate")

	rows, err := db.Conn.Query(r.Context(), query, parseDate(startDate), parseDate(endDate), nullString(statusName), nullString(customerName), nullString(carPlate))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var report []models.ReservationReport
	for rows.Next() {
		var res models.ReservationReport
		if err := rows.Scan(
			&res.ID,
			&res.CustomerID,
			&res.FirstName,
			&res.LastName,
			&res.CarPlate,
			&res.StartDate,
			&res.EndDate,
			&res.StatusName,
		); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		report = append(report, res)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}

func nullString(s string) interface{} {
	if s == "" {
		return nil
	}
	return s
}
