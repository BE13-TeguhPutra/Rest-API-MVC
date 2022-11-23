package repositories

import (
	"BE13/MVC/config"
	"BE13/MVC/models"
	"errors"
)

func GetAlluser() ([]models.User, error) {
	var users []models.User

	tx := config.DB.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return users, nil

}

func InsertUser(user models.User) error {
	tx := config.DB.Create(&user) // proses insert data
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}

	return nil

}

func UpdateUser(user models.User, id int) error {

	tx := config.DB.Model(user).Where("id = ?", id).Updates(user)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("id not found")
	}
	return nil

}

func Getbyid(user models.User, id int) (any, error) {

	tx := config.DB.First(&user, id)
	if tx.Error != nil {
		return nil, tx.Error

	}
	if tx.RowsAffected == 0 {
		return nil, errors.New("id not found")

	}
	return user, nil
}

func RemoveUser(user models.User, id int) error {
	tx := config.DB.Delete(&user, id)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("id not found")

	}
	return nil

}
