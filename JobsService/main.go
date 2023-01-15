package main

import (
	"fmt"
	"jobservice/database"
	"jobservice/proto"
	"jobservice/repository"
	"jobservice/service"
	"runtime"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	conn, err := grpc.Dial("localhost:9001", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	protoClient := proto.NewCallScrapeServiceClient(conn)

	db := database.SetDB()
	jobRepository := repository.NewJobRepository(db)

	jobService := service.NewJobService(protoClient, jobRepository)

	jobService.RetrieveJobs()
	runtime.Goexit()
	//TODO create a channel, and use it to end the main go loop
	fmt.Println("Exit")

}

// func (s *server) Add(ctx context.Context, request *proto.Request)(*proto.Response, error){
// 	a
// }
