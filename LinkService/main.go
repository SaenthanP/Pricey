package main

import (
	"linkservice/database"

	"linkservice/handler"
	"linkservice/repository"
	"linkservice/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.SetDB()

	linkRepository := repository.NewLinkRepository(db)
	approvedLinkRepository := repository.NewApprovedLinkRepository(db)
	userToLinkRepository := repository.NewUserToLinkRepository(db)
	linkService := service.NewLinkSevice(linkRepository, approvedLinkRepository, userToLinkRepository)

	linkHandler := handler.NewServer(linkService)

	router := gin.Default()

	router.POST("/api/createlink", linkHandler.CreateLink)
	router.GET("/api/geturl", linkHandler.GetLink)

	router.Run(":8081")
}
