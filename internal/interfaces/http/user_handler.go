package http

import (
	"strconv"
	"workshop_4/internal/domain"
	"workshop_4/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

// UserHandler handles HTTP requests for users
type UserHandler struct {
	userUseCase *usecase.UserUseCase
}

// NewUserHandler creates a new user handler
func NewUserHandler(userUseCase *usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

// UserResponse represents the API response for user data
type UserResponse struct {
	ID           int    `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Address      string `json:"address"`
	Avatar       string `json:"avatar"`
	MemberLevel  string `json:"member_level"`
	PointBalance int    `json:"point_balance"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

// toUserResponse converts domain user to response
func toUserResponse(user *domain.User) UserResponse {
	return UserResponse{
		ID:           user.ID,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Email:        user.Email,
		Phone:        user.Phone,
		Address:      user.Address,
		Avatar:       user.Avatar,
		MemberLevel:  user.MemberLevel,
		PointBalance: user.PointBalance,
		CreatedAt:    user.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:    user.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}

// GetUsers handles GET /users
func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	users, err := h.userUseCase.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Failed to fetch users",
		})
	}

	responses := make([]UserResponse, len(users))
	for i, user := range users {
		responses[i] = toUserResponse(user)
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    responses,
	})
}

// GetUser handles GET /users/:id
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid user ID",
		})
	}

	user, err := h.userUseCase.GetUserByID(id)
	if err == domain.ErrUserNotFound {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   "User not found",
		})
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Failed to fetch user",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    toUserResponse(user),
	})
}

// CreateUserRequest represents the request body for creating a user
type CreateUserRequest struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Address      string `json:"address"`
	Avatar       string `json:"avatar"`
	MemberLevel  string `json:"member_level"`
	PointBalance int    `json:"point_balance"`
}

// CreateUser handles POST /users
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid request body",
		})
	}

	input := usecase.CreateUserInput{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Email:        req.Email,
		Phone:        req.Phone,
		Address:      req.Address,
		Avatar:       req.Avatar,
		MemberLevel:  req.MemberLevel,
		PointBalance: req.PointBalance,
	}

	user, err := h.userUseCase.CreateUser(input)
	if err == domain.ErrFirstNameRequired || err == domain.ErrLastNameRequired || err == domain.ErrEmailRequired {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	if err == domain.ErrDuplicateEmail {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Failed to create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    toUserResponse(user),
	})
}

// UpdateUserRequest represents the request body for updating a user
type UpdateUserRequest struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Address      string `json:"address"`
	Avatar       string `json:"avatar"`
	MemberLevel  string `json:"member_level"`
	PointBalance int    `json:"point_balance"`
}

// UpdateUser handles PUT /users/:id
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid user ID",
		})
	}

	var req UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid request body",
		})
	}

	input := usecase.UpdateUserInput{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Email:        req.Email,
		Phone:        req.Phone,
		Address:      req.Address,
		Avatar:       req.Avatar,
		MemberLevel:  req.MemberLevel,
		PointBalance: req.PointBalance,
	}

	user, err := h.userUseCase.UpdateUser(id, input)
	if err == domain.ErrUserNotFound {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   "User not found",
		})
	}
	if err == domain.ErrFirstNameRequired || err == domain.ErrLastNameRequired || err == domain.ErrEmailRequired {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Failed to update user",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    toUserResponse(user),
	})
}

// DeleteUser handles DELETE /users/:id
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid user ID",
		})
	}

	err = h.userUseCase.DeleteUser(id)
	if err == domain.ErrUserNotFound {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   "User not found",
		})
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Failed to delete user",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "User deleted successfully",
	})
}
