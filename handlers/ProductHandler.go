package handlers

import (
	"fiber-first/models"
	"fiber-first/services"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllProduct(c *fiber.Ctx) error {
	code, msg, data := services.GetAllProduct()
	return c.JSON(fiber.Map{
		"code":    code,
		"message": msg,
		"data":    data,
	})
}

func GetProductById(c *fiber.Ctx) error {
	param := c.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		log.Fatal(err)
	}
	code, msg, data := services.GetProductById(id)
	return c.JSON(fiber.Map{
		"code":    code,
		"message": msg,
		"data":    data,
	})
}

func StoreProduct(c *fiber.Ctx) error {
	json := new(models.Product)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}
	code, msg := services.StoreProduct(json)
	return c.JSON(fiber.Map{
		"code":    code,
		"message": msg,
	})
}

func UpdateProduct(c *fiber.Ctx) error {
	json := new(models.Product)
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
		})
	}
	param := c.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		log.Fatal(err)
	}
	code, msg := services.UpdateProduct(id, json)
	return c.JSON(fiber.Map{
		"code":    code,
		"message": msg,
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	param := c.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		log.Fatal(err)
	}
	code, msg := services.DeleteProduct(id)
	return c.JSON(fiber.Map{
		"code":    code,
		"message": msg,
	})
}
