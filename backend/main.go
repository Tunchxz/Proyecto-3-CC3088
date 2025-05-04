package main

import (
	"backend/db"
	"backend/handlers"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	r.Get("/status", handlers.GetStatuses)
	r.Get("/report/reservations", handlers.GetReservationReport)
	r.Get("/report/contracts", handlers.GetContractReport)
	r.Get("/report/payments", handlers.GetPaymentReport)
	r.Get("/report/fines", handlers.GetFineReport)
	r.Get("/report/maintenance", handlers.GetMaintenanceReport)

	log.Println("🚀 Servidor corriendo en http://localhost:9000")
	http.ListenAndServe(":9000", r)
}
