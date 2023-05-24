package main

import (
	"go_fiber/app"
	"go_fiber/controller"

	"go_fiber/repository"
	"go_fiber/route"
	"go_fiber/service"

	"github.com/go-playground/validator/v10"

)

func main(){
	var db = app.NewDB()
	
	validator := validator.New()
	userRepository := repository.NewUserRepository()
    userService := service.NewUserService(userRepository, db, validator)
    userController := controller.NewUserController(userService)
	
	app := route.RouteInit(userController)

	app.Listen(":3000")
}

