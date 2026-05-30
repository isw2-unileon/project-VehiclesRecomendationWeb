package ports

import "github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/core/domain"

type CarRepository interface {
	FindAll() ([]domain.Car, error)
	FindByID(id int) (*domain.Car, error)
	FindByFilters(filters CarFilters) ([]domain.Car, error)
}

type CarFilters struct {
	Brand    string
	FuelType string
	MinPrice float64
	MaxPrice float64
	MinSeats int
}
