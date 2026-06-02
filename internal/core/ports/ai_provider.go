package ports

import "github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/core/domain"

type AIProvider interface {
	GenerateRecommendation(userPreferences string, availableCars []domain.Car) (string, error)
}
