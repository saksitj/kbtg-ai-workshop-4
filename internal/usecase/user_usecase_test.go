package usecase

import (
	"errors"
	"testing"
	"time"
	"workshop_4/internal/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository is a mock implementation of domain.UserRepository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindAll() ([]*domain.User, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.User), args.Error(1)
}

func (m *MockUserRepository) FindByID(id int) (*domain.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) FindByEmail(email string) (*domain.User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) Create(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) Update(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestGetAllUsers_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	useCase := NewUserUseCase(mockRepo)

	now := time.Now()
	expectedUsers := []*domain.User{
		{
			ID:           1,
			FirstName:    "John",
			LastName:     "Doe",
			Email:        "john@example.com",
			MemberLevel:  "Gold",
			PointBalance: 1000,
			CreatedAt:    now,
			UpdatedAt:    now,
		},
	}

	mockRepo.On("FindAll").Return(expectedUsers, nil)

	users, err := useCase.GetAllUsers()

	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, users)
	mockRepo.AssertExpectations(t)
}

func TestGetAllUsers_Error(t *testing.T) {
	mockRepo := new(MockUserRepository)
	useCase := NewUserUseCase(mockRepo)

	expectedError := errors.New("database error")
	mockRepo.On("FindAll").Return(nil, expectedError)

	users, err := useCase.GetAllUsers()

	assert.Error(t, err)
	assert.Nil(t, users)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

func TestGetUserByID_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	useCase := NewUserUseCase(mockRepo)

	now := time.Now()
	expectedUser := &domain.User{
		ID:           1,
		FirstName:    "John",
		LastName:     "Doe",
		Email:        "john@example.com",
		MemberLevel:  "Gold",
		PointBalance: 1000,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	mockRepo.On("FindByID", 1).Return(expectedUser, nil)

	user, err := useCase.GetUserByID(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
	mockRepo.AssertExpectations(t)
}

func TestGetUserByID_NotFound(t *testing.T) {
	mockRepo := new(MockUserRepository)
	useCase := NewUserUseCase(mockRepo)

	mockRepo.On("FindByID", 999).Return(nil, domain.ErrUserNotFound)

	user, err := useCase.GetUserByID(999)

	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, domain.ErrUserNotFound, err)
	mockRepo.AssertExpectations(t)
}

func TestCreateUser_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	useCase := NewUserUseCase(mockRepo)

	input := CreateUserInput{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		Phone:     "1234567890",
		Address:   "123 Main St",
	}

	mockRepo.On("FindByEmail", input.Email).Return(nil, nil)
	mockRepo.On("Create", mock.AnythingOfType("*domain.User")).Return(nil).Run(func(args mock.Arguments) {
		user := args.Get(0).(*domain.User)
		user.ID = 1
	})

	user, err := useCase.CreateUser(input)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, input.FirstName, user.FirstName)
	assert.Equal(t, input.LastName, user.LastName)
	assert.Equal(t, input.Email, user.Email)
	assert.Equal(t, "Bronze", user.MemberLevel)
	assert.Equal(t, 0, user.PointBalance)
	mockRepo.AssertExpectations(t)
}

func TestCreateUser_MissingFirstName(t *testing.T) {
	mockRepo := new(MockUserRepository)
	useCase := NewUserUseCase(mockRepo)

	input := CreateUserInput{
		LastName: "Doe",
		Email:    "john@example.com",
	}

	user, err := useCase.CreateUser(input)

	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, domain.ErrFirstNameRequired, err)
}

func TestCreateUser_MissingLastName(t *testing.T) {
	mockRepo := new(MockUserRepository)
	useCase := NewUserUseCase(mockRepo)

	input := CreateUserInput{
		FirstName: "John",
		Email:     "john@example.com",
	}

	user, err := useCase.CreateUser(input)

	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, domain.ErrLastNameRequired, err)
}

func TestCreateUser_MissingEmail(t *testing.T) {
	mockRepo := new(MockUserRepository)
	useCase := NewUserUseCase(mockRepo)

	input := CreateUserInput{
		FirstName: "John",
		LastName:  "Doe",
	}

	user, err := useCase.CreateUser(input)

	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, domain.ErrEmailRequired, err)
}

func TestCreateUser_DuplicateEmail(t *testing.T) {
	mockRepo := new(MockUserRepository)
	useCase := NewUserUseCase(mockRepo)

	input := CreateUserInput{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "existing@example.com",
	}

	existingUser := &domain.User{
		ID:    1,
		Email: input.Email,
	}

	mockRepo.On("FindByEmail", input.Email).Return(existingUser, nil)

	user, err := useCase.CreateUser(input)

	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, domain.ErrDuplicateEmail, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateUser_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	useCase := NewUserUseCase(mockRepo)

	now := time.Now()
	existingUser := &domain.User{
		ID:           1,
		FirstName:    "John",
		LastName:     "Doe",
		Email:        "john@example.com",
		MemberLevel:  "Gold",
		PointBalance: 1000,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	input := UpdateUserInput{
		FirstName:    "Jane",
		LastName:     "Smith",
		Email:        "jane@example.com",
		Phone:        "9876543210",
		MemberLevel:  "Platinum",
		PointBalance: 2000,
	}

	mockRepo.On("FindByID", 1).Return(existingUser, nil)
	mockRepo.On("Update", mock.AnythingOfType("*domain.User")).Return(nil)

	user, err := useCase.UpdateUser(1, input)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, input.FirstName, user.FirstName)
	assert.Equal(t, input.LastName, user.LastName)
	assert.Equal(t, input.Phone, user.Phone)
	assert.Equal(t, input.MemberLevel, user.MemberLevel)
	assert.Equal(t, input.PointBalance, user.PointBalance)
	assert.Equal(t, input.Email, user.Email)
	mockRepo.AssertExpectations(t)
}

func TestUpdateUser_NotFound(t *testing.T) {
	mockRepo := new(MockUserRepository)
	useCase := NewUserUseCase(mockRepo)

	input := UpdateUserInput{
		FirstName: "Jane",
	}

	mockRepo.On("FindByID", 999).Return(nil, domain.ErrUserNotFound)

	user, err := useCase.UpdateUser(999, input)

	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, domain.ErrUserNotFound, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteUser_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	useCase := NewUserUseCase(mockRepo)

	now := time.Now()
	existingUser := &domain.User{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		CreatedAt: now,
		UpdatedAt: now,
	}

	mockRepo.On("FindByID", 1).Return(existingUser, nil)
	mockRepo.On("Delete", 1).Return(nil)

	err := useCase.DeleteUser(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteUser_NotFound(t *testing.T) {
	mockRepo := new(MockUserRepository)
	useCase := NewUserUseCase(mockRepo)

	mockRepo.On("FindByID", 999).Return(nil, domain.ErrUserNotFound)

	err := useCase.DeleteUser(999)

	assert.Error(t, err)

	assert.Error(t, err)
	assert.Equal(t, domain.ErrUserNotFound, err)
	mockRepo.AssertExpectations(t)
}
