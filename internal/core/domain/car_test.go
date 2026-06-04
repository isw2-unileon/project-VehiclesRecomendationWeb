package domain_test

import (
	"testing"

	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/core/domain"
	"github.com/stretchr/testify/assert"
)

func intPtr(i int) *int { return &i }
func TestCar_FieldTypes(t *testing.T) {
	car := domain.Car{
		ID:           1,
		Brand:        "Toyota",
		Model:        "Corolla",
		Engine:       "1.8L",
		CapacityCC:   intPtr(1800),
		HorsePower:   140,
		TopSpeedKMH:  180,
		Acceleration: 9.5,
		Price:        20000.0,
		FuelType:     "Gasoline",
		Seats:        5,
		TorqueNM:     intPtr(172),
	}

	assert.Equal(t, 1, car.ID)
	assert.Equal(t, "Toyota", car.Brand)
	assert.Equal(t, "Corolla", car.Model)
	assert.Equal(t, "1.8L", car.Engine)
	assert.Equal(t, 1800, *car.CapacityCC)
	assert.Equal(t, 140, car.HorsePower)
	assert.Equal(t, 180, car.TopSpeedKMH)
	assert.Equal(t, 9.5, car.Acceleration)
	assert.Equal(t, 20000.0, car.Price)
	assert.Equal(t, "Gasoline", car.FuelType)
	assert.Equal(t, 5, car.Seats)
	assert.Equal(t, 172, *car.TorqueNM)
}

func TestCar_DefaultValues(t *testing.T) {
	car := domain.Car{}
	assert.Equal(t, 0, car.ID)
	assert.Equal(t, "", car.Brand)
	assert.Equal(t, "", car.Model)
	assert.Equal(t, 0.0, car.Price)
	assert.Equal(t, 0, car.Seats)
	assert.Nil(t, car.CapacityCC)
	assert.Nil(t, car.TorqueNM)
}

func TestCar_PriceIsFloat(t *testing.T) {
	car := domain.Car{Price: 19999.99}
	assert.Equal(t, 19999.99, car.Price)
}

func TestCar_AccelerationIsFloat(t *testing.T) {
	car := domain.Car{Acceleration: 4.2}
	assert.Equal(t, 4.2, car.Acceleration)
}
