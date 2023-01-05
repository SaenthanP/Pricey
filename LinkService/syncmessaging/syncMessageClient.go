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
	proto.AddServiceServer
}

var _linkService *service.LinkService

func NewSyncMessageClient(linkService *service.LinkService) {
	_linkService = linkService
}

func SetupGrpc() {
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
}

func (s *server) Add(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()
	result := a + b
	_linkService.TestCallFromRpc()
	return &proto.Response{Result: result}, nil
}
