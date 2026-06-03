package repositories

import (
	"database/sql"
	"fmt"

	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/core/domain"
)

type UserRepositorySQL struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepositorySQL {
	return &UserRepositorySQL{DB: db}
}

func (r *UserRepositorySQL) Create(user *domain.User, hashedPassword string) (*domain.User, error) {
	query := `INSERT INTO users (username, email, password, role)
		VALUES ($1, $2, $3, $4) RETURNING id`
	err := r.DB.QueryRow(query, user.Username, user.Email, hashedPassword, "user").Scan(&user.ID)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %w", err)
	}
	return user, nil
}

func (r *UserRepositorySQL) FindByEmail(email string) (*domain.User, error) {
	query := `SELECT id, username, email, password, role FROM users WHERE email = $1`
	user := &domain.User{}
	var hashedPassword string
	err := r.DB.QueryRow(query, email).Scan(
		&user.ID, &user.Username, &user.Email, &hashedPassword, &user.Role,
	)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword
	return user, nil
}

func (r *UserRepositorySQL) FindByUsername(username string) (*domain.User, error) {
	query := `SELECT id, username, email, role FROM users WHERE username = $1`
	user := &domain.User{}
	err := r.DB.QueryRow(query, username).Scan(
		&user.ID, &user.Username, &user.Email, &user.Role,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}
