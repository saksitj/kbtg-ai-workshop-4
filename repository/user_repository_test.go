package repository

import (
	"database/sql"
	"testing"
	"time"
	"workshop_4/models"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func setupTestDB(t *testing.T) (*sql.DB, func()) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Failed to open test database: %v", err)
	}

	// Create users table
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		phone TEXT,
		address TEXT,
		avatar TEXT,
		member_level TEXT DEFAULT 'Bronze',
		point_balance INTEGER DEFAULT 0,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		t.Fatalf("Failed to create test table: %v", err)
	}

	cleanup := func() {
		db.Close()
	}

	return db, cleanup
}

func TestUserRepository_Create(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewUserRepository(db)

	req := models.CreateUserRequest{
		FirstName:    "John",
		LastName:     "Doe",
		Email:        "john@example.com",
		Phone:        "0812345678",
		Address:      "123 Main St",
		Avatar:       "https://example.com/avatar.jpg",
		MemberLevel:  "Gold",
		PointBalance: 100,
	}

	user, err := repo.Create(req)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, 1, user.ID)
	assert.Equal(t, "John", user.FirstName)
	assert.Equal(t, "Doe", user.LastName)
	assert.Equal(t, "john@example.com", user.Email)
	assert.Equal(t, "Gold", user.MemberLevel)
	assert.Equal(t, 100, user.PointBalance)
}

func TestUserRepository_GetByID(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewUserRepository(db)

	// Create a user first
	req := models.CreateUserRequest{
		FirstName:    "Jane",
		LastName:     "Smith",
		Email:        "jane@example.com",
		MemberLevel:  "Silver",
		PointBalance: 50,
	}

	createdUser, _ := repo.Create(req)

	// Get the user by ID
	user, err := repo.GetByID(createdUser.ID)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, createdUser.ID, user.ID)
	assert.Equal(t, "Jane", user.FirstName)
	assert.Equal(t, "Smith", user.LastName)
	assert.Equal(t, "jane@example.com", user.Email)
	assert.Equal(t, "Silver", user.MemberLevel)
	assert.Equal(t, 50, user.PointBalance)
}

func TestUserRepository_GetByID_NotFound(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewUserRepository(db)

	user, err := repo.GetByID(999)

	assert.NoError(t, err)
	assert.Nil(t, user)
}

func TestUserRepository_GetAll(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewUserRepository(db)

	// Create multiple users
	users := []models.CreateUserRequest{
		{
			FirstName:    "User1",
			LastName:     "Test",
			Email:        "user1@example.com",
			MemberLevel:  "Bronze",
			PointBalance: 10,
		},
		{
			FirstName:    "User2",
			LastName:     "Test",
			Email:        "user2@example.com",
			MemberLevel:  "Gold",
			PointBalance: 200,
		},
		{
			FirstName:    "User3",
			LastName:     "Test",
			Email:        "user3@example.com",
			MemberLevel:  "Platinum",
			PointBalance: 500,
		},
	}

	for _, req := range users {
		repo.Create(req)
	}

	// Get all users
	allUsers, err := repo.GetAll()

	assert.NoError(t, err)
	assert.Len(t, allUsers, 3)
	assert.Equal(t, "User3", allUsers[0].FirstName) // Should be in DESC order
}

func TestUserRepository_Update(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewUserRepository(db)

	// Create a user first
	createReq := models.CreateUserRequest{
		FirstName:    "John",
		LastName:     "Doe",
		Email:        "john@example.com",
		MemberLevel:  "Bronze",
		PointBalance: 0,
	}

	createdUser, _ := repo.Create(createReq)

	// Update the user
	time.Sleep(time.Millisecond * 10) // Ensure updated_at is different

	updateReq := models.UpdateUserRequest{
		FirstName:    "John",
		LastName:     "Updated",
		Email:        "john.updated@example.com",
		Phone:        "0899999999",
		MemberLevel:  "Gold",
		PointBalance: 500,
	}

	updatedUser, err := repo.Update(createdUser.ID, updateReq)

	assert.NoError(t, err)
	assert.NotNil(t, updatedUser)
	assert.Equal(t, createdUser.ID, updatedUser.ID)
	assert.Equal(t, "Updated", updatedUser.LastName)
	assert.Equal(t, "john.updated@example.com", updatedUser.Email)
	assert.Equal(t, "Gold", updatedUser.MemberLevel)
	assert.Equal(t, 500, updatedUser.PointBalance)
	assert.NotEqual(t, createdUser.UpdatedAt, updatedUser.UpdatedAt)
}

func TestUserRepository_Delete(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewUserRepository(db)

	// Create a user first
	req := models.CreateUserRequest{
		FirstName: "Delete",
		LastName:  "Me",
		Email:     "delete@example.com",
	}

	createdUser, _ := repo.Create(req)

	// Delete the user
	err := repo.Delete(createdUser.ID)

	assert.NoError(t, err)

	// Verify user is deleted
	user, err := repo.GetByID(createdUser.ID)
	assert.NoError(t, err)
	assert.Nil(t, user)
}

func TestUserRepository_Create_DuplicateEmail(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewUserRepository(db)

	req := models.CreateUserRequest{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "duplicate@example.com",
	}

	// Create first user
	_, err := repo.Create(req)
	assert.NoError(t, err)

	// Try to create second user with same email
	_, err = repo.Create(req)
	assert.Error(t, err) // Should fail due to unique constraint
}

func TestUserRepository_GetAll_Empty(t *testing.T) {
	db, cleanup := setupTestDB(t)
	defer cleanup()

	repo := NewUserRepository(db)

	users, err := repo.GetAll()

	assert.NoError(t, err)
	assert.Empty(t, users)
}
