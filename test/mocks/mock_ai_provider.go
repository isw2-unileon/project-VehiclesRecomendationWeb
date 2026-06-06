package mocks

import (
	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/core/domain"
	"github.com/stretchr/testify/mock"
)

type MockAIProvider struct {
	mock.Mock
}

func (m *MockAIProvider) GenerateRecommendation(userPreferences string, availableCars []domain.Car) (string, error) {
	args := m.Called(userPreferences, availableCars)
	return args.String(0), args.Error(1)
}
