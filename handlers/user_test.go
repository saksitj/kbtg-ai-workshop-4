package handlers

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"
	"workshop_4/models"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository is a mock implementation of UserRepository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetAll() ([]models.User, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.User), args.Error(1)
}

func (m *MockUserRepository) GetByID(id int) (*models.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) Create(req models.CreateUserRequest) (*models.User, error) {
	args := m.Called(req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) Update(id int, req models.UpdateUserRequest) (*models.User, error) {
	args := m.Called(id, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func setupTestApp() *fiber.App {
	app := fiber.New()
	return app
}

func TestGetUsers_Success(t *testing.T) {
	// Setup
	app := setupTestApp()
	mockRepo := new(MockUserRepository)
	userRepo = mockRepo

	users := []models.User{
		{
			ID:           1,
			FirstName:    "John",
			LastName:     "Doe",
			Email:        "john@example.com",
			MemberLevel:  "Gold",
			PointBalance: 100,
		},
		{
			ID:           2,
			FirstName:    "Jane",
			LastName:     "Smith",
			Email:        "jane@example.com",
			MemberLevel:  "Silver",
			PointBalance: 50,
		},
	}

	mockRepo.On("GetAll").Return(users, nil)

	app.Get("/users", GetUsers)

	// Test
	req := httptest.NewRequest("GET", "/users", nil)
	resp, err := app.Test(req)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	var response map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&response)

	assert.True(t, response["success"].(bool))
	assert.NotNil(t, response["data"])
	mockRepo.AssertExpectations(t)
}

func TestGetUser_Success(t *testing.T) {
	// Setup
	app := setupTestApp()
	mockRepo := new(MockUserRepository)
	userRepo = mockRepo

	user := &models.User{
		ID:           1,
		FirstName:    "John",
		LastName:     "Doe",
		Email:        "john@example.com",
		MemberLevel:  "Gold",
		PointBalance: 100,
	}

	mockRepo.On("GetByID", 1).Return(user, nil)

	app.Get("/users/:id", GetUser)

	// Test
	req := httptest.NewRequest("GET", "/users/1", nil)
	resp, err := app.Test(req)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	var response map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&response)

	assert.True(t, response["success"].(bool))
	assert.NotNil(t, response["data"])
	mockRepo.AssertExpectations(t)
}

func TestGetUser_NotFound(t *testing.T) {
	// Setup
	app := setupTestApp()
	mockRepo := new(MockUserRepository)
	userRepo = mockRepo

	mockRepo.On("GetByID", 999).Return((*models.User)(nil), nil)

	app.Get("/users/:id", GetUser)

	// Test
	req := httptest.NewRequest("GET", "/users/999", nil)
	resp, err := app.Test(req)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 404, resp.StatusCode)

	var response map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&response)

	assert.False(t, response["success"].(bool))
	assert.Equal(t, "User not found", response["error"])
	mockRepo.AssertExpectations(t)
}

func TestGetUser_InvalidID(t *testing.T) {
	// Setup
	app := setupTestApp()
	mockRepo := new(MockUserRepository)
	userRepo = mockRepo

	app.Get("/users/:id", GetUser)

	// Test
	req := httptest.NewRequest("GET", "/users/invalid", nil)
	resp, err := app.Test(req)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode)

	var response map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&response)

	assert.False(t, response["success"].(bool))
	assert.Equal(t, "Invalid user ID", response["error"])
}

func TestCreateUser_Success(t *testing.T) {
	// Setup
	app := setupTestApp()
	mockRepo := new(MockUserRepository)
	userRepo = mockRepo

	createReq := models.CreateUserRequest{
		FirstName:    "John",
		LastName:     "Doe",
		Email:        "john@example.com",
		Phone:        "0812345678",
		MemberLevel:  "Bronze",
		PointBalance: 0,
	}

	createdUser := &models.User{
		ID:           1,
		FirstName:    "John",
		LastName:     "Doe",
		Email:        "john@example.com",
		Phone:        "0812345678",
		MemberLevel:  "Bronze",
		PointBalance: 0,
	}

	mockRepo.On("Create", mock.AnythingOfType("models.CreateUserRequest")).Return(createdUser, nil)

	app.Post("/users", CreateUser)

	// Test
	body, _ := json.Marshal(createReq)
	req := httptest.NewRequest("POST", "/users", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 201, resp.StatusCode)

	var response map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&response)

	assert.True(t, response["success"].(bool))
	assert.NotNil(t, response["data"])
	mockRepo.AssertExpectations(t)
}

