package services_test

import (
	"errors"
	"testing"

	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/core/domain"
	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/core/ports"
	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/core/services"
	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/test/mocks"
	"github.com/stretchr/testify/assert"
)

// ─── GetAllCars ───────────────────────────────────────────────────────────────

func TestGetAllCars_Success(t *testing.T) {
	mockRepo := new(mocks.MockCarRepository)

	expectedCars := []domain.Car{
		{ID: 1, Brand: "Toyota", Model: "Corolla", Price: 20000},
		{ID: 2, Brand: "BMW", Model: "M3", Price: 60000},
	}

	mockRepo.On("FindAll").Return(expectedCars, nil)

	svc := services.NewCarService(mockRepo)
	result, err := svc.GetAllCars()

	assert.NoError(t, err)
	assert.Equal(t, expectedCars, result)
	mockRepo.AssertExpectations(t)
}

func TestGetAllCars_Error(t *testing.T) {
	mockRepo := new(mocks.MockCarRepository)

	mockRepo.On("FindAll").Return([]domain.Car{}, errors.New("database error"))

	svc := services.NewCarService(mockRepo)
	result, err := svc.GetAllCars()

	assert.Error(t, err)
	assert.Empty(t, result)
	mockRepo.AssertExpectations(t)
}

// ─── GetCarByID ───────────────────────────────────────────────────────────────

func TestGetCarByID_Success(t *testing.T) {
	mockRepo := new(mocks.MockCarRepository)

	expectedCar := &domain.Car{ID: 1, Brand: "Toyota", Model: "Corolla", Price: 20000}

	mockRepo.On("FindByID", 1).Return(expectedCar, nil)

	svc := services.NewCarService(mockRepo)
	result, err := svc.GetCarByID(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedCar, result)
	mockRepo.AssertExpectations(t)
}

func TestGetCarByID_NotFound(t *testing.T) {
	mockRepo := new(mocks.MockCarRepository)

	mockRepo.On("FindByID", 99).Return(nil, errors.New("car not found"))

	svc := services.NewCarService(mockRepo)
	result, err := svc.GetCarByID(99)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

// ─── SearchCars ───────────────────────────────────────────────────────────────

func TestSearchCars_Success(t *testing.T) {
	mockRepo := new(mocks.MockCarRepository)

	filters := ports.CarFilters{
		Brand:    "BMW",
		FuelType: "Gasoline",
		MinPrice: 30000,
		MaxPrice: 80000,
		MinSeats: 2,
	}

	expectedCars := []domain.Car{
		{ID: 2, Brand: "BMW", Model: "M3", Price: 60000, FuelType: "Gasoline"},
	}

	mockRepo.On("FindByFilters", filters).Return(expectedCars, nil)

	svc := services.NewCarService(mockRepo)
	result, err := svc.SearchCars(filters)

	assert.NoError(t, err)
	assert.Equal(t, expectedCars, result)
	mockRepo.AssertExpectations(t)
}

func TestSearchCars_Error(t *testing.T) {
	mockRepo := new(mocks.MockCarRepository)

	filters := ports.CarFilters{}

	mockRepo.On("FindByFilters", filters).Return([]domain.Car{}, errors.New("database error"))

	svc := services.NewCarService(mockRepo)
	result, err := svc.SearchCars(filters)

	assert.Error(t, err)
	assert.Empty(t, result)
	mockRepo.AssertExpectations(t)
}

func TestSearchCars_EmptyFilters(t *testing.T) {
	mockRepo := new(mocks.MockCarRepository)

	filters := ports.CarFilters{}

	expectedCars := []domain.Car{
		{ID: 1, Brand: "Toyota", Model: "Corolla", Price: 20000},
		{ID: 2, Brand: "BMW", Model: "M3", Price: 60000},
	}

	mockRepo.On("FindByFilters", filters).Return(expectedCars, nil)

	svc := services.NewCarService(mockRepo)
	result, err := svc.SearchCars(filters)

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	mockRepo.AssertExpectations(t)
}
