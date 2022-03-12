package service

import (
	"context"

	v1 "realworld/api/realworld/v1"
)

func (s *RealWorldService) Login(ctx context.Context, req *v1.LoginRequest) (reply *v1.UserReply, err error) {
	return &v1.UserReply{
		User: &v1.UserReply_User{
			Username: "boom",
		},
	}, nil
}

func (s *RealWorldService) Register(ctx context.Context, req *v1.RegisterRequest) (reply *v1.UserReply, err error) {

	return &v1.UserReply{
		User: &v1.UserReply_User{
			Username: "boom",
		},
	}, nil
}
