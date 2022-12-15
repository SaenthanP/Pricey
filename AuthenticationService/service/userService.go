package service

import (
	"authenticationservice/dto"
	"authenticationservice/model"
	"authenticationservice/repository"
	"fmt"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserSevice(userRepository *repository.UserRepository) *UserService {
	return &UserService{userRepository}
}

func (userService *UserService) CreateUser(userToCreate *dto.CreateUserDto) string {

	user := &model.User{}
	user.Email = userToCreate.Email
	user.Username = userToCreate.Username

	if userService.userRepository.GetUserByEmail(user.Email) == nil {
		userService.userRepository.CreateUser(user)
	}

	user = userService.userRepository.LoginUser(user.Email)
	fmt.Println(user.UserId)
	token := user.GenerateJwt()
	fmt.Println(token)

	return token
}
