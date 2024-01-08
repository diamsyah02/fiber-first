package services

import (
	"fiber-first/models"
	"fiber-first/repositories"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetAllProduct() (int, string, []models.Product) {
	result, err := repositories.GetAllProduct()
	if err != nil {
		log.Fatal(err)
		return fiber.StatusBadRequest, err.Error(), nil
	}
	return fiber.StatusOK, "Get data product success", result
}

func GetProductById(id int) (int, string, models.Product) {
	result, err := repositories.GetProductById(id)
	if err != nil {
		log.Fatal(err)
		return fiber.StatusBadRequest, err.Error(), models.Product{}
	}
	return fiber.StatusOK, "Get data product success", result
}

func StoreProduct(json *models.Product) (int, string) {
	err := repositories.StoreProduct(json)
	if err != nil {
		log.Fatal(err)
		return fiber.StatusBadRequest, err.Error()
	}
	return fiber.StatusOK, "Insert data product success"
}

func UpdateProduct(id int, json *models.Product) (int, string) {
	err := repositories.UpdateProduct(id, json)
	if err != nil {
		log.Fatal(err)
		return fiber.StatusBadRequest, err.Error()
	}
	return fiber.StatusOK, "Update data product success"
}

func DeleteProduct(id int) (int, string) {
	err := repositories.DeleteProduct(id)
	if err != nil {
		log.Fatal(err)
		return fiber.StatusBadRequest, err.Error()
	}
	return fiber.StatusOK, "Delete data product success"
}
