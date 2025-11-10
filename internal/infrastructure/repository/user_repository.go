package repository

import (
	"database/sql"
	"workshop_4/internal/domain"
)

// sqliteUserRepository implements domain.UserRepository
type sqliteUserRepository struct {
	db *sql.DB
}

// NewSQLiteUserRepository creates a new SQLite user repository
func NewSQLiteUserRepository(db *sql.DB) domain.UserRepository {
	return &sqliteUserRepository{db: db}
}

// FindAll retrieves all users from the database
func (r *sqliteUserRepository) FindAll() ([]*domain.User, error) {
	query := `SELECT id, first_name, last_name, email, phone, address, avatar, member_level, point_balance, created_at, updated_at 
	          FROM users ORDER BY id DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*domain.User
	for rows.Next() {
		user := &domain.User{}
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

// FindByID retrieves a user by ID
func (r *sqliteUserRepository) FindByID(id int) (*domain.User, error) {
	query := `SELECT id, first_name, last_name, email, phone, address, avatar, member_level, point_balance, created_at, updated_at 
	          FROM users WHERE id = ?`

	user := &domain.User{}
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

	return user, nil
}

// FindByEmail retrieves a user by email
func (r *sqliteUserRepository) FindByEmail(email string) (*domain.User, error) {
	query := `SELECT id, first_name, last_name, email, phone, address, avatar, member_level, point_balance, created_at, updated_at 
	          FROM users WHERE email = ?`

	user := &domain.User{}
	err := r.db.QueryRow(query, email).Scan(
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

	return user, nil
}

// Create inserts a new user into the database
func (r *sqliteUserRepository) Create(user *domain.User) error {
	query := `INSERT INTO users (first_name, last_name, email, phone, address, avatar, member_level, point_balance, created_at, updated_at) 
			  VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := r.db.Exec(query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Phone,
		user.Address,
		user.Avatar,
		user.MemberLevel,
		user.PointBalance,
		user.CreatedAt,
		user.UpdatedAt,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = int(id)
	return nil
}

// Update modifies an existing user in the database
func (r *sqliteUserRepository) Update(user *domain.User) error {
	query := `UPDATE users 
	          SET first_name = ?, last_name = ?, email = ?, phone = ?, address = ?, avatar = ?, member_level = ?, point_balance = ?, updated_at = ? 
			  WHERE id = ?`

	_, err := r.db.Exec(query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Phone,
		user.Address,
		user.Avatar,
		user.MemberLevel,
		user.PointBalance,
		user.UpdatedAt,
		user.ID,
	)

	return err
}

// Delete removes a user from the database
func (r *sqliteUserRepository) Delete(id int) error {
	query := `DELETE FROM users WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}
