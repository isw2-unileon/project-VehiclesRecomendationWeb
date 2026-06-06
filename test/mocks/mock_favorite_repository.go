package mocks

import (
	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/core/domain"
	"github.com/stretchr/testify/mock"
)

type MockFavoriteRepository struct {
	mock.Mock
}

func (m *MockFavoriteRepository) Add(userID, carID int) error {
	args := m.Called(userID, carID)
	return args.Error(0)
}

func (m *MockFavoriteRepository) Remove(userID, carID int) error {
	args := m.Called(userID, carID)
	return args.Error(0)
}

func (m *MockFavoriteRepository) GetByUserID(userID int) ([]domain.Car, error) {
	args := m.Called(userID)
	return args.Get(0).([]domain.Car), args.Error(1)
}

func (m *MockFavoriteRepository) SearchByText(text string) ([]domain.Car, error) {
	args := m.Called(text)
	return args.Get(0).([]domain.Car), args.Error(1)
}
