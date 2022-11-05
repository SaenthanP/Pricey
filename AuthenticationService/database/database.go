package database

import (
	"os"

	"authenticationservice/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

func SetDB() (db *gorm.DB) {
	godotenv.Load()
	db, err := gorm.Open(postgres.Open(os.Getenv("CONNECTION_STRING")), &gorm.Config{})
	db.AutoMigrate(&model.User{})
	_ = err

	return db
}
