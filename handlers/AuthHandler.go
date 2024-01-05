package handlers

import (
	"fiber-first/models"
	"fiber-first/services"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	var token string
	json := new(models.User)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}
	code, msg, token := services.Login(json)
	return c.JSON(fiber.Map{
		"code":    code,
		"message": msg,
		"data": M{
			"token": token,
		},
	})
}

func Register(c *fiber.Ctx) error {
	json := new(models.User)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}
	code, msg := services.Register(json)
	return c.JSON(fiber.Map{
		"code":    code,
		"message": msg,
		"data":    json,
	})
}

type M map[string]interface{}
