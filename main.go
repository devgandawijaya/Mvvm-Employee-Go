package main

import (
	"go-api/config"
	"go-api/controller"
	"go-api/repository"
	"go-api/router"
	"go-api/service"

	"github.com/labstack/echo/v4"
)

func main() {
	config.InitDB()
	e := echo.New()
	userRepo := repository.NewUserRepository(config.DB)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)
	router.InitRouter(e, userController)
	e.Logger.Fatal(e.Start(":8080"))
}