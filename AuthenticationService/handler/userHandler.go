package handler

import (
	"authenticationservice/config"
	"authenticationservice/dto"
	"authenticationservice/model"
	"authenticationservice/service"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	userService *service.UserService
	config      *config.Config
}

func NewServer(userService *service.UserService, config *config.Config) *Server {
	return &Server{userService, config}
}

func (server *Server) RegisterUser(c *gin.Context) {
	url := server.config.GoogleConfig.AuthCodeURL("pseudo-random")
	
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (server *Server) Callback(c *gin.Context) {
	token, err := server.config.GoogleConfig.Exchange(context.Background(), c.Request.FormValue("code"))

	if err != nil {
		fmt.Print("could not get token")
		return
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)

	if err != nil {
		fmt.Print("could not get request")

	}
	defer resp.Body.Close()
	// content, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Print("could not parse resp")
	// }
	oauthResponse := model.OauthResponse{}

	if err := json.NewDecoder(resp.Body).Decode(&oauthResponse); err != nil {
		fmt.Println("Error %w", err)
	}

	userToCreate := dto.ToCreateUserDto(oauthResponse)

	jwt_token := server.userService.CreateUser(userToCreate)
	c.JSON(http.StatusOK, jwt_token)

}
