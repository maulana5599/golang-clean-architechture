package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseConnection() {
	var err error
	dsn := "maulana:maulana186@tcp(localhost:3306)/voting?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Connection refused")
		log.Println(err.Error())
	}

	fmt.Println("Connection successfully")
}
