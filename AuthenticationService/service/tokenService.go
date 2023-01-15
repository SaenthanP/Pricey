package service

import (
	"authenticationservice/model"
	"fmt"
	"log"
	"os"

	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateJwt(user *model.User) string {
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

func VerifyJwt(tknStr string) (bool, uuid.UUID) {
	type Claims struct {
		id       uuid.UUID
		email    string
		username string
		jwt.RegisteredClaims
	}

	tkn, err := jwt.ParseWithClaims(tknStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("jwt_string")), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			log.Panic("Token Error: ErrSignatureInvalid")
			return false, uuid.Nil
		}

	}
	if !tkn.Valid {
		log.Panic("Token Error: Not valid ")
		return false, uuid.Nil
	}

	claims := tkn.Claims.(*Claims)
	return true, claims.id
}
