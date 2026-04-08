package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
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
