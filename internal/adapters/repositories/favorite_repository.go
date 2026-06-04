package repositories

import (
	"database/sql"

	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/core/domain"
	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/core/ports"
)

type favoriteRepository struct {
	db *sql.DB
}

func NewFavoriteRepository(db *sql.DB) ports.FavoriteRepository {
	return &favoriteRepository{db: db}
}

func (r *favoriteRepository) Add(userID, carID int) error {
	query := `INSERT INTO user_favorites (user_id, car_id) VALUES ($1, $2) ON CONFLICT DO NOTHING`
	_, err := r.db.Exec(query, userID, carID)
	return err
}

func (r *favoriteRepository) Remove(userID, carID int) error {
	query := `DELETE FROM user_favorites WHERE user_id = $1 AND car_id = $2`
	_, err := r.db.Exec(query, userID, carID)
	return err
}

func (r *favoriteRepository) GetByUserID(userID int) ([]domain.Car, error) {
	query := `SELECT c.id, c.company, c.car_name, c.engine, c.capacity_cc,
	                 c.power_hp, c.max_speed_kmh, c.acceleration_0_100_sec, c.price, c.fuel_type, c.seats, c.torque_nm 
	          FROM cars c
	          JOIN user_favorites f ON c.id = f.car_id
	          WHERE f.user_id = $1`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cars []domain.Car
	for rows.Next() {
		var c domain.Car
		err := rows.Scan(
			&c.ID, &c.Brand, &c.Model, &c.Engine, &c.CapacityCC,
			&c.HorsePower, &c.TopSpeedKMH, &c.Acceleration, &c.Price,
			&c.FuelType, &c.Seats, &c.TorqueNM,
		)
		if err != nil {
			return nil, err
		}
		cars = append(cars, c)
	}
	return cars, nil
}

func (r *favoriteRepository) SearchByText(text string) ([]domain.Car, error) {
	query := `SELECT id, company, car_name, engine, capacity_cc,
	                 power_hp, max_speed_kmh, acceleration_0_100_sec, price, fuel_type, seats, torque_nm 
	          FROM cars 
	          WHERE company ILIKE $1 
	             OR car_name ILIKE $1 
	             OR engine ILIKE $1 
	             OR (company || ' ' || car_name) ILIKE $1
	          LIMIT 20`

	rows, err := r.db.Query(query, "%"+text+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cars []domain.Car
	for rows.Next() {
		var c domain.Car
		err := rows.Scan(
			&c.ID, &c.Brand, &c.Model, &c.Engine, &c.CapacityCC,
			&c.HorsePower, &c.TopSpeedKMH, &c.Acceleration, &c.Price,
			&c.FuelType, &c.Seats, &c.TorqueNM,
		)
		if err != nil {
			return nil, err
		}
		cars = append(cars, c)
	}
	return cars, nil
}
