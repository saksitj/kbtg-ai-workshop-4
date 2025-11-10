package usecase

import (
	"time"
	"workshop_4/internal/domain"
)

// UserUseCase handles user business logic
type UserUseCase struct {
	userRepo domain.UserRepository
}

// NewUserUseCase creates a new user use case
func NewUserUseCase(userRepo domain.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
	}
}

// GetAllUsers retrieves all users
func (uc *UserUseCase) GetAllUsers() ([]*domain.User, error) {
	return uc.userRepo.FindAll()
}

// GetUserByID retrieves a user by ID
func (uc *UserUseCase) GetUserByID(id int) (*domain.User, error) {
	if id <= 0 {
		return nil, domain.ErrInvalidUserID
	}

	user, err := uc.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, domain.ErrUserNotFound
	}

	return user, nil
}

// CreateUserInput represents input for creating a user
type CreateUserInput struct {
	FirstName    string
	LastName     string
	Email        string
	Phone        string
	Address      string
	Avatar       string
	MemberLevel  string
	PointBalance int
}

// CreateUser creates a new user
func (uc *UserUseCase) CreateUser(input CreateUserInput) (*domain.User, error) {
	// Set defaults
	if input.MemberLevel == "" {
		input.MemberLevel = "Bronze"
	}

	// Create user entity
	user := &domain.User{
		FirstName:    input.FirstName,
		LastName:     input.LastName,
		Email:        input.Email,
		Phone:        input.Phone,
		Address:      input.Address,
		Avatar:       input.Avatar,
		MemberLevel:  input.MemberLevel,
		PointBalance: input.PointBalance,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// Validate
	if err := user.Validate(); err != nil {
		return nil, err
	}

	// Check if email already exists
	existing, err := uc.userRepo.FindByEmail(input.Email)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, domain.ErrDuplicateEmail
	}

	// Save to repository
	if err := uc.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateUserInput represents input for updating a user
type UpdateUserInput struct {
	FirstName    string
	LastName     string
	Email        string
	Phone        string
	Address      string
	Avatar       string
	MemberLevel  string
	PointBalance int
}

// UpdateUser updates an existing user
func (uc *UserUseCase) UpdateUser(id int, input UpdateUserInput) (*domain.User, error) {
	if id <= 0 {
		return nil, domain.ErrInvalidUserID
	}

	// Get existing user
	user, err := uc.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, domain.ErrUserNotFound
	}

	// Update fields
	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.Email = input.Email
	user.Phone = input.Phone
	user.Address = input.Address
	user.Avatar = input.Avatar
	user.MemberLevel = input.MemberLevel
	user.PointBalance = input.PointBalance
	user.UpdatedAt = time.Now()

	// Validate
	if err := user.Validate(); err != nil {
		return nil, err
	}

	// Save to repository
	if err := uc.userRepo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser deletes a user by ID
func (uc *UserUseCase) DeleteUser(id int) error {
	if id <= 0 {
		return domain.ErrInvalidUserID
	}

	// Check if user exists
	user, err := uc.userRepo.FindByID(id)
	if err != nil {
		return err
	}
	if user == nil {
		return domain.ErrUserNotFound
	}

	return uc.userRepo.Delete(id)
}
