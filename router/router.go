package router

import (
	"go-api/controller"

	"github.com/labstack/echo/v4"
)

func InitRouter(e *echo.Echo, userController *controller.UserController) {
	e.GET("/users", userController.GetAllUsers)
	e.POST("/users", userController.CreateUser)
	e.PUT("/users/:id", userController.UpdateUser)
	e.DELETE("/users/:id", userController.DeleteUser)
	e.GET("/users/openlibrary", userController.GetOpenLibraryData)
}