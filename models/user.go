package models

import (
	"BE13/MVC/entities"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Phone    string
	Address  string `gorm:"type:varchar(15)"`
	Books    []Book
}

func ModelstoCore(user User) entities.UserCore {
	return entities.UserCore{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		Address:   user.Address,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

}

func CoreToModel(dataCore entities.UserCore) User {
	userGorm := User{

		Name:     dataCore.Name,
		Email:    dataCore.Email,
		Password: dataCore.Password,
		Phone:    dataCore.Phone,
		Address:  dataCore.Address,
	}
	return userGorm
}

func ListModeltoCore(user []User) []entities.UserCore {
	var dataCore []entities.UserCore
	for _, v := range user {
		dataCore = append(dataCore, ModelstoCore(v))
	}
	return dataCore
}
