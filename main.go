package main

import (
	"boilerplate/app/config"
	pg "boilerplate/app/db"
	"boilerplate/app/middlewares"
	"boilerplate/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	// Init database
	pg.Init()

	// Migrate database
	pg.Migrate()

	// Create new Fiber instance
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		},
	})

	// Enable CORS
	app.Use(middlewares.Cors())

	// Enable limiter
	app.Use(middlewares.Limiter())

	// Enable logger
	app.Use(middlewares.Logger())

	// Create route for "/"
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	// Create route for "/api"
	api := app.Group("/api")

	// Create route for "/api/v1"
	v1 := api.Group("/v1")

	// Create routes for "/api/v1"
	routes.AuthRoute(v1)
	routes.UserRoute(v1)

	// Custom 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Sorry can't find that!")
	})

	// Start server
	var host = config.APP_HOST
	var port = config.APP_PORT

	err := app.Listen(host + ":" + port)

	if err != nil {
		return
	}
}
