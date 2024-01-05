package routes

import (
	"fiber-first/handlers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(app *fiber.App) {
	app.Post("/login", handlers.Login)
	app.Post("/register", handlers.Register)
}
