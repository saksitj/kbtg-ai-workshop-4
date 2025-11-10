package models

import "time"

// User represents a user in the system
type User struct {
	ID           int       `json:"id"`
	FirstName    string    `json:"first_name" validate:"required"`
	LastName     string    `json:"last_name" validate:"required"`
	Email        string    `json:"email" validate:"required,email"`
	Phone        string    `json:"phone"`
	Address      string    `json:"address"`
	Avatar       string    `json:"avatar"`
	MemberLevel  string    `json:"member_level"`  // e.g., "Gold", "Silver", "Bronze", "Platinum"
	PointBalance int       `json:"point_balance"` // แต้มคงเหลือ
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// CreateUserRequest represents the request body for creating a user
type CreateUserRequest struct {
	FirstName    string `json:"first_name" validate:"required"`
	LastName     string `json:"last_name" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	Phone        string `json:"phone"`
	Address      string `json:"address"`
	Avatar       string `json:"avatar"`
	MemberLevel  string `json:"member_level"`
	PointBalance int    `json:"point_balance"`
}

// UpdateUserRequest represents the request body for updating a user
type UpdateUserRequest struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email" validate:"email"`
	Phone        string `json:"phone"`
	Address      string `json:"address"`
	Avatar       string `json:"avatar"`
	MemberLevel  string `json:"member_level"`
	PointBalance int    `json:"point_balance"`
}

// UserResponse represents the response structure for user data
type UserResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
}
