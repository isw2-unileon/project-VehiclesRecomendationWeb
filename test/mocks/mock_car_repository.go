package mocks

import (
	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/core/domain"
	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/core/ports"
	"github.com/stretchr/testify/mock"
)

type MockCarRepository struct {
	mock.Mock
}

func (m *MockCarRepository) FindAll() ([]domain.Car, error) {
	args := m.Called()
	return args.Get(0).([]domain.Car), args.Error(1)
}

func (m *MockCarRepository) FindByID(id int) (*domain.Car, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Car), args.Error(1)
}

func (m *MockCarRepository) FindByFilters(filters ports.CarFilters) ([]domain.Car, error) {
	args := m.Called(filters)
	return args.Get(0).([]domain.Car), args.Error(1)
}
