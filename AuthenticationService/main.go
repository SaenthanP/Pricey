package main

import (
	"authenticationservice/config"
	"authenticationservice/database"
	"authenticationservice/handler"
	"authenticationservice/repository"
	"authenticationservice/service"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.SetDB()

	fmt.Println(os.Getenv("CONNECTION_STRING"))
	config := config.NewConfig()
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserSevice(userRepository)
	userHandler := handler.NewServer(userService, config)

	router := gin.Default()

	router.GET("/api/ping", userHandler.RegisterUser)
	router.GET("/api/callback", userHandler.Callback)
	router.GET("/api/verify", userHandler.VerifyToken)

	router.Run(":8080")
}
