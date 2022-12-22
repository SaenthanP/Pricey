package database

import (
	"os"

	"jobservice/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

func SetDB() (db *gorm.DB) {
	godotenv.Load()
	db, err := gorm.Open(postgres.Open(os.Getenv("CONNECTION_STRING")), &gorm.Config{})
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	db.AutoMigrate(&model.Job{})
	_ = err

	return db
}
