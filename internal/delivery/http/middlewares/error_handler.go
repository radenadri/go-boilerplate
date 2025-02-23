package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/radenadri/go-boilerplate/internal/delivery/dto/response"
)

func NotFoundHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotFound, response.Response{
			Success: false,
			Error:   "URL not found",
		})
	}
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Only handle errors if there are any
		if len(c.Errors) > 0 {
			c.JSON(http.StatusInternalServerError, response.Response{
				Success: false,
				Error:   "Internal Server Error",
			})
			return
		}
	}
}
