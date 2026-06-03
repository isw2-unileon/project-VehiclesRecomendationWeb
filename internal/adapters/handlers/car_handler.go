package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/core/ports"
)

type CarHandler struct {
	service ports.CarService
}

func NewCarHandler(service ports.CarService) *CarHandler {
	return &CarHandler{service: service}
}

// GetAllCars handles GET /api/cars
func (h *CarHandler) GetAllCars(w http.ResponseWriter, r *http.Request) {
	cars, err := h.service.GetAllCars()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error fetching cars")
		return
	}
	respondWithJSON(w, http.StatusOK, cars)
}

// GetCarByID handles GET /api/cars/{id}
func (h *CarHandler) GetCarByID(w http.ResponseWriter, r *http.Request) {
	// Extract id from URL: /api/cars/42
	parts := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid car ID")
		return
	}

	car, err := h.service.GetCarByID(id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Car not found")
		return
	}
	respondWithJSON(w, http.StatusOK, car)
}

// SearchCars handles GET /api/cars/search?brand=BMW&fuel_type=Gasoline&min_price=10000&max_price=50000
func (h *CarHandler) SearchCars(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	minPrice, _ := strconv.ParseFloat(q.Get("min_price"), 64)
	maxPrice, _ := strconv.ParseFloat(q.Get("max_price"), 64)
	minSeats, _ := strconv.Atoi(q.Get("min_seats"))

	filters := ports.CarFilters{
		Brand:    q.Get("brand"),
		FuelType: q.Get("fuel_type"),
		MinPrice: minPrice,
		MaxPrice: maxPrice,
		MinSeats: minSeats,
	}

	cars, err := h.service.SearchCars(filters)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error searching cars")
		return
	}
	respondWithJSON(w, http.StatusOK, cars)
}

// --- Helpers ---

func respondWithJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func respondWithError(w http.ResponseWriter, status int, message string) {
	respondWithJSON(w, status, map[string]string{"error": message})
}
