package services

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

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
	prefLower := strings.ToLower(preferences)

	cleanPref := strings.ReplaceAll(prefLower, ",", "")
	cleanPref = strings.ReplaceAll(cleanPref, ".", "")

	if filters.MaxPrice == 0 {
		rePrice := regexp.MustCompile(`(?:under|less than|below|max|budget)[^\d]*(\d{3,7})`)
		matches := rePrice.FindStringSubmatch(cleanPref)
		if len(matches) > 1 {
			if price, err := strconv.ParseFloat(matches[1], 64); err == nil {
				filters.MaxPrice = price
			}
		}
	}

	if filters.MinSeats == 0 {
		reSeats := regexp.MustCompile(`(\d)\s*(?:seat|seater|passenger)`)
		matches := reSeats.FindStringSubmatch(prefLower)
		if len(matches) > 1 {
			if seats, err := strconv.Atoi(matches[1]); err == nil {
				filters.MinSeats = seats
			}
		}
	}

	if filters.FuelType == "" {
		if strings.Contains(prefLower, "hybrid") {
			filters.FuelType = "Hybrid"
		} else if strings.Contains(prefLower, "electric") || strings.Contains(prefLower, "ev") {
			filters.FuelType = "Electric"
		} else if strings.Contains(prefLower, "diesel") {
			filters.FuelType = "Diesel"
		} else if strings.Contains(prefLower, "petrol") || strings.Contains(prefLower, "gasoline") {
			filters.FuelType = "Gasoline"
		}
	}

	if filters.Brand == "" {
		brands := []string{"toyota", "nissan", "bmw", "audi", "mercedes", "ford", "kia", "volkswagen", "aston martin", "mazda", "acura", "honda", "hyundai"}
		for _, b := range brands {
			if strings.Contains(prefLower, b) {
				filters.Brand = b
				break
			}
		}
	}

	cars, err := s.carRepo.FindByFilters(filters)
	if err != nil {
		return "", err
	}

	if len(cars) == 0 {
		emptyFilters := ports.CarFilters{}
		cars, _ = s.carRepo.FindByFilters(emptyFilters)
	}

	if len(cars) > 15 {
		cars = cars[:15]
	}

	enrichedPrompt := fmt.Sprintf(`You are an expert vehicle recommendation assistant. Your goal is to suggest cars that best fit the user's lifestyle, budget, and specific requests. 

CRITICAL RULES:
1. You MUST ONLY recommend vehicles that exist in the provided database context. NEVER invent or suggest car models that are not in the provided list.
2. MATCHING VEHICLES: If you find cars in the catalog that meet the user's requirements (price, fuel, seats, etc.), confidently recommend them as the best options. Always include the exact 'Brand' and 'Model' in your text so the user can easily copy and paste it into our search tool.
3. NO MATCHING VEHICLES: ONLY if no cars in the database meet the user's core criteria, you must:
   - Politely inform the user that we don't have exactly what they are looking for.
   - Explain WHICH specific criteria caused the lack of results.
   - Proactively suggest the closest alternative models from the database.

Always answer in the same language the user used to ask the question.

USER'S REQUEST: 
"%s"`, preferences)

	return s.aiProvider.GenerateRecommendation(enrichedPrompt, cars)
}
