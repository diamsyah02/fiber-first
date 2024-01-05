package routes

import (
	"fiber-first/handlers"
	"fiber-first/middleware"

	"github.com/gofiber/fiber/v2"
)

func ProductRoute(app *fiber.App) {
	group := app.Group("/product", middleware.Authentication)
	group.Get("/", handlers.GetAllProduct)
	group.Get("/:id", handlers.GetProductById)
	group.Post("/", handlers.StoreProduct)
	group.Put("/:id", handlers.UpdateProduct)
	group.Delete("/:id", handlers.DeleteProduct)
}
