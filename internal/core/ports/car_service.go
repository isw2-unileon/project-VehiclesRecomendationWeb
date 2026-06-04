package ports

import "github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/core/domain"

type CarFilters struct {
	Brand    string
	FuelType string
	MinPrice float64
	MaxPrice float64
	MinSeats int
	MinHP    int
}

type CarService interface {
	GetAllCars() ([]domain.Car, error)
	GetCarByID(id int) (*domain.Car, error)
	SearchCars(filters CarFilters) ([]domain.Car, error)
}
