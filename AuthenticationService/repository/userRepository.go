package repository

import (
	"authenticationservice/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (userRepository *UserRepository) CreateUser(user *model.User) {

	userRepository.db.Create(user)
}
