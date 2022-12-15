package handler

import (
	"fmt"
	"linkservice/dto"
	"linkservice/service"
	"net/http"
	"linkservice/model"
	"github.com/gin-gonic/gin"
)

type Server struct {
	linkService *service.LinkService
}

func NewServer(linkService *service.LinkService) *Server {
	return &Server{linkService}
}

func (server *Server) CreateLink(c *gin.Context) {
	var requestBody dto.CreateLinkDto
   // bind the headers to data
   header := &model.Header{}

   if err := c.ShouldBindHeader(header); err != nil {
		c.JSON(400, err.Error())
		return
	}


	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Println(err)
	}

	server.linkService.CreateLink(&requestBody,header.UserId)
	c.JSON(http.StatusOK, "test")

}
func (server *Server) GetLink(c *gin.Context) {

	c.JSON(http.StatusOK, "test")

}
