package ports

import "github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/core/domain"

type FavoriteRepository interface {
	Add(userID, carID int) error
	Remove(userID, carID int) error
	GetByUserID(userID int) ([]domain.Car, error)
	SearchByText(text string) ([]domain.Car, error)
}

type FavoriteService interface {
	AddFavorite(userID, carID int) error
	RemoveFavorite(userID, carID int) error
	GetFavorites(userID int) ([]domain.Car, error)
	SearchByText(text string) ([]domain.Car, error)
}
