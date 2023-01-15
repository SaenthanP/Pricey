package main

import (
	"linkservice/asyncmessaging"
	"linkservice/database"
	"linkservice/syncmessaging"

	"linkservice/handler"
	"linkservice/repository"
	"linkservice/service"

	"github.com/gin-gonic/gin"
)

func main() {
	asyncMessagingClient := asyncmessaging.NewAsyncMessageClient()

	db := database.SetDB()

	linkRepository := repository.NewLinkRepository(db)
	approvedLinkRepository := repository.NewApprovedLinkRepository(db)
	userToLinkRepository := repository.NewUserToLinkRepository(db)
	linkService := service.NewLinkSevice(asyncMessagingClient, linkRepository, approvedLinkRepository, userToLinkRepository)
	
	syncmessaging.NewSyncMessageClient(linkService)

	linkHandler := handler.NewServer(linkService)

	router := gin.Default()

	router.POST("/api/link/createlink", linkHandler.CreateLink)
	router.GET("/api/link/geturl", linkHandler.GetLink)

	router.Run(":8081")
}
