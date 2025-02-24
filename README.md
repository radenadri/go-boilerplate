# Go Boilerplate

A robust and scalable Go boilerplate for building modern web applications with best practices and essential features pre-configured. This boilerplate follows clean architecture principles and includes production-ready features out of the box.

## Table of Contents

- [Features](#features)
- [Tech Stack](#tech-stack)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [Project Structure](#project-structure)
- [Configuration](#configuration)
- [Makefile Commands](#makefile-commands)
- [API Documentation](#api-documentation)
- [Development](#development)
- [Contributing](#contributing)
- [License](#license)

## Features

### Core Features
- **Clean Architecture**: Implements Uncle Bob's Clean Architecture principles for better code organization and maintainability
- **RESTful API**: High-performance HTTP routing with [Gin Framework](https://github.com/gin-gonic/gin)
- **Database Integration**: PostgreSQL with [GORM v2](https://gorm.io) for reliable data persistence
- **Caching**: Integrated [Redis](https://github.com/redis/go-redis) for caching and rate limiting
- **Authentication**: Secure JWT authentication using [golang-jwt](https://github.com/golang-jwt/jwt)

### Developer Experience
- **Hot Reload**: Live code reloading with [Air](https://github.com/cosmtrek/air)
- **Environment Management**: Configuration using [GoDotEnv](https://github.com/joho/godotenv)
- **Containerization**: Docker and docker-compose support for consistent development and deployment
- **API Documentation**: Automated API documentation with [Swag](https://github.com/swaggo/swag)

### Reliability & Monitoring
- **Structured Logging**: High-performance logging with [Zap](https://github.com/uber-go/zap)
- **Input Validation**: Request validation using [go-playground/validator](https://github.com/go-playground/validator)
- **Error Handling**: Consistent error handling with custom error types
- **Monitoring**: Integrated monitoring with [Sentry](https://sentry.io)

### Security
- **CORS**: Configurable CORS middleware
- **Rate Limiting**: IP-based and token-based rate limiting
- **Security Headers**: Secure headers middleware
- **Input Sanitization**: XSS protection and input sanitization
- **Password Hashing**: Secure password hashing with [Bcrypt](https://github.com/golang/crypto)

## Tech Stack

- **Web Framework**: [Gin](https://github.com/gin-gonic/gin) v1.9.1
- **ORM**: [GORM](https://gorm.io) v2.0.0
- **Database**: PostgreSQL 15
- **Caching**: Redis 7.0
- **Authentication**: [golang-jwt](https://github.com/golang-jwt/jwt) v5.0.0
- **Validation**: [go-playground/validator](https://github.com/go-playground/validator) v10.0.0
- **Logging**: [Uber Zap](https://github.com/uber-go/zap) v1.24.0
- **Configuration**: [GoDotEnv](https://github.com/joho/godotenv) v1.5.1
- **Documentation**: [Swag](https://github.com/swaggo/swag) v1.16.0
- **Testing**: [Testify](https://github.com/stretchr/testify) v1.8.0

## Prerequisites

- Go 1.21 or higher
- Docker 24.0 or higher
- Docker Compose v2.0 or higher
- PostgreSQL 15
- Redis 7.0

## Getting Started

### Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/go-boilerplate.git
cd go-boilerplate
```

2. Copy the environment file:
```bash
cp .env.example .env
```

3. Update the environment variables in `.env` file.

### Development with Docker

1. Build and start the containers:
```bash
docker-compose up -d
```

2. Watch logs:
```bash
docker-compose logs -f app
```

The application will be available at `http://localhost:6500`

### Local Development

1. Install dependencies:
```bash
go mod download
```

2. Install Air for hot reload:
```bash
go install github.com/cosmtrek/air@latest
```

3. Start the application with hot reload:
```bash
air
```

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Project Structure

```
.
├── cmd/                    # Application entry points
│   └── api/               # API server
│       └── main.go        # Main application file
├── config/                # Configuration
│   ├── config.go          # Configuration structs
│   └── database.go        # Database configuration
├── internal/              # Private application code
│   ├── delivery/          # HTTP handlers
│   │   ├── http/         # HTTP transport layer
│   │   └── dto/          # Data Transfer Objects
│   ├── domain/           # Business logic and entities
│   │   ├── entity/       # Domain entities
│   │   ├── repository/   # Repository interfaces
│   │   └── service/      # Service interfaces
│   ├── repository/       # Repository implementations
│   └── service/          # Service implementations
├── migrations/           # Database migrations
├── pkg/                  # Public libraries
│   ├── auth/            # Authentication package
│   ├── database/        # Database utilities
│   ├── logger/          # Logging utilities
│   └── validator/       # Validation utilities
├── scripts/             # Build and deployment scripts
├── test/                # Test utilities and fixtures
└── web/                # Web assets and templates
```

## Configuration

### Environment Variables

| Variable | Description | Default | Required |
|----------|-------------|---------|----------|
| APP_NAME | Application name | Golang Boilerplate | No |
| APP_ENV | Environment (development/production) | development | No |
| APP_PORT | HTTP server port | 8080 | No |
| APP_URL | Application URL | http://localhost:8080 | No |
| DB_HOST | Database host | 127.0.0.1 | Yes |
| DB_PORT | Database port | 5432 | Yes |
| DB_NAME | Database name | go_boilerplate | Yes |
| DB_USER | Database username | - | Yes |
| DB_PASSWORD | Database password | - | Yes |
| REDIS_HOST | Redis host | 127.0.0.1 | Yes |
| REDIS_PORT | Redis port | 6379 | Yes |
| REDIS_PASSWORD | Redis password | - | No |
| JWT_SECRET | JWT signing key | - | Yes |
| JWT_EXPIRY | JWT expiry time in hours | 24 | No |
| CORS_ORIGIN | Allowed CORS origins | * | No |
| LOG_LEVEL | Logging level (debug/info/warn/error) | info | No |

## Makefile Commands

This project includes a Makefile to help with common development tasks. Here are the available commands:

### Development Commands

```bash
# Start the application in development mode with hot reload
make dev

# Build the application
make build

# Run the application
make run

# Clean build artifacts
make clean
```

### Database Commands

```bash
# Create a new migration
make migrate-create name=migration_name

# Run all pending migrations
make migrate-up

# Rollback the last migration
make migrate-down

# Show migration status
make migrate-status
```

### Docker Commands

```bash
# Build and start all containers
make docker-up

# Stop all containers
make docker-down

# Build all containers
make docker-build

# Show container logs
make docker-logs

# Remove all containers and volumes
make docker-clean
```

### Testing Commands

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Run specific tests
make test-file file=path/to/test/file

# Run linter
make lint
```

### Tools and Dependencies

```bash
# Install development tools (air, swag, etc.)
make install-tools

# Update dependencies
make deps-update

# Tidy go.mod
make deps-tidy
```

### Documentation Commands

```bash
# Generate Swagger documentation
make swagger

# Generate and view test coverage
make coverage-html
```

### Helper Commands

```bash
# Show all available make commands
make help

# Format all Go files
make fmt

# Run security checks
make security-check

# Generate mocks for testing
make generate-mocks
```

Example usage:

```bash
# Start development with hot reload
make dev

# Create a new migration
make migrate-create name=add_users_table

# Run tests with coverage and view report
make test-coverage
make coverage-html
```

## API Documentation

### Swagger Documentation

This project uses Swagger for API documentation. Access the Swagger UI at:

```
http://localhost:6500/swagger/index.html
```

To generate/update Swagger documentation:

```bash
# Install swag
go install github.com/swaggo/swag/cmd/swag@latest

# Generate documentation
make swag
```

### Authentication Endpoints

#### Register User
```http
POST /api/v1/auth/register
Content-Type: application/json

{
    "name": "User Example",
    "username": "user123",
    "email": "user@example.com",
    "password": "password123"
}
```

#### Login
```http
POST /api/v1/auth/login
Content-Type: application/json

{
    "username": "user123",
    "password": "password123"
}
```

### User Endpoints

#### Get All Users
```http
GET /api/v1/users/page=1&perPage=5
Authorization: Bearer <token>
```

## Development

### Adding New Features

1. Create a new branch:
```bash
git checkout -b feature/your-feature-name
```

2. Implement your changes following the project structure
3. Add tests for new functionality
4. Update documentation if necessary
5. Submit a pull request

### Code Style

This project follows the official Go style guide and best practices:

- Use `gofmt` for code formatting
- Follow [Effective Go](https://golang.org/doc/effective_go) guidelines
- Run `golangci-lint` before committing

```bash
# Install golangci-lint
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run linter
golangci-lint run
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
