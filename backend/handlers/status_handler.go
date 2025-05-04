package handlers

import (
	"backend/db"
	"backend/models"
	"encoding/json"
	"net/http"
)

func GetStatuses(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Conn.Query(r.Context(), "SELECT id, status_name FROM OperationStatus")
	if err != nil {
		http.Error(w, "Error al obtener estados", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var statuses []models.Status
	for rows.Next() {
		var s models.Status
		if err := rows.Scan(&s.ID, &s.Name); err != nil {
			http.Error(w, "Error al leer filas", http.StatusInternalServerError)
			return
		}
		statuses = append(statuses, s)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(statuses)
}
