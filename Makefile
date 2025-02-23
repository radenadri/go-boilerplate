.PHONY: build run test clean dev migrate migrate-create migrate-up migrate-down lint

BUILD_DIR=./tmp
BINARY_NAME=main
MAIN_FILE=./cmd/api/main.go

build:
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_FILE)

run: build
	$(BUILD_DIR)/$(BINARY_NAME)

dev:
	air -c .air.toml

test:
	go test -v ./...

clean:
	rm -rf $(BUILD_DIR)

# Database commands
db-create:
	psql -U $(shell grep DB_USERNAME .env | cut -d '=' -f2) -h $(shell grep DB_HOST .env | cut -d '=' -f2) -c 'CREATE DATABASE $(shell grep DB_DATABASE .env | cut -d '=' -f2);'

db-drop:
	psql -U $(shell grep DB_USERNAME .env | cut -d '=' -f2) -h $(shell grep DB_HOST .env | cut -d '=' -f2) -c 'DROP DATABASE IF EXISTS $(shell grep DB_DATABASE .env | cut -d '=' -f2);'

# Migration commands
migrate-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir migrations -seq $$name

migrate-up:
	migrate -path migrations -database "postgres://$(shell grep DB_USERNAME .env | cut -d '=' -f2):$(shell grep DB_PASSWORD .env | cut -d '=' -f2)@$(shell grep DB_HOST .env | cut -d '=' -f2):$(shell grep DB_PORT .env | cut -d '=' -f2)/$(shell grep DB_DATABASE .env | cut -d '=' -f2)?sslmode=$(shell if [ "$(shell grep DB_SSL_ENABLED .env | cut -d '=' -f2)" = "true" ]; then echo "enable"; else echo "disable"; fi)" up

migrate-down:

# Linting
lint:
	golangci-lint run
	migrate -path migrations -database "postgres://$(shell grep DB_USERNAME .env | cut -d '=' -f2):$(shell grep DB_PASSWORD .env | cut -d '=' -f2)@$(shell grep DB_HOST .env | cut -d '=' -f2):$(shell grep DB_PORT .env | cut -d '=' -f2)/$(shell grep DB_DATABASE .env | cut -d '=' -f2)?sslmode=$(shell if [ "$(shell grep DB_SSL_ENABLED .env | cut -d '=' -f2)" = "true" ]; then echo "enable"; else echo "disable"; fi)" down

# Swagger documentation
swag:
	swag init -g cmd/api/main.go

# Help command
help:
	@echo "Available commands:"
	@echo "  build          - Build the application"
	@echo "  run            - Build and run the application"
	@echo "  dev            - Run the application with hot reload using Air"
	@echo "  test           - Run tests"
	@echo "  clean          - Remove build artifacts"
	@echo "  db-create      - Create database"
	@echo "  db-drop        - Drop database"
	@echo "  migrate-create - Create a new migration file"
	@echo "  migrate-up     - Run all pending migrations"
	@echo "  migrate-down   - Rollback the last migration"
	@echo "  swag           - Generate Swagger documentation"
	@echo "  help           - Show this help message"
