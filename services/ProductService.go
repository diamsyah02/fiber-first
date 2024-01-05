package services

import (
	"fiber-first/models"
	"fiber-first/repositories"
	"log"
)

func GetAllProduct() (int, string, []models.Product) {
	result, err := repositories.GetAllProduct()
	if err != nil {
		log.Fatal(err)
	}
	return 200, "Get data product success", result
}

func GetProductById(id int) (int, string, models.Product) {
	result, err := repositories.GetProductById(id)
	if err != nil {
		log.Fatal(err)
	}
	return 200, "Get data product success", result
}

func StoreProduct(json *models.Product) (int, string) {
	err := repositories.StoreProduct(json)
	if err != nil {
		return 400, err.Error()
	}
	return 200, "Insert data product success"
}

func UpdateProduct(id int, json *models.Product) (int, string) {
	err := repositories.UpdateProduct(id, json)
	if err != nil {
		return 400, err.Error()
	}
	return 200, "Update data product success"
}

func DeleteProduct(id int) (int, string) {
	err := repositories.DeleteProduct(id)
	if err != nil {
		return 400, err.Error()
	}
	return 200, "Delete data product success"
}
