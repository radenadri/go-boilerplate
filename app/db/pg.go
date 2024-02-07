package db

import (
	"boilerplate/app/config"
	"boilerplate/app/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME,
		config.DB_USER,
		config.DB_PASSWORD,
		config.DB_SSL_MODE,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		panic("failed to connect database")
	}

	DB = db
}

func Migrate() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		panic("failed to migrate database")
	}
}
