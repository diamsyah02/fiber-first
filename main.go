package main

import (
	"fiber-first/configs"
	"fiber-first/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(
		cors.Config{
			AllowOrigins: "*",
			AllowHeaders: "Origin, Content-Type, Accept",
		}))

	configs.ConnDB()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Rest Api with Fiber Golang")
	})

	routes.AuthRoute(app)
	routes.ProductRoute(app)

	app.Listen(":3000")
}
