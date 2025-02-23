package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/radenadri/go-boilerplate/config"
	"github.com/radenadri/go-boilerplate/internal/delivery/dto/response"
)

func AuthenticateJWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		secretKey := config.JWTSecret

		if secretKey == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Response{
				Success: false,
				Error:   "JWT_SECRET not found",
			})
			return
		}

		authToken := c.Request.Header.Get("Authorization")
		if authToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Response{
				Success: false,
				Error:   "Authorization header not found",
			})
			return
		}

		token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, response.Response{
				Success: false,
				Error:   "Invalid token",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
