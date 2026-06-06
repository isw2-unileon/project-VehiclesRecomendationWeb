package services_test

import (
	"errors"
	"testing"

	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/core/domain"
	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/core/services"
	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/test/mocks"
	"github.com/stretchr/testify/assert"
)

func TestAddFavorite_Success(t *testing.T) {
	mockRepo := new(mocks.MockFavoriteRepository)
	mockRepo.On("Add", 1, 10).Return(nil)

	svc := services.NewFavoriteService(mockRepo)
	err := svc.AddFavorite(1, 10)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestAddFavorite_Error(t *testing.T) {
	mockRepo := new(mocks.MockFavoriteRepository)
	mockRepo.On("Add", 1, 10).Return(errors.New("database error"))

	svc := services.NewFavoriteService(mockRepo)
	err := svc.AddFavorite(1, 10)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestRemoveFavorite_Success(t *testing.T) {
	mockRepo := new(mocks.MockFavoriteRepository)
	mockRepo.On("Remove", 1, 10).Return(nil)

	svc := services.NewFavoriteService(mockRepo)
	err := svc.RemoveFavorite(1, 10)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestRemoveFavorite_Error(t *testing.T) {
	mockRepo := new(mocks.MockFavoriteRepository)
	mockRepo.On("Remove", 1, 10).Return(errors.New("database error"))

	svc := services.NewFavoriteService(mockRepo)
	err := svc.RemoveFavorite(1, 10)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetFavorites_Success(t *testing.T) {
	mockRepo := new(mocks.MockFavoriteRepository)

	expectedCars := []domain.Car{
		{ID: 1, Brand: "Toyota", Model: "Corolla"},
		{ID: 2, Brand: "BMW", Model: "M3"},
	}

	mockRepo.On("GetByUserID", 1).Return(expectedCars, nil)

	svc := services.NewFavoriteService(mockRepo)
	result, err := svc.GetFavorites(1)

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	mockRepo.AssertExpectations(t)
}

func TestGetFavorites_Error(t *testing.T) {
	mockRepo := new(mocks.MockFavoriteRepository)
	mockRepo.On("GetByUserID", 1).Return([]domain.Car{}, errors.New("database error"))

	svc := services.NewFavoriteService(mockRepo)
	result, err := svc.GetFavorites(1)

	assert.Error(t, err)
	assert.Empty(t, result)
	mockRepo.AssertExpectations(t)
}

func TestSearchByText_Success(t *testing.T) {
	mockRepo := new(mocks.MockFavoriteRepository)

	expectedCars := []domain.Car{
		{ID: 1, Brand: "Toyota", Model: "Corolla"},
	}

	mockRepo.On("SearchByText", "toyota").Return(expectedCars, nil)

	svc := services.NewFavoriteService(mockRepo)
	result, err := svc.SearchByText("toyota")

	assert.NoError(t, err)
	assert.Len(t, result, 1)
	mockRepo.AssertExpectations(t)
}

func TestSearchByText_Error(t *testing.T) {
	mockRepo := new(mocks.MockFavoriteRepository)
	mockRepo.On("SearchByText", "toyota").Return([]domain.Car{}, errors.New("database error"))

	svc := services.NewFavoriteService(mockRepo)
	result, err := svc.SearchByText("toyota")

	assert.Error(t, err)
	assert.Empty(t, result)
	mockRepo.AssertExpectations(t)
}
