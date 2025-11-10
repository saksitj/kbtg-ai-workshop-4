# Workshop 4 - Go + Fiber Backend API

## Description
Backend API built with Go and Fiber framework for Workshop 4.

## Features
- 🚀 Fast HTTP server using Fiber v2
- 📁 Clean project structure
- 🛣️ RESTful API routes
- 🔒 Middleware support (CORS, Logger, Recovery)
- ⚙️ Environment configuration
- 📝 Example CRUD operations

## Prerequisites
- Go 1.21 or higher
- Git

## Installation

1. Clone the repository:
```bash
git clone <your-repo-url>
cd workshop_4
```

2. Install dependencies:
```bash
go mod download
```

3. Create environment file:
```bash
cp .env.example .env
```

4. Edit `.env` file with your configuration

## Running the Application

### Development
```bash
go run main.go
```

### Build and Run
```bash
go build -o bin/app
./bin/app
```

## Project Structure
```
workshop_4/
├── config/          # Configuration files
│   └── config.go
├── handlers/        # Request handlers
│   └── user.go
├── middleware/      # Custom middleware
│   └── auth.go
├── routes/          # Route definitions
│   └── routes.go
├── .env.example     # Example environment variables
├── .gitignore       # Git ignore file
├── go.mod           # Go module file
├── go.sum           # Go dependencies checksum
├── main.go          # Application entry point
└── README.md        # This file
```

## API Endpoints

### Health Check
```
GET /health
```

### Welcome
```
GET /
```

### Users API (v1)
```
GET    /api/v1/users     - Get all users
GET    /api/v1/users/:id - Get user by ID
POST   /api/v1/users     - Create new user
PUT    /api/v1/users/:id - Update user
DELETE /api/v1/users/:id - Delete user
```

## Example API Requests

### Get all users
```bash
curl http://localhost:3000/api/v1/users
```

### Create user
```bash
curl -X POST http://localhost:3000/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","email":"john@example.com"}'
```

### Get user by ID
```bash
curl http://localhost:3000/api/v1/users/1
```

### Update user
```bash
curl -X PUT http://localhost:3000/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"John Updated","email":"john.updated@example.com"}'
```

### Delete user
```bash
curl -X DELETE http://localhost:3000/api/v1/users/1
```

## Environment Variables

| Variable    | Description                | Default          |
|-------------|----------------------------|------------------|
| PORT        | Server port                | 3000             |
| ENVIRONMENT | Environment (dev/prod)     | development      |
| APP_NAME    | Application name           | Workshop 4 API   |

## Technologies Used
- [Go](https://go.dev/) - Programming language
- [Fiber](https://gofiber.io/) - Web framework
- [Fiber Middleware](https://docs.gofiber.io/api/middleware) - CORS, Logger, Recovery

## Documentation
- [Fiber Documentation](https://docs.gofiber.io/)
- [Go Documentation](https://go.dev/doc/)

## License
MIT
