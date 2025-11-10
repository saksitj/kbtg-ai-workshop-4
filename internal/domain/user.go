package domain

import "time"

// User represents the core business entity
type User struct {
	ID           int
	FirstName    string
	LastName     string
	Email        string
	Phone        string
	Address      string
	Avatar       string
	MemberLevel  string
	PointBalance int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// Validate validates the user entity
func (u *User) Validate() error {
	if u.FirstName == "" {
		return ErrFirstNameRequired
	}
	if u.LastName == "" {
		return ErrLastNameRequired
	}
	if u.Email == "" {
		return ErrEmailRequired
	}
	return nil
}

// GetFullName returns the user's full name
func (u *User) GetFullName() string {
	return u.FirstName + " " + u.LastName
}

// IsActive checks if user has member level
func (u *User) IsActive() bool {
	return u.MemberLevel != ""
}
