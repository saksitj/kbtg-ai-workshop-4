package domain

import "errors"

// Domain errors
var (
	ErrUserNotFound        = errors.New("user not found")
	ErrFirstNameRequired   = errors.New("first name is required")
	ErrLastNameRequired    = errors.New("last name is required")
	ErrEmailRequired       = errors.New("email is required")
	ErrInvalidEmail        = errors.New("invalid email format")
	ErrDuplicateEmail      = errors.New("email already exists")
	ErrInvalidUserID       = errors.New("invalid user ID")
	ErrInvalidMemberLevel  = errors.New("invalid member level")
	ErrInvalidPointBalance = errors.New("point balance cannot be negative")
)
