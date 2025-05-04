package main

import (
	"backend/db"
	"backend/handlers"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()

	// Configurar CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Get("/status", handlers.GetStatuses)
	r.Get("/report/reservations", handlers.GetReservationReport)
	r.Get("/report/contracts", handlers.GetContractReport)
	r.Get("/report/payments", handlers.GetPaymentReport)
	r.Get("/report/fines", handlers.GetFineReport)
	r.Get("/report/maintenance", handlers.GetMaintenanceReport)

	log.Println("🚀 Servidor corriendo en http://localhost:9000")
	http.ListenAndServe(":9000", r)
}
