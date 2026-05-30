package services

import (
	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/core/ports"
)

type RecommendationService struct {
	carRepo    ports.CarRepository
	aiProvider ports.AIProvider
}

func NewRecommendationService(carRepo ports.CarRepository, aiProvider ports.AIProvider) *RecommendationService {
	return &RecommendationService{
		carRepo:    carRepo,
		aiProvider: aiProvider,
	}
}

func (s *RecommendationService) GetAIRecommendation(preferences string, filters ports.CarFilters) (string, error) {
	// 1. Prioritize database efficiency by pre-filtering candidates via SQL
	cars, err := s.carRepo.FindByFilters(filters)
	if err != nil {
		return "", err
	}

	// Limit context to avoid hitting LLM token limits (e.g., top 15 results)
	if len(cars) > 15 {
		cars = cars[:15]
	}

	// 2. Delegate the natural language generation to the AI provider
	return s.aiProvider.GenerateRecommendation(preferences, cars)
}
