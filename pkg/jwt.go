package pkg

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/radenadri/go-boilerplate/config"
	"github.com/radenadri/go-boilerplate/internal/domain/models"
)

func GenerateJWT(user models.User) (string, error) {
	claims := &jwt.MapClaims{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWTSecret))
}