func TestCreateUser_MissingFields(t *testing.T) {
	// Setup
	app := setupTestApp()
	mockRepo := new(MockUserRepository)
	userRepo = mockRepo

	app.Post("/users", CreateUser)

	// Test with missing required fields
	createReq := models.CreateUserRequest{
		Email: "john@example.com",
		// Missing FirstName and LastName
	}

	body, _ := json.Marshal(createReq)
	req := httptest.NewRequest("POST", "/users", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode)

	var response map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&response)

	assert.False(t, response["success"].(bool))
	assert.Contains(t, response["error"], "required")
}

func TestUpdateUser_Success(t *testing.T) {
	// Setup
	app := setupTestApp()
	mockRepo := new(MockUserRepository)
	userRepo = mockRepo

	existingUser := &models.User{
		ID:           1,
		FirstName:    "John",
		LastName:     "Doe",
		Email:        "john@example.com",
		MemberLevel:  "Bronze",
		PointBalance: 0,
	}

	updateReq := models.UpdateUserRequest{
		FirstName:    "John",
		LastName:     "Updated",
		Email:        "john.updated@example.com",
		MemberLevel:  "Gold",
		PointBalance: 500,
	}

	updatedUser := &models.User{
		ID:           1,
		FirstName:    "John",
		LastName:     "Updated",
		Email:        "john.updated@example.com",
		MemberLevel:  "Gold",
		PointBalance: 500,
	}

	mockRepo.On("GetByID", 1).Return(existingUser, nil)
	mockRepo.On("Update", 1, mock.AnythingOfType("models.UpdateUserRequest")).Return(updatedUser, nil)

	app.Put("/users/:id", UpdateUser)

	// Test
	body, _ := json.Marshal(updateReq)
	req := httptest.NewRequest("PUT", "/users/1", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	var response map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&response)

	assert.True(t, response["success"].(bool))
	assert.NotNil(t, response["data"])
	mockRepo.AssertExpectations(t)
}

func TestUpdateUser_NotFound(t *testing.T) {
	// Setup
	app := setupTestApp()
	mockRepo := new(MockUserRepository)
	userRepo = mockRepo

	mockRepo.On("GetByID", 999).Return((*models.User)(nil), nil)

	app.Put("/users/:id", UpdateUser)

	updateReq := models.UpdateUserRequest{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
	}

	// Test
	body, _ := json.Marshal(updateReq)
	req := httptest.NewRequest("PUT", "/users/999", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 404, resp.StatusCode)

	var response map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&response)

	assert.False(t, response["success"].(bool))
	assert.Equal(t, "User not found", response["error"])
	mockRepo.AssertExpectations(t)
}

func TestDeleteUser_Success(t *testing.T) {
	// Setup
	app := setupTestApp()
	mockRepo := new(MockUserRepository)
	userRepo = mockRepo

	existingUser := &models.User{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
	}

	mockRepo.On("GetByID", 1).Return(existingUser, nil)
	mockRepo.On("Delete", 1).Return(nil)

	app.Delete("/users/:id", DeleteUser)

	// Test
	req := httptest.NewRequest("DELETE", "/users/1", nil)
	resp, err := app.Test(req)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	var response map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&response)

	assert.True(t, response["success"].(bool))
	assert.Equal(t, "User deleted successfully", response["message"])
	mockRepo.AssertExpectations(t)
}

func TestDeleteUser_NotFound(t *testing.T) {
	// Setup
	app := setupTestApp()
	mockRepo := new(MockUserRepository)
	userRepo = mockRepo

	mockRepo.On("GetByID", 999).Return((*models.User)(nil), nil)

	app.Delete("/users/:id", DeleteUser)

	// Test
	req := httptest.NewRequest("DELETE", "/users/999", nil)
	resp, err := app.Test(req)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 404, resp.StatusCode)

	var response map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&response)

	assert.False(t, response["success"].(bool))
	assert.Equal(t, "User not found", response["error"])
	mockRepo.AssertExpectations(t)
}
