package repositories

import (
	"database/sql"
	"strconv"

	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/core/domain"
	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/core/ports"
)

type carRepository struct {
	db *sql.DB
}

func NewCarRepository(db *sql.DB) ports.CarRepository {
	return &carRepository{db: db}
}

const selectCars = `SELECT id, company, car_name, engine, capacity_cc,
	power_hp, max_speed_kmh, acceleration_0_100_sec, price, fuel_type, seats, torque_nm FROM cars`

func scanCar(rows *sql.Rows) (domain.Car, error) {
	var c domain.Car
	err := rows.Scan(
		&c.ID, &c.Brand, &c.Model, &c.Engine, &c.CapacityCC,
		&c.HorsePower, &c.TopSpeedKMH, &c.Acceleration, &c.Price,
		&c.FuelType, &c.Seats, &c.TorqueNM,
	)
	return c, err
}

func (r *carRepository) FindAll() ([]domain.Car, error) {
	rows, err := r.db.Query(selectCars)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cars []domain.Car
	for rows.Next() {
		c, err := scanCar(rows)
		if err != nil {
			return nil, err
		}
		cars = append(cars, c)
	}
	return cars, nil
}

func (r *carRepository) FindByID(id int) (*domain.Car, error) {
	query := selectCars + " WHERE id = $1"
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		c, err := scanCar(rows)
		if err != nil {
			return nil, err
		}
		return &c, nil
	}
	return nil, sql.ErrNoRows
}

func (r *carRepository) FindByFilters(filters ports.CarFilters) ([]domain.Car, error) {
	query := selectCars + " WHERE 1=1"
	var args []interface{}
	i := 1

	if filters.Brand != "" {
		query += " AND company ILIKE $" + strconv.Itoa(i)
		args = append(args, "%"+filters.Brand+"%")
		i++
	}
	if filters.FuelType != "" {
		query += " AND fuel_type ILIKE $" + strconv.Itoa(i)
		args = append(args, filters.FuelType)
		i++
	}
	if filters.MaxPrice > 0 {
		query += " AND price <= $" + strconv.Itoa(i)
		args = append(args, filters.MaxPrice)
		i++
	}
	if filters.MinSeats > 0 {
		query += " AND seats >= $" + strconv.Itoa(i)
		args = append(args, filters.MinSeats)
		i++
	}
	if filters.MinHP > 0 {
		query += " AND power_hp >= $" + strconv.Itoa(i)
		args = append(args, filters.MinHP)
		i++
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cars []domain.Car
	for rows.Next() {
		c, err := scanCar(rows)
		if err != nil {
			return nil, err
		}
		cars = append(cars, c)
	}
	return cars, nil
}
