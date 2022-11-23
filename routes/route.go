package routes

import (
	"BE13/MVC/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {

	e := echo.New()
	e.GET("/users", controllers.GetUserController)
	e.GET("/users/:id", controllers.GetUserbyIdController)
	e.POST("/users", controllers.PostUserController)
	e.PUT("/users/:id", controllers.PutUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)

	return e

}
