package repositories

import (
	"database/sql"
	"fmt"

	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/core/domain"
	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/core/ports"
)

type CarRepositorySQL struct {
	DB *sql.DB
}

func NewCarRepository(db *sql.DB) *CarRepositorySQL {
	return &CarRepositorySQL{DB: db}
}

func (r *CarRepositorySQL) FindAll() ([]domain.Car, error) {
	rows, err := r.DB.Query(`SELECT id, company, car_name, engine, capacity_cc,
		power_hp, max_speed_kmh, acceleration_0_100_sec, price, fuel_type, seats, torque_nm
		FROM cars`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanCars(rows)
}

func (r *CarRepositorySQL) FindByID(id int) (*domain.Car, error) {
	row := r.DB.QueryRow(`SELECT id, company, car_name, engine, capacity_cc,
		power_hp, max_speed_kmh, acceleration_0_100_sec, price, fuel_type, seats, torque_nm
		FROM cars WHERE id = $1`, id)

	car := &domain.Car{}
	err := row.Scan(&car.ID, &car.Brand, &car.Model, &car.Engine, &car.CapacityCC,
		&car.HorsePower, &car.TopSpeedKMH, &car.Acceleration, &car.Price,
		&car.FuelType, &car.Seats, &car.TorqueNM)
	if err != nil {
		return nil, err
	}
	return car, nil
}

func (r *CarRepositorySQL) FindByFilters(filters ports.CarFilters) ([]domain.Car, error) {
	query := `SELECT id, company, car_name, engine, capacity_cc,
		power_hp, max_speed_kmh, acceleration_0_100_sec, price, fuel_type, seats, torque_nm
		FROM cars WHERE 1=1`
	args := []interface{}{}
	i := 1

	if filters.Brand != "" {
		query += fmt.Sprintf(" AND LOWER(company) LIKE LOWER($%d)", i)
		args = append(args, "%"+filters.Brand+"%")
		i++
	}
	if filters.FuelType != "" {
		query += fmt.Sprintf(" AND LOWER(fuel_type) = LOWER($%d)", i)
		args = append(args, filters.FuelType)
		i++
	}
	if filters.MinPrice > 0 {
		query += fmt.Sprintf(" AND price >= $%d", i)
		args = append(args, filters.MinPrice)
		i++
	}
	if filters.MaxPrice > 0 {
		query += fmt.Sprintf(" AND price <= $%d", i)
		args = append(args, filters.MaxPrice)
		i++
	}
	if filters.MinSeats > 0 {
		query += fmt.Sprintf(" AND seats >= $%d", i)
		args = append(args, filters.MinSeats)
		i++
	}

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanCars(rows)
}

func scanCars(rows *sql.Rows) ([]domain.Car, error) {
	var cars []domain.Car
	for rows.Next() {
		var car domain.Car
		err := rows.Scan(&car.ID, &car.Brand, &car.Model, &car.Engine, &car.CapacityCC,
			&car.HorsePower, &car.TopSpeedKMH, &car.Acceleration, &car.Price,
			&car.FuelType, &car.Seats, &car.TorqueNM)
		if err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}
	return cars, nil
}
