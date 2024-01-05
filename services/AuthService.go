package services

import (
	"fiber-first/models"
	"fiber-first/repositories"
	"fiber-first/utils"
	"log"
)

func Login(json *models.User) (int, string, string) {
	check, err := repositories.CheckUser(json.Username)
	if err != nil {
		return 400, "Login failed, because username has not been registered", ""
	}
	result := utils.ComparePass(check.Password, []byte(json.Password))
	if result {
		token := utils.GenerateToken(json.Username)
		return 200, "Login success", token
	}
	return 400, "Login failed", ""
}

func Register(json *models.User) (int, string) {
	_, err_ := repositories.CheckUser(json.Username)
	if err_ == nil {
		return 400, "Register failed, because username is already exist!"
	}
	err := repositories.Register(json)
	if err != nil {
		log.Fatal(err)
		return 400, "Register failed"
	}
	return 200, "Register success"
}
