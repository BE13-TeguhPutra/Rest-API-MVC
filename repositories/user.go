package repositories

import (
	"BE13/MVC/config"
	"BE13/MVC/entities"
	"BE13/MVC/models"
	"errors"
)

func GetAlluser() ([]entities.UserCore, error) {
	var users []models.User

	tx := config.DB.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	dataCore := models.ListModeltoCore(users)

	// var dataCore []entities.UserCore
	// for _, v := range users {
	// 	dataCore = append(dataCore, entities.UserCore{
	// 		ID:        v.ID,
	// 		Name:      v.Name,
	// 		Email:     v.Email,
	// 		Phone:     v.Phone,
	// 		Address:   v.Address,
	// 		CreatedAt: v.CreatedAt,
	// 		UpdatedAt: v.UpdatedAt,
	// 	})

	// }
	return dataCore, nil

}

func InsertUser(Core entities.UserCore) error {

	user := models.CoreToModel(Core)
	tx := config.DB.Create(&user) // proses insert data
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}

	return nil

}

func UpdateUser(userCore entities.UserCore, id int) error {

	user := models.CoreToModel(userCore)
	tx := config.DB.Model(user).Where("id = ?", id).Updates(user)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("id not found")
	}
	return nil

}

func Getbyid(id int) (entities.UserCore, error) {

	user := models.User{}
	tx := config.DB.First(&user, id)
	if tx.Error != nil {
		return entities.UserCore{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return entities.UserCore{}, errors.New("id not found")

	}
	dCore := models.ModelstoCore(user)
	return dCore, nil
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
