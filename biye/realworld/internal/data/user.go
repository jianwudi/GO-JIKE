package data

import (
	"context"
	"realworld/internal/biz"

	userv1 "realworld/api/account/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) error {
	//发grpc消息
	_, err := r.data.uc.Register(ctx, &userv1.RegisterRequest{
		User: &userv1.RegisterRequest_User{
			Email:    u.Email,
			Username: u.Username,
			Password: u.Password,
		},
	})

	if err != nil {
		return err
	}
	return nil
}

func (u *userRepo) GetUserByEmail(ctx context.Context, email string) (*biz.User, error) {
	return nil, nil
}
