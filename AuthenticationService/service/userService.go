package service

import (
	"authenticationservice/dto"
	"authenticationservice/model"
	"authenticationservice/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserSevice(userRepository *repository.UserRepository) *UserService {
	return &UserService{userRepository}
}

func (userService *UserService) CreateUser(userToCreate *dto.CreateUserDto) {

	user := &model.User{}
	user.Email = userToCreate.Email
	user.Username = userToCreate.Username
	userService.userRepository.CreateUser(user)
}
