package model

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

func (user *User) GenerateJwt() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.UserId,
		"email":    user.Email,
		"username": user.Username,
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("jwt_string")))
	fmt.Println(err)
	// fmt.Println(tokenString, err)

	return tokenString
}
