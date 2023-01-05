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

	syncmessaging.SetupGrpc()

	asyncMessagingClient := asyncmessaging.NewAsyncMessageClient()

	// go func() {
	// 	consumer, _ := kafka.NewConsumer(&kafka.ConfigMap{
	// 		"bootstrap.servers": "localhost:9092",
	// 		"group.id":          "foo",
	// 		"auto.offset.reset": "smallest"})
	// 	err = consumer.Subscribe(topic, nil)

	// 	for {
	// 		ev := consumer.Poll(100)
	// 		switch e := ev.(type) {
	// 		case *kafka.Message:
	// 			fmt.Println("RECEIVED %s\n", string(e.Value))
	// 			// application-specific processing
	// 		case kafka.Error:
	// 			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)

	// 		}
	// 	}
	// 	// consumer.Close()

	// }()

	db := database.SetDB()

	linkRepository := repository.NewLinkRepository(db)
	approvedLinkRepository := repository.NewApprovedLinkRepository(db)
	userToLinkRepository := repository.NewUserToLinkRepository(db)
	linkService := service.NewLinkSevice(asyncMessagingClient, linkRepository, approvedLinkRepository, userToLinkRepository)
	
	syncmessaging.NewAsyncMessageClient(linkService)

	linkHandler := handler.NewServer(linkService)

	router := gin.Default()

	router.POST("/api/link/createlink", linkHandler.CreateLink)
	router.GET("/api/link/geturl", linkHandler.GetLink)

	router.Run(":8081")
}
