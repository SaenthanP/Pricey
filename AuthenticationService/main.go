package main

import (
	"os"

	"authenticationservice/handler"
	"authenticationservice/model"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

type User struct {
	Name    string
	NewName string
}

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("CONNECTION_STRING")), &gorm.Config{})
	db.AutoMigrate(&model.User{})
	_ = err
	router := gin.Default()

	router.GET("/api/ping", handler.RegisterUser)

	router.Run()
}
