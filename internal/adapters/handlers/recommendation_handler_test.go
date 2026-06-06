package handlers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/adapters/handlers"
	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/core/ports"
	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/test/mocks"
	"github.com/stretchr/testify/assert"
)

func TestRecommendationHandler_Success(t *testing.T) {
	mockSvc := new(mocks.MockRecommendationService)

	filters := ports.CarFilters{}
	mockSvc.On("GetAIRecommendation", "I want a cheap car", filters).Return("I recommend the Toyota Corolla", nil)

	handler := handlers.NewRecommendationHandler(mockSvc)

	body := `{"preferences":"I want a cheap car","filters":{}}`
	req := httptest.NewRequest(http.MethodPost, "/api/recommendations", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	handler.GetRecommendation(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var result map[string]string
	err := json.NewDecoder(rec.Body).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, "I recommend the Toyota Corolla", result["recommendation"])
	mockSvc.AssertExpectations(t)
}

func TestRecommendationHandler_ServiceError(t *testing.T) {
	mockSvc := new(mocks.MockRecommendationService)

	filters := ports.CarFilters{}
	mockSvc.On("GetAIRecommendation", "I want a cheap car", filters).Return("", errors.New("AI service unavailable"))

	handler := handlers.NewRecommendationHandler(mockSvc)

	body := `{"preferences":"I want a cheap car","filters":{}}`
	req := httptest.NewRequest(http.MethodPost, "/api/recommendations", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	handler.GetRecommendation(rec, req)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockSvc.AssertExpectations(t)
}

func TestRecommendationHandler_MethodNotAllowed(t *testing.T) {
	mockSvc := new(mocks.MockRecommendationService)

	handler := handlers.NewRecommendationHandler(mockSvc)

	req := httptest.NewRequest(http.MethodGet, "/api/recommendations", nil)
	rec := httptest.NewRecorder()

	handler.GetRecommendation(rec, req)

	assert.Equal(t, http.StatusMethodNotAllowed, rec.Code)
}

func TestRecommendationHandler_InvalidBody(t *testing.T) {
	mockSvc := new(mocks.MockRecommendationService)

	handler := handlers.NewRecommendationHandler(mockSvc)

	req := httptest.NewRequest(http.MethodPost, "/api/recommendations", bytes.NewBufferString("invalid json"))
	rec := httptest.NewRecorder()

	handler.GetRecommendation(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
