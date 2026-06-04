package handlers_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/adapters/handlers"
	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/core/domain"
	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/core/ports"
	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/test/mocks"
	"github.com/stretchr/testify/assert"
)

// Helper function to create a pointer to an int
func intPtr(i int) *int {
	return &i
}

// ─── GetAllCars ───────────────────────────────────────────────────────────────

func TestHandler_GetAllCars_Success(t *testing.T) {
	mockSvc := new(mocks.MockCarService)

	expectedCars := []domain.Car{
		{ID: 1, Brand: "Toyota", Model: "Corolla", Price: 20000, CapacityCC: intPtr(1600)},
		{ID: 2, Brand: "BMW", Model: "M3", Price: 60000, CapacityCC: intPtr(3000)},
	}

	mockSvc.On("GetAllCars").Return(expectedCars, nil)

	handler := handlers.NewCarHandler(mockSvc)

	req := httptest.NewRequest(http.MethodGet, "/api/cars", nil)
	rec := httptest.NewRecorder()

	handler.GetAllCars(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var result []domain.Car
	err := json.NewDecoder(rec.Body).Decode(&result)
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	mockSvc.AssertExpectations(t)
}

func TestHandler_GetAllCars_Error(t *testing.T) {
	mockSvc := new(mocks.MockCarService)

	mockSvc.On("GetAllCars").Return([]domain.Car{}, errors.New("database error"))

	handler := handlers.NewCarHandler(mockSvc)

	req := httptest.NewRequest(http.MethodGet, "/api/cars", nil)
	rec := httptest.NewRecorder()

	handler.GetAllCars(rec, req)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockSvc.AssertExpectations(t)
}

// ─── GetCarByID ───────────────────────────────────────────────────────────────

func TestHandler_GetCarByID_Success(t *testing.T) {
	mockSvc := new(mocks.MockCarService)

	expectedCar := &domain.Car{ID: 1, Brand: "Toyota", Model: "Corolla", Price: 20000, CapacityCC: intPtr(1600)}

	mockSvc.On("GetCarByID", 1).Return(expectedCar, nil)

	handler := handlers.NewCarHandler(mockSvc)

	req := httptest.NewRequest(http.MethodGet, "/api/cars/1", nil)
	rec := httptest.NewRecorder()

	handler.GetCarByID(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var result domain.Car
	err := json.NewDecoder(rec.Body).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, 1, result.ID)
	mockSvc.AssertExpectations(t)
}

func TestHandler_GetCarByID_InvalidID(t *testing.T) {
	mockSvc := new(mocks.MockCarService)

	handler := handlers.NewCarHandler(mockSvc)

	req := httptest.NewRequest(http.MethodGet, "/api/cars/abc", nil)
	rec := httptest.NewRecorder()

	handler.GetCarByID(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	mockSvc.AssertExpectations(t)
}

func TestHandler_GetCarByID_NotFound(t *testing.T) {
	mockSvc := new(mocks.MockCarService)

	mockSvc.On("GetCarByID", 99).Return(nil, errors.New("car not found"))

	handler := handlers.NewCarHandler(mockSvc)

	req := httptest.NewRequest(http.MethodGet, "/api/cars/99", nil)
	rec := httptest.NewRecorder()

	handler.GetCarByID(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
	mockSvc.AssertExpectations(t)
}

// ─── SearchCars ───────────────────────────────────────────────────────────────

func TestHandler_SearchCars_Success(t *testing.T) {
	mockSvc := new(mocks.MockCarService)

	expectedCars := []domain.Car{
		{ID: 2, Brand: "BMW", Model: "M3", Price: 60000, FuelType: "Gasoline", CapacityCC: intPtr(3000)},
	}

	filters := ports.CarFilters{
		Brand:    "BMW",
		FuelType: "Gasoline",
		MinPrice: 0,
		MaxPrice: 0,
		MinSeats: 0,
		MinHP:    0,
	}

	mockSvc.On("SearchCars", filters).Return(expectedCars, nil)

	handler := handlers.NewCarHandler(mockSvc)
	req := httptest.NewRequest(http.MethodGet, "/api/cars/search?brand=BMW&fuel_type=Gasoline", nil)
	rec := httptest.NewRecorder()

	handler.SearchCars(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockSvc.AssertExpectations(t)
}

func TestHandler_SearchCars_Error(t *testing.T) {
	mockSvc := new(mocks.MockCarService)
	filters := ports.CarFilters{}
	mockSvc.On("SearchCars", filters).Return([]domain.Car{}, errors.New("database error"))

	handler := handlers.NewCarHandler(mockSvc)
	req := httptest.NewRequest(http.MethodGet, "/api/cars/search", nil)
	rec := httptest.NewRecorder()

	handler.SearchCars(rec, req)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}
