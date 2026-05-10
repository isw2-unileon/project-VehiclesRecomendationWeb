package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/adapters/repositories"
	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/services"
)

func main() {
	// --- EXTRA FUNCTIONALITY: DATABASE CONNECTION ---
	// Define connection string for PostgreSQL
	connStr := "host=localhost port=5432 user=postgres password=pass dbname=cars sslmode=disable"

	// Initialize the DB connection using your custom repository
	db, err := repositories.InitDB(connStr)
	if err != nil {
		log.Fatalf("CRITICAL: Failed to connect to database: %v", err)
	}
	defer db.Close()

	// --- EXTRA FUNCTIONALITY: BACKGROUND SIMULATOR ---
	// Create an instance of the simulator with the active DB connection
	simulator := services.ApiSimulator{DB: db}

	// Launch the simulator in a separate goroutine (concurrent thread)
	// This ensures the simulator runs in the background without blocking the API
	go simulator.Start()

	// Basic health check endpoint
	http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "ok", "message": "Vehicles Recommendation API is up and running!"}`))
	})

	port := ":8080"
	fmt.Printf("Server starting on port %s...\n", port)

	// Start the server
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
