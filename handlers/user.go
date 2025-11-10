package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// GetUsers returns a list of all users
func GetUsers(c *fiber.Ctx) error {
	// TODO: Implement database query
	users := []fiber.Map{
		{"id": 1, "name": "John Doe", "email": "john@example.com"},
		{"id": 2, "name": "Jane Smith", "email": "jane@example.com"},
	}
	return c.JSON(fiber.Map{
		"success": true,
		"data":    users,
	})
}

// GetUser returns a single user by ID
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	// TODO: Implement database query
	return c.JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"id":    id,
			"name":  "John Doe",
			"email": "john@example.com",
		},
	})
}

// CreateUser creates a new user
func CreateUser(c *fiber.Ctx) error {
	type CreateUserRequest struct {
		Name  string `json:"name" validate:"required"`
		Email string `json:"email" validate:"required,email"`
	}

	var req CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid request body",
		})
	}

	// TODO: Implement database insert
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"id":    3,
			"name":  req.Name,
			"email": req.Email,
		},
	})
}

// UpdateUser updates an existing user
func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")

	type UpdateUserRequest struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	var req UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid request body",
		})
	}

	// TODO: Implement database update
	return c.JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"id":    id,
			"name":  req.Name,
			"email": req.Email,
		},
	})
}

// DeleteUser deletes a user by ID
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	// TODO: Implement database delete
	return c.JSON(fiber.Map{
		"success": true,
		"message": "User deleted successfully",
		"id":      id,
	})
}
