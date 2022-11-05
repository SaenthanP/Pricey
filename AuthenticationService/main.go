package main

import (
	"authenticationservice/database"
	"authenticationservice/handler"
	"authenticationservice/repository"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name    string
	NewName string
}

func main() {
	db := database.SetDB()

	userRepository:=repository.NewUserRepository(db)

	test:=handler.NewUserHandler(userRepository)
	
	router := gin.Default()

	router.GET("/api/ping", test.RegisterUser)

	router.Run()
}
