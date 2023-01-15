package syncmessaging

import (
	"context"
	"fmt"
	"linkservice/proto"
	"linkservice/service"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.CallScrapeServiceServer
}

var _linkService *service.LinkService

func NewSyncMessageClient(linkService *service.LinkService) {
	_linkService = linkService
	setupGrpc()

}

func setupGrpc() {
	listener, err := net.Listen("tcp", "localhost:9001")
	if err != nil {
		fmt.Printf("GRPC connection failed: %v", err)
	}
	grpcServer := grpc.NewServer()

	proto.RegisterCallScrapeServiceServer(grpcServer, &server{})
	//Needed to seriealize and deserialize data
	reflection.Register(grpcServer)

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			fmt.Printf("Grpc server failed %v", err)
		}
	}()
}

func (s *server) CallScrape(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	_linkService.ScrapeJob()
	
	return &proto.Response{Result: 1}, nil
}
