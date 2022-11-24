package config

import (
	"BE13/MVC/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	conesctionString := "root:Teguh12345@tcp(127.0.0.1:3306)/db_be13_gorm?parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(conesctionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

}
func InitialMigration() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Book{})
}
