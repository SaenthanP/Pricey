package handler

import (
	"fmt"
	"linkservice/dto"
	"linkservice/model"
	"linkservice/service"
	"net/http"

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

	link, err := server.linkService.CreateLink(&requestBody, header.UserId)

	//Must break it out if the resource already exists later
	if len(err) == 0 {
		 c.JSON(http.StatusCreated, link)
		 return
	}

	 c.JSON(http.StatusInternalServerError, err)
	 return
}
func (server *Server) GetLink(c *gin.Context) {

	c.JSON(http.StatusOK, "test")

}
