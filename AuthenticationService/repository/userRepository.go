package repository

import (
	"gorm.io/gorm"
	"authenticationservice/model"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}


func (userRepository *UserRepository) CreateUser(){
	user:=&model.User{}

	user.Email="saenthan"
	userRepository.db.Create(user)
}