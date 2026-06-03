package ports

import "github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/core/domain"

type UserRepository interface {
	Create(user *domain.User, hashedPassword string) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	FindByUsername(username string) (*domain.User, error)
}
