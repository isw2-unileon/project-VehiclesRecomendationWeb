package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/core/ports"
)

type RecommendationHandler struct {
	service ports.RecommendationService
}

func NewRecommendationHandler(service ports.RecommendationService) *RecommendationHandler {
	return &RecommendationHandler{service: service}
}

type RecommendationRequest struct {
	Preferences string           `json:"preferences"`
	Filters     ports.CarFilters `json:"filters"`
}

func (h *RecommendationHandler) GetRecommendation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req RecommendationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	recommendation, err := h.service.GetAIRecommendation(req.Preferences, req.Filters)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"recommendation": recommendation})
}
