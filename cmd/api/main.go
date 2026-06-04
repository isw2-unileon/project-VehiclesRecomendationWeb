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

func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next(w, r)
	}
}

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

	// Wiring
	carRepo := repositories.NewCarRepository(db)
	carService := services.NewCarService(carRepo)
	carHandler := handlers.NewCarHandler(carService)

	userRepo := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(authService)

	groqClient := groq.NewGroqClient()
	aiService := services.NewRecommendationService(carRepo, groqClient)
	aiHandler := handlers.NewRecommendationHandler(aiService)

	sim := &simulator.ApiSimulator{DB: db}
	go sim.Start()

	mux := http.NewServeMux()

	mux.HandleFunc("/api/health", enableCORS(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "ok", "message": "Vehicles Recommendation API is up and running!"}`))
	}))

	mux.HandleFunc("/api/cars/", enableCORS(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/cars/search" {
			carHandler.SearchCars(w, r)
			return
		}
		carHandler.GetCarByID(w, r)
	}))
	mux.HandleFunc("/api/cars", enableCORS(carHandler.GetAllCars))

	mux.HandleFunc("/api/recommend", enableCORS(aiHandler.GetRecommendation))

	mux.HandleFunc("/api/auth/register", enableCORS(authHandler.Register))
	mux.HandleFunc("/api/auth/login", enableCORS(authHandler.Login))

	mux.Handle("/", http.FileServer(http.Dir("public/")))

	port := ":8080"
	fmt.Printf("Server starting on port %s...\n", port)

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
