package handler

import (
	"scraperservice/service"
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
	c.JSON(http.StatusOK, "test")

}
