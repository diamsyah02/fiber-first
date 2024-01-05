package repositories

import (
	"fiber-first/configs"
	"fiber-first/models"
	"fiber-first/utils"
)

func CheckUser(username string) (*models.User, error) {
	user := models.User{}
	query := models.User{Username: username}
	db := configs.DB
	result := db.First(&user, &query)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func Register(json *models.User) error {
	user := models.User{Username: json.Username, Password: utils.HashAndSalt([]byte(json.Password))}
	db := configs.DB
	result := db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
