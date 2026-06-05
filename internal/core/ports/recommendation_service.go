package ports

type RecommendationService interface {
	GetAIRecommendation(preferences string, filters CarFilters) (string, error)
}
