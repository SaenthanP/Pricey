package service

import (
	"authenticationservice/model"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v4"
)
type user struct {
    *model.User
}

func (user user) GenerateJwt(){
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"id":user.UserId,
		"email":user.Email,
		"username":user.Username,
	})
	tokenString,err:=token.SignedString(os.Getenv("jwt_string"))

	fmt.Println(tokenString,err)
}