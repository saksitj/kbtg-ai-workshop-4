# Workshop 4 - Go Fiber Backend Project Structure

```
workshop_4/
├── config/
│   └── config.go              # Configuration management
│
├── handlers/
│   └── user.go                # User handler functions (controllers)
│
├── middleware/
│   └── auth.go                # Authentication and custom middleware
│
├── models/
│   └── user.go                # Data models and structs
│
├── routes/
│   └── routes.go              # Route definitions and setup
│
├── .env.example               # Example environment variables
├── .gitignore                 # Git ignore rules
├── API_TESTING.md             # API testing examples with curl
├── go.mod                     # Go module definition
├── go.sum                     # Go dependencies checksum
├── main.go                    # Application entry point
├── Makefile                   # Build and run commands
├── PROJECT_STRUCTURE.md       # This file
└── README.md                  # Project documentation
```

## File Descriptions

### Core Files

- **main.go**: Entry point of the application. Sets up Fiber, middleware, and starts the server.
- **go.mod**: Defines the Go module and dependencies.
- **go.sum**: Contains checksums of dependencies for verification.

### Configuration

- **config/config.go**: Manages application configuration from environment variables.
- **.env.example**: Template for environment variables.

### Routing

- **routes/routes.go**: Central place for all route definitions and API endpoint setup.

### Handlers

- **handlers/user.go**: Contains handler functions (controllers) for user-related operations:
  - GetUsers: GET /api/v1/users
  - GetUser: GET /api/v1/users/:id
  - CreateUser: POST /api/v1/users
  - UpdateUser: PUT /api/v1/users/:id
  - DeleteUser: DELETE /api/v1/users/:id

### Models

- **models/user.go**: Data structures and types:
  - User struct
  - Request/Response structures
  - Validation tags

### Middleware

- **middleware/auth.go**: Custom middleware functions:
  - Authentication
  - Authorization
  - Request logging

### Documentation

- **README.md**: Complete project documentation with setup and usage instructions.
- **API_TESTING.md**: cURL examples for testing all API endpoints.
- **PROJECT_STRUCTURE.md**: This file, explaining the project structure.

### Build Tools

- **Makefile**: Convenient commands for common tasks:
  - `make run`: Run the application
  - `make build`: Build executable
  - `make test`: Run tests
  - `make clean`: Clean build artifacts
  - `make tidy`: Tidy dependencies

## Quick Start

1. Install dependencies:
   ```bash
   go mod download
   ```

2. Copy environment file:
   ```bash
   cp .env.example .env
   ```

3. Run the application:
   ```bash
   make run
   # or
   go run main.go
   ```

4. Server will start at http://localhost:3000

## Available Endpoints

- `GET /health` - Health check
- `GET /` - Welcome message
- `GET /api/v1/users` - Get all users
- `GET /api/v1/users/:id` - Get user by ID
- `POST /api/v1/users` - Create new user
- `PUT /api/v1/users/:id` - Update user
- `DELETE /api/v1/users/:id` - Delete user

## Next Steps

1. **Database Integration**: Add database support (PostgreSQL, MySQL, MongoDB)
2. **Authentication**: Implement JWT authentication
3. **Validation**: Add request validation middleware
4. **Testing**: Write unit and integration tests
5. **Logging**: Add structured logging
6. **Docker**: Add Dockerfile and docker-compose
7. **API Documentation**: Add Swagger/OpenAPI documentation
