package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/adapters/groq"
	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/adapters/handlers"
	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/adapters/repositories"
	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/adapters/simulator"
	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/core/services"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	connStr := "host=localhost port=5432 user=postgres password=pass dbname=cars sslmode=disable"

	db, err := repositories.InitDB(connStr)
	if err != nil {
		log.Fatalf("CRITICAL: Failed to connect to database: %v", err)
	}
	defer db.Close()

	// repository → service → handler
	carRepo := repositories.NewCarRepository(db)
	carService := services.NewCarService(carRepo)
	carHandler := handlers.NewCarHandler(carService)

	// components IA (Groq)
	groqClient := groq.NewGroqClient()
	aiService := services.NewRecommendationService(carRepo, groqClient)
	aiHandler := handlers.NewRecommendationHandler(aiService)

	sim := simulator.ApiSimulator{DB: db}
	go sim.Start()

	// check endpoint
	http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "ok", "message": "Vehicles Recommendation API is up and running!"}`))
	})

	http.HandleFunc("/api/cars/search", carHandler.SearchCars) // antes que /api/cars/
	http.HandleFunc("/api/cars/", carHandler.GetCarByID)
	http.HandleFunc("/api/cars", carHandler.GetAllCars)

	// AI recommendation endpoint
	http.HandleFunc("/api/recommend", aiHandler.GetRecommendation)

	port := ":8080"
	fmt.Printf("Server starting on port %s...\n", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
