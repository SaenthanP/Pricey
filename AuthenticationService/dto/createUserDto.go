package dto

import (
	"authenticationservice/model"
	"fmt"
)

type CreateUserDto struct {
	Email    string
	Username string
}

func ToCreateUserDto(oauthResponse model.OauthResponse) *CreateUserDto {
	fmt.Println(oauthResponse.Email)
	return &CreateUserDto{Email: oauthResponse.Email, Username: oauthResponse.Name }

}
