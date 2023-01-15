package repository

import (
	"authenticationservice/model"
	"time"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (userRepository *UserRepository) CreateUser(user *model.User) {

	userRepository.db.FirstOrCreate(user)
}

func (userRepository *UserRepository) GetUserByEmail(email string) *model.User {
	var result model.User
	err := userRepository.db.Model(&model.User{}).Where("email=?", email).First(&result).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	} else {
		return &result
	}

}

func (userRepository *UserRepository) LoginUser(email string) *model.User {

	userRepository.db.Model(&model.User{}).Where("email=?", email).Update("last_login", time.Now()).Debug()
	user := userRepository.GetUserByEmail(email)

	return user
}
