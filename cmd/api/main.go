package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/adapters/repositories"
	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/adapters/simulator"
)

func main() {
	connStr := "host=localhost port=5432 user=postgres password=pass dbname=cars sslmode=disable"

	db, err := repositories.InitDB(connStr)
	if err != nil {
		log.Fatalf("CRITICAL: Failed to connect to database: %v", err)
	}
	defer db.Close()

	sim := simulator.ApiSimulator{DB: db}
	go sim.Start()

	http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "ok", "message": "Vehicles Recommendation API is up and running!"}`))
	})

	port := ":8080"
	fmt.Printf("Server starting on port %s...\n", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
