package middlewares

import (
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
)

func Swagger() fiber.Handler {
	return swagger.New(swagger.Config{
		BasePath: "/api/v1",
		FilePath: "./docs/swagger.json",
		Path:     "/swagger",
		Title:    "Go Fiber boilerplate API",
		CacheAge: 0,
	})
}
