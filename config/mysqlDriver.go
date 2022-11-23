package config

import (
	"BE13/MVC/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	conesctionString := os.Getenv("DB_CONNECTION")

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
