package service

import (
	"context"

	v1 "week04/code/api/helloworld/v1"
	"week04/code/internal/apiserver/biz"
)

type GreeterService struct {
	v1.UnimplementedGreeterServer
	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

// SayHello implements helloworld.GreeterServer
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	/* 	if in.GetName() == "error" {
		return nil, v1.ErrorUserNotFound("user not found: %s", in.GetName())
	} */
	return &v1.HelloReply{Message: "Hello " + in.GetName()}, nil
}
