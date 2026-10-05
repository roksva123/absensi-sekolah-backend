.PHONY: help build run dev test clean migrate docker-build docker-up docker-down docker-logs

help:
	@echo "Absensi Sekolah Backend - Available Commands"
	@echo "============================================"
	@echo "Development:"
	@echo "  make run          - Run the application"
	@echo "  make dev          - Run with hot reload (requires air)"
	@echo "  make build        - Build the binary"
	@echo ""
	@echo "Testing:"
	@echo "  make test         - Run tests"
	@echo ""
	@echo "Database:"
	@echo "  make migrate      - Run auto-migrations"
	@echo ""
	@echo "Docker:"
	@echo "  make docker-build - Build Docker image"
	@echo "  make docker-up    - Start Docker containers"
	@echo "  make docker-down  - Stop Docker containers"
	@echo "  make docker-logs  - View container logs"
	@echo ""
	@echo "Cleanup:"
	@echo "  make clean        - Clean build artifacts"
	@echo "  make deps         - Download dependencies"

build:
	@echo "Building binary..."
	go build -o absensi-backend .
	@echo "✅ Binary built: absensi-backend"

run: build
	@echo "Starting server..."
	./absensi-backend

dev:
	@echo "Starting server with hot reload..."
	@command -v air >/dev/null 2>&1 || go install github.com/cosmtrek/air@latest
	air

test:
	@echo "Running tests..."
	go test -v ./...

clean:
	@echo "Cleaning up..."
	rm -f absensi-backend
	go clean
	@echo "✅ Clean complete"

deps:
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy
	@echo "✅ Dependencies downloaded"

migrate:
	@echo "Running migrations..."
	go run main.go
	@echo "✅ Migrations complete"

docker-build:
	@echo "Building Docker image..."
	docker-compose build
	@echo "✅ Docker image built"

docker-up:
	@echo "Starting Docker containers..."
	docker-compose up -d
	@echo "✅ Containers started"
	@echo "Backend: http://localhost:8080"

docker-down:
	@echo "Stopping Docker containers..."
	docker-compose down
	@echo "✅ Containers stopped"

docker-logs:
	docker-compose logs -f backend

docker-ps:
	docker-compose ps

fmt:
	@echo "Formatting code..."
	go fmt ./...
	@echo "✅ Code formatted"

lint:
	@echo "Running linter..."
	@command -v golangci-lint >/dev/null 2>&1 || go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run ./...
	@echo "✅ Linting complete"

vet:
	@echo "Running go vet..."
	go vet ./...
	@echo "✅ Vet check complete"
