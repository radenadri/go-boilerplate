# Go Boilerplate

A robust and scalable Go boilerplate for building modern web applications with best practices and essential features pre-configured. This boilerplate follows clean architecture principles and includes production-ready features out of the box.

## Table of Contents

- [Features](#features)
- [Tech Stack](#tech-stack)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [Project Structure](#project-structure)
- [Configuration](#configuration)
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
- **Metrics**: Prometheus metrics integration for monitoring
- **Tracing**: Distributed tracing with OpenTelemetry

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
swag init -g cmd/api/main.go
```

### Authentication Endpoints

#### Register User
```http
POST /api/v1/auth/register
Content-Type: application/json

{
    "username": "user123",
    "email": "user@example.com",
    "password": "password123"
}
```

Response:
```json
{
    "status": "success",
    "message": "User registered successfully",
    "data": {
        "id": "uuid",
        "username": "user123",
        "email": "user@example.com",
        "created_at": "2024-02-24T12:00:00Z"
    }
}
```

#### Login
```http
POST /api/v1/auth/login
Content-Type: application/json

{
    "email": "user@example.com",
    "password": "password123"
}
```

Response:
```json
{
    "status": "success",
    "message": "Login successful",
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIs...",
        "token_type": "Bearer",
        "expires_in": 86400
    }
}
```

### User Endpoints

#### Get All Users
```http
GET /api/v1/users
Authorization: Bearer <token>
```

Response:
```json
{
    "status": "success",
    "data": {
        "users": [
            {
                "id": "uuid",
                "username": "user123",
                "email": "user@example.com",
                "created_at": "2024-02-24T12:00:00Z"
            }
        ],
        "pagination": {
            "current_page": 1,
            "per_page": 10,
            "total": 100
        }
    }
}
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
