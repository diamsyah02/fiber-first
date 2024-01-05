package configs

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnDB() {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/latihan?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}
