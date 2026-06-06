package handlers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/adapters/handlers"
	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/internal/core/domain"
	"github.com/isw2-unileon/project-VehiclesRecomendationWeb/test/mocks"
	"github.com/stretchr/testify/assert"
)

func TestAuthHandler_Register_Success(t *testing.T) {
	mockSvc := new(mocks.MockAuthService)

	mockSvc.On("Register", "andres", "andres@email.com", "password123").Return(&domain.User{
		ID:       1,
		Username: "andres",
		Email:    "andres@email.com",
		Role:     "user",
	}, nil)

	handler := handlers.NewAuthHandler(mockSvc)

	body := `{"username":"andres","email":"andres@email.com","password":"password123"}`
	req := httptest.NewRequest(http.MethodPost, "/api/auth/register", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	handler.Register(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockSvc.AssertExpectations(t)
}

func TestAuthHandler_Register_MissingFields(t *testing.T) {
	mockSvc := new(mocks.MockAuthService)

	handler := handlers.NewAuthHandler(mockSvc)

	body := `{"username":"","email":"","password":""}`
	req := httptest.NewRequest(http.MethodPost, "/api/auth/register", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	handler.Register(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	mockSvc.AssertExpectations(t)
}

func TestAuthHandler_Register_ServiceError(t *testing.T) {
	mockSvc := new(mocks.MockAuthService)

	mockSvc.On("Register", "andres", "andres@email.com", "password123").Return(nil, errors.New("user already exists"))

	handler := handlers.NewAuthHandler(mockSvc)

	body := `{"username":"andres","email":"andres@email.com","password":"password123"}`
	req := httptest.NewRequest(http.MethodPost, "/api/auth/register", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	handler.Register(rec, req)

	assert.Equal(t, http.StatusConflict, rec.Code)
	mockSvc.AssertExpectations(t)
}

func TestAuthHandler_Register_MethodNotAllowed(t *testing.T) {
	mockSvc := new(mocks.MockAuthService)

	handler := handlers.NewAuthHandler(mockSvc)

	req := httptest.NewRequest(http.MethodGet, "/api/auth/register", nil)
	rec := httptest.NewRecorder()

	handler.Register(rec, req)

	assert.Equal(t, http.StatusMethodNotAllowed, rec.Code)
}

func TestAuthHandler_Login_Success(t *testing.T) {
	mockSvc := new(mocks.MockAuthService)

	mockSvc.On("Login", "andres@email.com", "password123").Return("mocked-jwt-token", nil)
	mockSvc.On("GetUserByEmail", "andres@email.com").Return(&domain.User{
		ID:       1,
		Username: "andres",
		Email:    "andres@email.com",
	}, nil)

	handler := handlers.NewAuthHandler(mockSvc)

	body := `{"email":"andres@email.com","password":"password123"}`
	req := httptest.NewRequest(http.MethodPost, "/api/auth/login", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	handler.Login(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var result map[string]string
	err := json.NewDecoder(rec.Body).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, "mocked-jwt-token", result["token"])
	mockSvc.AssertExpectations(t)
}

func TestAuthHandler_Login_InvalidCredentials(t *testing.T) {
	mockSvc := new(mocks.MockAuthService)

	mockSvc.On("Login", "andres@email.com", "wrongpassword").Return("", errors.New("invalid credentials"))

	handler := handlers.NewAuthHandler(mockSvc)

	body := `{"email":"andres@email.com","password":"wrongpassword"}`
	req := httptest.NewRequest(http.MethodPost, "/api/auth/login", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()

	handler.Login(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	mockSvc.AssertExpectations(t)
}

func TestAuthHandler_Login_MethodNotAllowed(t *testing.T) {
	mockSvc := new(mocks.MockAuthService)

	handler := handlers.NewAuthHandler(mockSvc)

	req := httptest.NewRequest(http.MethodGet, "/api/auth/login", nil)
	rec := httptest.NewRecorder()

	handler.Login(rec, req)

	assert.Equal(t, http.StatusMethodNotAllowed, rec.Code)
}
