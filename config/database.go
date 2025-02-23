package config

import (
	"fmt"
	"strconv"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var RedisClient *redis.Client

func InitDB() {
	var dbIsSSLEnabled string

	isSSLEnabled, parseErr := strconv.ParseBool(DBSSLEnabled)

	if parseErr != nil {
		panic("Failed to parse DBSSLEnabled!")
	}

	if isSSLEnabled {
		dbIsSSLEnabled = "enable"
	} else {
		dbIsSSLEnabled = "disable"
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		DBHost, DBUsername, DBPassword, DBDatabase, DBPort, dbIsSSLEnabled,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
}

func InitRedis() {

	redisDB := RedisDB
	db, err := strconv.Atoi(redisDB)
	if err != nil {
		db = 0
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", RedisHost, RedisPort),
		Password:     RedisPassword,
		DB:           db,
		PoolSize:     10,
		MinIdleConns: 5,
	})
}
