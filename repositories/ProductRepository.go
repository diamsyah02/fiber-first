package repositories

import (
	"fiber-first/configs"
	"fiber-first/models"
)

func GetAllProduct() ([]models.Product, error) {
	product := []models.Product{}
	db := configs.DB
	result := db.Model(&models.Product{}).Find(&product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func GetProductById(id int) (models.Product, error) {
	product := models.Product{}
	db := configs.DB
	result := db.Model(&models.Product{}).Where("id = ?", id).Find(&product)
	if result.Error != nil {
		return product, result.Error
	}
	return product, nil
}

func StoreProduct(json *models.Product) error {
	product := models.Product{
		Name:  json.Name,
		Qty:   json.Qty,
		Price: json.Price,
	}
	db := configs.DB
	result := db.Create(&product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateProduct(id int, json *models.Product) error {
	product := models.Product{
		Name:  json.Name,
		Qty:   json.Qty,
		Price: json.Price,
	}
	db := configs.DB
	result := db.Model(&models.Product{}).Where("id = ?", id).Updates(&product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteProduct(id int) error {
	product := models.Product{}
	db := configs.DB
	result := db.Where("id = ?", id).Delete(&product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
