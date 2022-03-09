package service

import (
	v1 "account-service/api/realworld/v1"
	"account-service/internal/pkg/errors"
	"context"
	"fmt"
)

var (
	ErrMissingEmail = errors.NewHTTPError(422, "email", "cannot be empty")
	ErrRegisterFail = errors.NewHTTPError(423, "register", "fail")
)

func (s *AccountService) Login(ctx context.Context, req *v1.LoginRequest) (reply *v1.UserReply, err error) {
	if len(req.User.Email) == 0 {
		return nil, ErrMissingEmail
	}

	fmt.Println("Login")
	u, err := s.uc.Login(ctx, req.User.Email, req.User.Password)
	if err != nil {
		return nil, ErrRegisterFail
	}
	return &v1.UserReply{
		User: &v1.UserReply_User{
			Username: u.UserName,
			Token:    u.Token,
			Email:    u.Email,
		},
	}, nil
}

func (s *AccountService) Register(ctx context.Context, req *v1.RegisterRequest) (reply *v1.UserReply, err error) {
	u, err := s.uc.Register(ctx, req.User.Username, req.User.Email, req.User.Password)
	if err != nil {
		return nil, err
	}
	return &v1.UserReply{
		User: &v1.UserReply_User{
			Username: u.UserName,
			Token:    u.Token,
			Email:    u.Email,
		},
	}, nil
}

func (s *AccountService) GetCurrentUser(ctx context.Context, req *v1.GetCurrentUserRequest) (reply *v1.UserReply, err error) {

	return &v1.UserReply{
		User: &v1.UserReply_User{
			Username: "boom",
		},
	}, nil
}

func (s *AccountService) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (reply *v1.UserReply, err error) {
	return &v1.UserReply{
		User: &v1.UserReply_User{
			Username: "boom",
		},
	}, nil
}
