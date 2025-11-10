package domain

// UserRepository defines the interface for user data operations
type UserRepository interface {
	FindAll() ([]*User, error)
	FindByID(id int) (*User, error)
	FindByEmail(email string) (*User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id int) error
}
