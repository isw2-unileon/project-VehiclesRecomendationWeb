package handlers_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/adapters/handlers"
	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/core/domain"
	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/test/mocks"
	"github.com/stretchr/testify/assert"
)

// Token JWT válido con user_id=1 para tests
const validToken = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImFuZHJlcyIsInJvbGUiOiJ1c2VyIn0.signature"

func TestFavoriteHandler_GetFavorites_Success(t *testing.T) {
	mockSvc := new(mocks.MockFavoriteService)

	expectedCars := []domain.Car{
		{ID: 1, Brand: "Toyota", Model: "Corolla"},
	}

	mockSvc.On("GetFavorites", 1).Return(expectedCars, nil)

	handler := handlers.NewFavoriteHandler(mockSvc)

	req := httptest.NewRequest(http.MethodGet, "/api/favorites", nil)
	req.Header.Set("Authorization", validToken)
	rec := httptest.NewRecorder()

	handler.HandleFavorites(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockSvc.AssertExpectations(t)
}

func TestFavoriteHandler_GetFavorites_Unauthorized(t *testing.T) {
	mockSvc := new(mocks.MockFavoriteService)

	handler := handlers.NewFavoriteHandler(mockSvc)

	req := httptest.NewRequest(http.MethodGet, "/api/favorites", nil)
	rec := httptest.NewRecorder()

	handler.HandleFavorites(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestFavoriteHandler_AddFavorite_Success(t *testing.T) {
	mockSvc := new(mocks.MockFavoriteService)

	mockSvc.On("AddFavorite", 1, 10).Return(nil)

	handler := handlers.NewFavoriteHandler(mockSvc)

	body := `{"car_id":10}`
	req := httptest.NewRequest(http.MethodPost, "/api/favorites", bytes.NewBufferString(body))
	req.Header.Set("Authorization", validToken)
	rec := httptest.NewRecorder()

	handler.HandleFavorites(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockSvc.AssertExpectations(t)
}

func TestFavoriteHandler_AddFavorite_Error(t *testing.T) {
	mockSvc := new(mocks.MockFavoriteService)

	mockSvc.On("AddFavorite", 1, 10).Return(errors.New("database error"))

	handler := handlers.NewFavoriteHandler(mockSvc)

	body := `{"car_id":10}`
	req := httptest.NewRequest(http.MethodPost, "/api/favorites", bytes.NewBufferString(body))
	req.Header.Set("Authorization", validToken)
	rec := httptest.NewRecorder()

	handler.HandleFavorites(rec, req)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockSvc.AssertExpectations(t)
}

func TestFavoriteHandler_RemoveFavorite_Success(t *testing.T) {
	mockSvc := new(mocks.MockFavoriteService)

	mockSvc.On("RemoveFavorite", 1, 10).Return(nil)

	handler := handlers.NewFavoriteHandler(mockSvc)

	req := httptest.NewRequest(http.MethodDelete, "/api/favorites?car_id=10", nil)
	req.Header.Set("Authorization", validToken)
	rec := httptest.NewRecorder()

	handler.HandleFavorites(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockSvc.AssertExpectations(t)
}

func TestFavoriteHandler_MethodNotAllowed(t *testing.T) {
	mockSvc := new(mocks.MockFavoriteService)

	handler := handlers.NewFavoriteHandler(mockSvc)

	req := httptest.NewRequest(http.MethodPut, "/api/favorites", nil)
	req.Header.Set("Authorization", validToken)
	rec := httptest.NewRecorder()

	handler.HandleFavorites(rec, req)

	assert.Equal(t, http.StatusMethodNotAllowed, rec.Code)
}

func TestFavoriteHandler_TextSearch_Success(t *testing.T) {
	mockSvc := new(mocks.MockFavoriteService)

	expectedCars := []domain.Car{
		{ID: 1, Brand: "Toyota", Model: "Corolla"},
	}

	mockSvc.On("SearchByText", "toyota").Return(expectedCars, nil)

	handler := handlers.NewFavoriteHandler(mockSvc)

	req := httptest.NewRequest(http.MethodGet, "/api/favorites/search?q=toyota", nil)
	rec := httptest.NewRecorder()

	handler.TextSearch(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockSvc.AssertExpectations(t)
}

func TestFavoriteHandler_TextSearch_EmptyQuery(t *testing.T) {
	mockSvc := new(mocks.MockFavoriteService)

	handler := handlers.NewFavoriteHandler(mockSvc)

	req := httptest.NewRequest(http.MethodGet, "/api/favorites/search", nil)
	rec := httptest.NewRecorder()

	handler.TextSearch(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}
