# Go Boilerplate

A robust and scalable Go boilerplate for building modern web applications with best practices and essential features pre-configured.

## Features

- **Clean Architecture**: Follows clean architecture principles for better code organization and maintainability
- **RESTful API**: Built with Gin framework for high-performance HTTP routing
- **Database Integration**: PostgreSQL with GORM for reliable data persistence
- **Redis Support**: Integrated Redis for caching and rate limiting
- **JWT Authentication**: Secure authentication using JWT tokens
- **Environment Configuration**: Easy environment management with .env files
- **Docker Support**: Containerization ready with Docker and docker-compose
- **Logging**: Structured logging with Zap logger
- **Input Validation**: Request validation using go-playground/validator
- **Error Handling**: Consistent error handling across the application
- **Middleware**: CORS, logging, and other essential middleware included
- **Database Migrations**: Organized database schema management

## Prerequisites

- Go 1.21 or higher
- Docker and Docker Compose
- PostgreSQL
- Redis

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

3. Update the environment variables in `.env` file with your configuration.

### Running with Docker

1. Build and start the containers:
```bash
docker-compose up -d
```

The application will be available at `http://localhost:6500`

### Running Locally

1. Install dependencies:
```bash
go mod download
```

2. Start the application:
```bash
go run cmd/api/main.go
```

## Project Structure

```
.
├── cmd/                # Application entry points
├── config/            # Configuration setup
├── internal/          # Private application code
│   ├── delivery/      # HTTP handlers and DTOs
│   ├── domain/        # Business logic and entities
│   ├── repositories/  # Data access layer
│   └── services/      # Business logic implementation
├── migrations/        # Database migrations
├── pkg/               # Public libraries
└── utils/             # Utility functions
```

## Environment Variables

| Variable | Description | Default |
|----------|-------------|----------|
| APP_NAME | Application name | Golang Boilerplate |
| APP_ENV | Environment (development/production) | development |
| APP_PORT | HTTP server port | 8080 |
| DB_HOST | Database host | 127.0.0.1 |
| DB_PORT | Database port | 5432 |
| DB_DATABASE | Database name | go_boilerplate |
| DB_USERNAME | Database username | - |
| DB_PASSWORD | Database password | - |
| REDIS_HOST | Redis host | 127.0.0.1 |
| REDIS_PORT | Redis port | 6379 |
| JWT_SECRET | JWT signing key | - |

## API Documentation

### Swagger Documentation

This project uses Swagger for API documentation. You can access the Swagger UI interface at:

```
http://localhost:8080/swagger/index.html
```

The Swagger documentation provides:
- Detailed API endpoint descriptions
- Request/response schemas
- Authentication requirements
- Interactive API testing interface

### Authentication

#### Register User
```http
POST /api/v1/auth/register
Content-Type: application/json

{
    "username": "user123",
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

### Users

#### Get All Users
```http
GET /api/v1/users
Authorization: Bearer <token>
```

#### Get User by Username
```http
GET /api/v1/users/:username
Authorization: Bearer <token>
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
