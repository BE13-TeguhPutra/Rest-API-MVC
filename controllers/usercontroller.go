package controllers

import (
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
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("Success Get All Users", users))

}

func PostUserController(c echo.Context) error {
	user := models.User{}
	errBind := c.Bind(&user)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error binding"))
	}
	err := repositories.InsertUser(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed to Create User"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("Create User Success"))

}

func PutUserController(c echo.Context) error {
	user := models.User{}
	idParam := c.Param("id")
	id, errconv := strconv.Atoi(idParam)
	if errconv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error convert"))
	}

	errBind := c.Bind(&user)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error binding"))
	}

	errupdate := repositories.UpdateUser(user, id)
	if errupdate != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Update Failed"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("Update Success"))

}

func GetUserbyIdController(c echo.Context) error {
	user := models.User{}
	idParam := c.Param("id")
	id, errconv := strconv.Atoi(idParam)
	if errconv != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error convert"))
	}

	res, errbyid := repositories.Getbyid(user, id)
	if errbyid != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed to get user"))
	}
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
