package main

import (
	"fmt"
	"linkservice/database"
	"net"
	"os"
	"time"

	"linkservice/handler"
	"linkservice/repository"
	"linkservice/service"

	"context"
	"linkservice/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/confluentinc/confluent-kafka-go/kafka"
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

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"client.id":         "linkservice",
		"acks":              "all"})
	topic := "test"
	time.Sleep(1000)
	delivery_chan := make(chan kafka.Event, 10000)

	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: int32(kafka.PartitionAny)},
		Value:          []byte("val")},
		delivery_chan,
	)
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: int32(kafka.PartitionAny)},
		Value:          []byte("val")},
		delivery_chan,
	)
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: int32(kafka.PartitionAny)},
		Value:          []byte("val")},
		delivery_chan,
	)
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: int32(kafka.PartitionAny)},
		Value:          []byte("val")},
		delivery_chan,
	)
	fmt.Println("REACH?")

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	go func() {
		consumer, _ := kafka.NewConsumer(&kafka.ConfigMap{
			"bootstrap.servers": "localhost:9092",
			"group.id":          "foo",
			"auto.offset.reset": "smallest"})
		err = consumer.Subscribe(topic, nil)

		for {
			ev := consumer.Poll(100)
			switch e := ev.(type) {
			case *kafka.Message:
				fmt.Println("RECEIVED %s\n", string(e.Value))
				// application-specific processing
			case kafka.Error:
				fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)

			}
		}
		// consumer.Close()

	}()
	// go func() {
	// 	for num := 1; num < 2000; num++ {

	// 		err = p.Produce(&kafka.Message{
	// 			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: int32(kafka.PartitionAny)},
	// 			Value:          []byte(strconv.Itoa(num))},
	// 			delivery_chan,
	// 		)
	// 		fmt.Println("run: %s", strconv.Itoa(num))

	// 	}
	// }()
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
