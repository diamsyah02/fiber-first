package middleware

import (
	"fiber-first/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Authentication(c *fiber.Ctx) error {
	authorizationHeader := c.Request().Header.Peek("Authorization")
	if !strings.Contains(string(authorizationHeader), "Bearer") {
		return c.JSON(fiber.Map{
			"code":    401,
			"message": "Invalid Token",
		})
	}
	tokenString := strings.Replace(string(authorizationHeader), "Bearer ", "", -1)
	ok := utils.VerifyToken(tokenString)
	if ok {
		return c.Next()
	}
	return c.JSON(fiber.Map{
		"code":    401,
		"message": "Invalid Token",
	})
}
