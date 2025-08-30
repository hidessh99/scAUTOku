# Makefile for VPN Account Management API

.PHONY: build run dev test clean deps

# Build the application
build:
	go build -o bin/api main.go

# Run the application
run:
	go run main.go

# Run in development mode with hot reload (requires air)
dev:
	air

# Run tests
test:
	go test -v ./...

# Clean build artifacts
clean:
	rm -rf bin/

# Install dependencies
deps:
	go mod tidy

# Install air for hot reloading (development)
install-air:
	go install github.com/cosmtrek/air@latest

# Build Docker image
docker-build:
	docker build -t vpn-account-api .

# Run Docker container
docker-run:
	docker run -p 3000:3000 vpn-account-api

# Build and run with docker-compose
docker-up:
	docker-compose up --build

# Stop docker-compose
docker-down:
	docker-compose down