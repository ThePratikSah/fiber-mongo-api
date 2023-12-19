package routes

import (
	"fiber-mongo-api/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	app.Post("/users", controllers.CreateUser)
	app.Get("/users", controllers.GetAllUsers)
	app.Put("/users/:id", controllers.UpdateUser)
}
