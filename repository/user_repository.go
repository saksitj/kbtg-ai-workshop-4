package repository

import (
	"database/sql"
	"time"
	"workshop_4/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// GetAll retrieves all users from the database
func (r *UserRepository) GetAll() ([]models.User, error) {
	query := `SELECT id, first_name, last_name, email, phone, address, avatar, member_level, point_balance, created_at, updated_at FROM users ORDER BY id DESC`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Phone,
			&user.Address,
			&user.Avatar,
			&user.MemberLevel,
			&user.PointBalance,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// GetByID retrieves a user by ID
func (r *UserRepository) GetByID(id int) (*models.User, error) {
	query := `SELECT id, first_name, last_name, email, phone, address, avatar, member_level, point_balance, created_at, updated_at FROM users WHERE id = ?`

	var user models.User
	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Phone,
		&user.Address,
		&user.Avatar,
		&user.MemberLevel,
		&user.PointBalance,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Create inserts a new user into the database
func (r *UserRepository) Create(req models.CreateUserRequest) (*models.User, error) {
	query := `INSERT INTO users (first_name, last_name, email, phone, address, avatar, member_level, point_balance, created_at, updated_at) 
			  VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	now := time.Now()
	result, err := r.db.Exec(query, req.FirstName, req.LastName, req.Email, req.Phone, req.Address, req.Avatar, req.MemberLevel, req.PointBalance, now, now)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return r.GetByID(int(id))
}

// Update modifies an existing user in the database
func (r *UserRepository) Update(id int, req models.UpdateUserRequest) (*models.User, error) {
	query := `UPDATE users SET first_name = ?, last_name = ?, email = ?, phone = ?, address = ?, avatar = ?, member_level = ?, point_balance = ?, updated_at = ? 
			  WHERE id = ?`

	now := time.Now()
	_, err := r.db.Exec(query, req.FirstName, req.LastName, req.Email, req.Phone, req.Address, req.Avatar, req.MemberLevel, req.PointBalance, now, id)
	if err != nil {
		return nil, err
	}

	return r.GetByID(id)
}

// Delete removes a user from the database
func (r *UserRepository) Delete(id int) error {
	query := `DELETE FROM users WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}
