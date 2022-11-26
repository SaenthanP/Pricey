package main

import (
	"authenticationservice/config"
	"authenticationservice/database"
	"authenticationservice/handler"
	"authenticationservice/repository"
	"authenticationservice/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.SetDB()
	config := config.NewConfig()
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserSevice(userRepository)
	userHandler := handler.NewServer(userService, config)

	router := gin.Default()

	router.GET("/api/ping", userHandler.RegisterUser)
	router.GET("/callback", userHandler.Callback)

	router.Run(":8080")
}
