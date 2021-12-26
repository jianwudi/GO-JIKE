package server

import (
	v1 "week04/code/api/helloworld/v1"
	"week04/code/internal/apiserver/service"

	"google.golang.org/grpc"
)

func NewGRPCServer(greeter *service.GreeterService) *grpc.Server {
	srv := grpc.NewServer()
	v1.RegisterGreeterServer(srv, greeter)
	return srv
}
