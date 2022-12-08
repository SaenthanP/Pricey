package main

import (
	"scraperservice/database"

	"scraperservice/handler"
	"scraperservice/repository"
	"scraperservice/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.SetDB()

	linkRepository := repository.NewLinkRepository(db)
	linkService := service.NewLinkSevice(linkRepository)
	linkHandler := handler.NewServer(linkService)

	router := gin.Default()

	router.POST("/api/create_url", linkHandler.CreateLink)

	router.Run(":8081")
}
