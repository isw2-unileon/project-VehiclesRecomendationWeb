package mocks

import (
	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/core/ports"
	"github.com/stretchr/testify/mock"
)

type MockRecommendationService struct {
	mock.Mock
}

func (m *MockRecommendationService) GetAIRecommendation(preferences string, filters ports.CarFilters) (string, error) {
	args := m.Called(preferences, filters)
	return args.String(0), args.Error(1)
}
