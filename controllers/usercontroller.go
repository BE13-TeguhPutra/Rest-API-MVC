package controllers

import (
	"BE13/MVC/entities"
	"BE13/MVC/helper"
	"BE13/MVC/models"
	"BE13/MVC/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetUserController(c echo.Context) error {
	users, err := repositories.GetAlluser()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}
	response := entities.ListUserCoreToResponse(users)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success Get All Users", response))

}

func PostUserController(c echo.Context) error {
	//Mengambil data dari form/body postman menggunakan bind, lalu dimasukkan ke struct user request
	user := entities.UserRequest{}
	errBind := c.Bind(&user)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error binding"))
	}
	//Data yang ada di struct request akan dimasukan ke struct Core detail ada di entities.
	dataCore := entities.UserRequestToCore(user)

	//Data yang ada di Struct UserCore dimasukan ke Models(database) liha repositori dan models.
	err := repositories.InsertUser(dataCore)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed to Create User"))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Create User Success"))

}

func PutUserController(c echo.Context) error {

	idParam := c.Param("id")
	id, errconv := strconv.Atoi(idParam)
	if errconv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error convert"))
	}
	userReq := entities.UserRequest{}
	errBind := c.Bind(&userReq)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error binding"))
	}

	userCore := entities.UserRequestToCore(userReq)

	errupdate := repositories.UpdateUser(userCore, id)
	if errupdate != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Update Failed"))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("Update Success"))

}

func GetUserbyIdController(c echo.Context) error {
	idParam := c.Param("id")
	id, errconv := strconv.Atoi(idParam)
	if errconv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error convert"))
	}

	dataCore, errGet := repositories.Getbyid(id)
	if errGet != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error Get by id"))
	}

	res := entities.UserCoretoRespone(dataCore)
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Get user by id success", res))

}

func DeleteUserController(c echo.Context) error {
	user := models.User{}
	idParam := c.Param("id")
	id, errconv := strconv.Atoi(idParam)
	if errconv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error convert"))
	}

	errdelete := repositories.RemoveUser(user, id)
	if errdelete != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Delete Failed"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("Delete Success"))
}
