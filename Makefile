.PHONY: run build clean test help dev

# Variables
APP_NAME=workshop4
BUILD_DIR=bin
MAIN_FILE=main.go

# Help command
help:
	@echo "Available commands:"
	@echo "  make run      - Run the application"
	@echo "  make dev      - Run with hot reload (requires air)"
	@echo "  make build    - Build the application"
	@echo "  make clean    - Clean build artifacts"
	@echo "  make test     - Run tests"
	@echo "  make install  - Install dependencies"
	@echo "  make tidy     - Tidy and verify dependencies"

# Run the application
run:
	@echo "🚀 Starting application..."
	@go run $(MAIN_FILE)

# Development mode with hot reload
dev:
	@echo "🔥 Starting development mode..."
	@air

# Build the application
build:
	@echo "🔨 Building application..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)
	@echo "✅ Build complete: $(BUILD_DIR)/$(APP_NAME)"

# Clean build artifacts
clean:
	@echo "🧹 Cleaning..."
	@rm -rf $(BUILD_DIR)
	@go clean
	@echo "✅ Clean complete"

# Run tests
test:
	@echo "🧪 Running tests..."
	@go test -v ./...

# Run tests with coverage
test-coverage:
	@echo "🧪 Running tests with coverage..."
	@go test -cover ./...

# Run tests with detailed coverage
test-coverage-html:
	@echo "🧪 Generating coverage report..."
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "✅ Coverage report generated: coverage.html"

# Install dependencies
install:
	@echo "📦 Installing dependencies..."
	@go mod download
	@echo "✅ Dependencies installed"

# Tidy dependencies
tidy:
	@echo "🔧 Tidying dependencies..."
	@go mod tidy
	@echo "✅ Dependencies tidied"

# Default target
.DEFAULT_GOAL := help
