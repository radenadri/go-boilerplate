package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	AppName       = GetEnvOrDefault("APP_NAME", "Go Boilerplate")
	AppEnv        = GetEnvOrDefault("APP_ENV", "development")
	AppPort       = GetEnvOrDefault("APP_PORT", "8080")
	AppTimezone   = GetEnvOrDefault("APP_TIMEZONE", "UTC")
	AppApiVersion = GetEnvOrDefault("APP_API_VERSION", "v1")

	CORSAllowOrigins = GetEnvOrDefault("CORS_ALLOWED_ORIGINS", "*")

	DBHost       = GetEnvOrDefault("DB_HOST", "localhost")
	DBPort       = GetEnvOrDefault("DB_PORT", "5432")
	DBDatabase   = GetEnvOrDefault("DB_DATABASE", "")
	DBUsername   = GetEnvOrDefault("DB_USERNAME", "")
	DBPassword   = GetEnvOrDefault("DB_PASSWORD", "")
	DBSSLEnabled = GetEnvOrDefault("DB_SSL_ENABLED", "false")

	JWTSecret = GetEnvOrDefault("JWT_SECRET", "")

	RedisHost     = GetEnvOrDefault("REDIS_HOST", "localhost")
	RedisPort     = GetEnvOrDefault("REDIS_PORT", "6379")
	RedisPassword = GetEnvOrDefault("REDIS_PASSWORD", "")
	RedisDB       = GetEnvOrDefault("REDIS_DB", "0")

	SentryDSN = GetEnvOrDefault("SENTRY_DSN", "")
)

func GetEnvOrDefault(key, defaultValue string) string {
	if err := godotenv.Load(); err != nil {
		println("Warning: Error loading .env file:", err)
	}

	value := os.Getenv(key)
	if value != "" {
		return value
	}

	return defaultValue
}
