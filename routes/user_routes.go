package routes

import (
	"boilerplate/app/controllers"
	"boilerplate/app/middlewares"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app fiber.Router) {
	app.Get("/users", middlewares.EnableJWT(), controllers.GetUsers)
	app.Get("/users/:id", middlewares.EnableJWT(), controllers.GetUser)
}
