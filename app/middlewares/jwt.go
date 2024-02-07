package middlewares

import (
	"boilerplate/app/config"
	"boilerplate/app/models"
	"boilerplate/app/utils"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func EnableJWT() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(config.JWT_SECRET)},
		ErrorHandler: JwtError,
	})
}

func GenerateJWT(user models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["user_id"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	t, err := token.SignedString([]byte(config.JWT_SECRET))
	if err != nil {
		return "error", err
	}

	return t, nil
}

func JwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(utils.Response{
				Status:  false,
				Message: "Missing or malformed JWT",
				Data:    nil,
			})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(utils.Response{
			Status:  false,
			Message: "Invalid or expired JWT",
			Data:    nil,
		})
}
