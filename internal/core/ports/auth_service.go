package ports

import "github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/core/domain"

type AuthService interface {
	Register(username, email, password string) (*domain.User, error)
	Login(email, password string) (string, error)
	GetUserByEmail(email string) (*domain.User, error)
}
