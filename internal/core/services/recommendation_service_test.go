package services_test

import (
	"errors"
	"testing"

	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/core/domain"
	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/core/ports"
	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/core/services"
	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAIRecommendation_Success(t *testing.T) {
	mockRepo := new(mocks.MockCarRepository)
	mockAI := new(mocks.MockAIProvider)

	filters := ports.CarFilters{}
	expectedCars := []domain.Car{
		{ID: 1, Brand: "Toyota", Model: "Corolla", Price: 20000},
	}

	mockRepo.On("FindByFilters", filters).Return(expectedCars, nil)
	mockAI.On("GenerateRecommendation", mock.AnythingOfType("string"), expectedCars).Return("I recommend the Toyota Corolla", nil)

	svc := services.NewRecommendationService(mockRepo, mockAI)
	result, err := svc.GetAIRecommendation("I want a cheap car", filters)

	assert.NoError(t, err)
	assert.NotEmpty(t, result)
	mockRepo.AssertExpectations(t)
	mockAI.AssertExpectations(t)
}

func TestGetAIRecommendation_RepositoryError(t *testing.T) {
	mockRepo := new(mocks.MockCarRepository)
	mockAI := new(mocks.MockAIProvider)

	filters := ports.CarFilters{}

	mockRepo.On("FindByFilters", filters).Return([]domain.Car{}, errors.New("database error"))

	svc := services.NewRecommendationService(mockRepo, mockAI)
	result, err := svc.GetAIRecommendation("I want a cheap car", filters)

	assert.Error(t, err)
	assert.Empty(t, result)
	mockRepo.AssertExpectations(t)
}

func TestGetAIRecommendation_AIError(t *testing.T) {
	mockRepo := new(mocks.MockCarRepository)
	mockAI := new(mocks.MockAIProvider)

	filters := ports.CarFilters{}
	expectedCars := []domain.Car{
		{ID: 1, Brand: "Toyota", Model: "Corolla", Price: 20000},
	}

	mockRepo.On("FindByFilters", filters).Return(expectedCars, nil)
	mockAI.On("GenerateRecommendation", mock.AnythingOfType("string"), expectedCars).Return("", errors.New("AI service unavailable"))

	svc := services.NewRecommendationService(mockRepo, mockAI)
	result, err := svc.GetAIRecommendation("I want a cheap car", filters)

	assert.Error(t, err)
	assert.Empty(t, result)
	mockRepo.AssertExpectations(t)
	mockAI.AssertExpectations(t)
}

func TestGetAIRecommendation_ParsesMaxPrice(t *testing.T) {
	mockRepo := new(mocks.MockCarRepository)
	mockAI := new(mocks.MockAIProvider)

	filters := ports.CarFilters{}
	expectedCars := []domain.Car{
		{ID: 1, Brand: "Toyota", Model: "Corolla", Price: 20000},
	}

	mockRepo.On("FindByFilters", mock.MatchedBy(func(f ports.CarFilters) bool {
		return f.MaxPrice == 30000
	})).Return(expectedCars, nil)
	mockAI.On("GenerateRecommendation", mock.AnythingOfType("string"), expectedCars).Return("I recommend the Toyota Corolla", nil)

	svc := services.NewRecommendationService(mockRepo, mockAI)
	result, err := svc.GetAIRecommendation("I want a car under 30000", filters)

	assert.NoError(t, err)
	assert.NotEmpty(t, result)
	mockRepo.AssertExpectations(t)
}

func TestGetAIRecommendation_ParsesFuelType(t *testing.T) {
	mockRepo := new(mocks.MockCarRepository)
	mockAI := new(mocks.MockAIProvider)

	filters := ports.CarFilters{}
	expectedCars := []domain.Car{
		{ID: 3, Brand: "Tesla", Model: "Model 3", FuelType: "Electric"},
	}

	mockRepo.On("FindByFilters", mock.MatchedBy(func(f ports.CarFilters) bool {
		return f.FuelType == "Electric"
	})).Return(expectedCars, nil)
	mockAI.On("GenerateRecommendation", mock.AnythingOfType("string"), expectedCars).Return("I recommend the Tesla Model 3", nil)

	svc := services.NewRecommendationService(mockRepo, mockAI)
	result, err := svc.GetAIRecommendation("I want an electric car", filters)

	assert.NoError(t, err)
	assert.NotEmpty(t, result)
	mockRepo.AssertExpectations(t)
}
