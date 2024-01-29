package services

import (
	"context"

	pb "github.com/gameboy86/grpc-server-fx/example/hello"
	"google.golang.org/grpc"
)

func (s *HelloService) SayHello(
	context.Context,
	*pb.HelloRequest,
) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello"}, nil
}

func (s *HelloService) Description() *grpc.ServiceDesc {
	return &pb.Greeter_ServiceDesc
}

func NewHelloService() *HelloService {
	return &HelloService{}
}

type HelloService struct {
	pb.UnsafeGreeterServer
}
