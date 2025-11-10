package repository

import "workshop_4/models"

// UserRepositoryInterface defines the methods for user repository
type UserRepositoryInterface interface {
	GetAll() ([]models.User, error)
	GetByID(id int) (*models.User, error)
	Create(req models.CreateUserRequest) (*models.User, error)
	Update(id int, req models.UpdateUserRequest) (*models.User, error)
	Delete(id int) error
}
