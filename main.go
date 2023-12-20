package main

import (
	"fiber-mongo-api/src/configs"
	"fiber-mongo-api/src/routes"
	"flag"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// Fiber instance
	app := fiber.New()

	// health test route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Ok")
	})

	// connect to mongodb
	configs.ConnectDB()

	seed := flag.Bool("seed", false, "🌱 Seeding the database")

	flag.Parse()
	if *seed {
		configs.SeedData()
	}

	// register routes here
	routes.UserRoutes(app)

	// Start server
	app.Listen(":3000")
}
