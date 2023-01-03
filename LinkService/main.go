package main

import (
	"fmt"
	"linkservice/database"
	"net"

	"linkservice/handler"
	"linkservice/repository"
	"linkservice/service"

	"context"
	"linkservice/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.AddServiceServer
}

func main() {
	listener, err := net.Listen("tcp", "localhost:9001")
	if err != nil {
		fmt.Printf("GRPC connection failed: %v", err)
	}
	grpcServer := grpc.NewServer()

	proto.RegisterAddServiceServer(grpcServer, &server{})
	//Needed to seriealize and deserialize data
	reflection.Register(grpcServer)
	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			fmt.Printf("Grpc server failed %v", err)
		}
	}()

	db := database.SetDB()

	linkRepository := repository.NewLinkRepository(db)
	approvedLinkRepository := repository.NewApprovedLinkRepository(db)
	userToLinkRepository := repository.NewUserToLinkRepository(db)
	linkService := service.NewLinkSevice(linkRepository, approvedLinkRepository, userToLinkRepository)

	linkHandler := handler.NewServer(linkService)

	router := gin.Default()

	router.POST("/api/link/createlink", linkHandler.CreateLink)
	router.GET("/api/link/geturl", linkHandler.GetLink)

	router.Run(":8081")
}
func (s *server) Add(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a + b
	return &proto.Response{Result: result}, nil
}
