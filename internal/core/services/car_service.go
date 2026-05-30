package services

import (
	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/core/domain"
	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/core/ports"
)

// CarService contains the business logic for cars
type CarService struct {
	repo ports.CarRepository
}

// NewCarService creates a new CarService with the given repository
func NewCarService(repo ports.CarRepository) *CarService {
	return &CarService{repo: repo}
}

func (s *CarService) GetAllCars() ([]domain.Car, error) {
	return s.repo.FindAll()
}

func (s *CarService) GetCarByID(id int) (*domain.Car, error) {
	return s.repo.FindByID(id)
}

func (s *CarService) SearchCars(filters ports.CarFilters) ([]domain.Car, error) {
	return s.repo.FindByFilters(filters)
}
