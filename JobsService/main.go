package main

import (
	"context"
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
	client := proto.NewAddServiceClient(conn)
	req := &proto.Request{A: int64(2), B: int64(3)}

	if response, err := client.Add(context.Background(), req); err == nil {
		fmt.Println(response.Result)
	}
	db := database.SetDB()
	jobRepository := repository.NewJobRepository(db)

	jobService := service.NewJobService(jobRepository)
	jobService.Test()

	jobService.RetrieveJobs()
	fmt.Println("test")
	runtime.Goexit()
	//TODO create a channel, and use it to end the main go loop
	fmt.Println("Exit")

}

// func (s *server) Add(ctx context.Context, request *proto.Request)(*proto.Response, error){
// 	a
// }
