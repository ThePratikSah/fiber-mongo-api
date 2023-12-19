package main

import (
	"fiber-mongo-api/src/configs"
	"fiber-mongo-api/src/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// Fiber instance
	app := fiber.New()

	// test route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Ok")
	})

	// connect to mongodb
	configs.ConnectDB()

	routes.UserRoutes(app)

	// Start server
	app.Listen(":3000")
}
