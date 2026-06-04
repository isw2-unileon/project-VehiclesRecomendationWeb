package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/core/ports"
)

type FavoriteHandler struct {
	service ports.FavoriteService
}

func NewFavoriteHandler(service ports.FavoriteService) *FavoriteHandler {
	return &FavoriteHandler{service: service}
}

func extractUserID(r *http.Request) (int, error) {
	authHeader := r.Header.Get("Authorization")
	if len(authHeader) < 7 || !strings.HasPrefix(authHeader, "Bearer ") {
		return 0, fmt.Errorf("unauthorized: missing token")
	}
	tokenString := authHeader[7:]
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return 0, fmt.Errorf("invalid token format")
	}
	payloadSegment := parts[1]
	if l := len(payloadSegment) % 4; l > 0 {
		payloadSegment += strings.Repeat("=", 4-l)
	}
	decoded, err := base64.URLEncoding.DecodeString(payloadSegment)
	if err != nil {
		decoded, err = base64.StdEncoding.DecodeString(payloadSegment)
		if err != nil {
			return 0, fmt.Errorf("failed to decode token payload")
		}
	}
	var claims map[string]interface{}
	if err := json.Unmarshal(decoded, &claims); err != nil {
		return 0, fmt.Errorf("failed to parse token claims")
	}

	for _, key := range []string{"user_id", "id", "sub"} {
		if val, exists := claims[key]; exists {
			if f, ok := val.(float64); ok {
				return int(f), nil
			}
			if s, ok := val.(string); ok {
				if id, err := strconv.Atoi(s); err == nil {
					return id, nil
				}
			}
		}
	}
	return 0, fmt.Errorf("user identification not found in token")
}

func (h *FavoriteHandler) HandleFavorites(w http.ResponseWriter, r *http.Request) {
	userID, err := extractUserID(r)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	switch r.Method {
	case http.MethodPost:
		var req struct {
			CarID int `json:"car_id"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.CarID == 0 {
			log.Printf("[FAVORITOS ERROR] Body inválido recibido")
			respondWithError(w, http.StatusBadRequest, "Invalid request body")
			return
		}
		if err := h.service.AddFavorite(userID, req.CarID); err != nil {
			log.Printf("[FAVORITOS ERROR CRÍTICO] Error de Base de Datos al añadir: %v", err)
			respondWithError(w, http.StatusInternalServerError, "Could not save favorite")
			return
		}
		respondWithJSON(w, http.StatusOK, map[string]string{"status": "success", "message": "Added to favorites"})

	case http.MethodDelete:
		carID, err := strconv.Atoi(r.URL.Query().Get("car_id"))
		if err != nil || carID == 0 {
			respondWithError(w, http.StatusBadRequest, "Invalid car_id param")
			return
		}
		if err := h.service.RemoveFavorite(userID, carID); err != nil {
			respondWithError(w, http.StatusInternalServerError, "Could not remove favorite")
			return
		}
		respondWithJSON(w, http.StatusOK, map[string]string{"status": "success", "message": "Removed from favorites"})

	case http.MethodGet:
		cars, err := h.service.GetFavorites(userID)
		if err != nil {
			log.Printf("[FAVORITOS ERROR] Error al obtener favoritos del usuario %d: %v", userID, err)
			respondWithError(w, http.StatusInternalServerError, "Could not fetch favorites")
			return
		}
		respondWithJSON(w, http.StatusOK, cars)

	default:
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func (h *FavoriteHandler) TextSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		respondWithJSON(w, http.StatusOK, []interface{}{})
		return
	}
	cars, err := h.service.SearchByText(query)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error processing text search")
		return
	}
	respondWithJSON(w, http.StatusOK, cars)
}
