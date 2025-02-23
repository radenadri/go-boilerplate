package repositories

import (
	"github.com/radenadri/go-boilerplate/internal/domain/models"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	DB *gorm.DB
}

type UserRepository interface {
	FindAll(limit int, offset int) ([]models.User, error)
	FindByUsername(username string) (*models.User, error)
	Create(user *models.User) error
	Count() (int64, error)
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return &GormUserRepository{DB: DB}
}

func (r *GormUserRepository) FindAll(limit int, offset int) ([]models.User, error) {
	var results []models.User

	if err := r.DB.Limit(limit).Offset(offset).Find(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
}

func (r *GormUserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User

	if err := r.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *GormUserRepository) Create(user *models.User) error {
	if err := r.DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *GormUserRepository) Count() (int64, error) {
	var totalItems int64

	if err := r.DB.Model(&models.User{}).Count(&totalItems).Error; err != nil {
		return 0, err
	}

	return totalItems, nil
}
