package handler

import (
	"authenticationservice/repository"

	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userRepository *repository.UserRepository
}

func NewUserHandler(userRepository *repository.UserRepository) *UserHandler {
	return &UserHandler{userRepository}
}
func (userHandler *UserHandler) RegisterUser(c *gin.Context) {

	userHandler.userRepository.CreateUser()
	c.JSON(http.StatusOK, gin.H{
		"message": "sucess",
	})
}
