package services

import (
	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/core/domain"
	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/core/ports"
)

type favoriteService struct {
	repo ports.FavoriteRepository
}

func NewFavoriteService(repo ports.FavoriteRepository) ports.FavoriteService {
	return &favoriteService{repo: repo}
}

func (s *favoriteService) AddFavorite(userID, carID int) error {
	return s.repo.Add(userID, carID)
}

func (s *favoriteService) RemoveFavorite(userID, carID int) error {
	return s.repo.Remove(userID, carID)
}

func (s *favoriteService) GetFavorites(userID int) ([]domain.Car, error) {
	return s.repo.GetByUserID(userID)
}

func (s *favoriteService) SearchByText(text string) ([]domain.Car, error) {
	return s.repo.SearchByText(text)
}
