package main

import (
	"log"
	"workshop_4/config"
	"workshop_4/database"
	"workshop_4/internal/infrastructure/repository"
	httphandler "workshop_4/internal/interfaces/http"
	"workshop_4/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	if err := database.InitDB(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer database.CloseDB()

	// Initialize Clean Architecture layers
	// Infrastructure Layer - Repository
	userRepo := repository.NewSQLiteUserRepository(database.DB)

	// Use Case Layer - Business Logic
	userUseCase := usecase.NewUserUseCase(userRepo)

	// Interface Layer - HTTP Handlers
	userHandler := httphandler.NewUserHandler(userUseCase)

	// Create a new Fiber instance
	app := fiber.New(fiber.Config{
		AppName: cfg.AppName,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${method} ${path}\n",
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Setup routes
	setupRoutes(app, userHandler)

	// Start server
	log.Printf("🚀 Server starting on port %s (Environment: %s)", cfg.Port, cfg.Environment)
	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func setupRoutes(app *fiber.App, userHandler *httphandler.UserHandler) {
	// Root endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to Workshop 4 API",
			"status":  "running",
		})
	})

	// API v1 routes
	api := app.Group("/api/v1")

	// User routes
	users := api.Group("/users")
	users.Get("/", userHandler.GetUsers)
	users.Get("/:id", userHandler.GetUser)
	users.Post("/", userHandler.CreateUser)
	users.Put("/:id", userHandler.UpdateUser)
	users.Delete("/:id", userHandler.DeleteUser)
}
