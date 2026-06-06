package services_test

import (
	"errors"
	"os"
	"testing"

	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/core/domain"
	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/core/services"
	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister_Success(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)

	expectedUser := &domain.User{
		ID:       1,
		Username: "andres",
		Email:    "andres@email.com",
		Role:     "user",
	}

	mockRepo.On("Create", &domain.User{
		Username: "andres",
		Email:    "andres@email.com",
		Role:     "user",
	}, mock.AnythingOfType("string")).Return(expectedUser, nil)

	svc := services.NewAuthService(mockRepo)
	result, err := svc.Register("andres", "andres@email.com", "password123")

	assert.NoError(t, err)
	assert.Equal(t, expectedUser.Username, result.Username)
	assert.Equal(t, expectedUser.Email, result.Email)
}

func TestRegister_Error(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)

	mockRepo.On("Create", &domain.User{
		Username: "andres",
		Email:    "andres@email.com",
		Role:     "user",
	}, mock.AnythingOfType("string")).Return(nil, errors.New("user already exists"))

	svc := services.NewAuthService(mockRepo)
	result, err := svc.Register("andres", "andres@email.com", "password123")

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestLogin_Success(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	os.Setenv("JWT_SECRET", "test-secret")

	hashedPassword := "$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi"

	mockRepo.On("FindByEmail", "andres@email.com").Return(&domain.User{
		ID:       1,
		Username: "andres",
		Email:    "andres@email.com",
		Password: hashedPassword,
		Role:     "user",
	}, nil)

	svc := services.NewAuthService(mockRepo)
	token, err := svc.Login("andres@email.com", "password")

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	mockRepo.AssertExpectations(t)
}

func TestLogin_UserNotFound(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)

	mockRepo.On("FindByEmail", "noexiste@email.com").Return(nil, errors.New("user not found"))

	svc := services.NewAuthService(mockRepo)
	token, err := svc.Login("noexiste@email.com", "password123")

	assert.Error(t, err)
	assert.Empty(t, token)
	mockRepo.AssertExpectations(t)
}

func TestLogin_WrongPassword(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)

	hashedPassword := "$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi"

	mockRepo.On("FindByEmail", "andres@email.com").Return(&domain.User{
		ID:       1,
		Username: "andres",
		Email:    "andres@email.com",
		Password: hashedPassword,
		Role:     "user",
	}, nil)

	svc := services.NewAuthService(mockRepo)
	token, err := svc.Login("andres@email.com", "wrongpassword")

	assert.Error(t, err)
	assert.Empty(t, token)
	mockRepo.AssertExpectations(t)
}

func TestGetUserByEmail_Success(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)

	expectedUser := &domain.User{
		ID:       1,
		Username: "andres",
		Email:    "andres@email.com",
	}

	mockRepo.On("FindByEmail", "andres@email.com").Return(expectedUser, nil)

	svc := services.NewAuthService(mockRepo)
	result, err := svc.GetUserByEmail("andres@email.com")

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, result)
	mockRepo.AssertExpectations(t)
}

func TestGetUserByEmail_NotFound(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)

	mockRepo.On("FindByEmail", "noexiste@email.com").Return(nil, errors.New("user not found"))

	svc := services.NewAuthService(mockRepo)
	result, err := svc.GetUserByEmail("noexiste@email.com")

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}
