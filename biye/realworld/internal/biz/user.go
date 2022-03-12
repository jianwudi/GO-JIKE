package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Email    string
	Username string
	Password string
}

type UserLogin struct {
	Email    string
	UserName string
	Token    string
	Bio      string
	Image    string
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}
type UserUsecase struct {
	ur  UserRepo
	log *log.Helper
}

func NewUserUsecase(ur UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{ur: ur, log: log.NewHelper(logger)}
}
func (uc *UserUsecase) Register(ctx context.Context, username, email, password string) (*UserLogin, error) {
	u := &User{
		Email:    email,
		Username: username,
		Password: password,
	}
	if err := uc.ur.CreateUser(ctx, u); err != nil {
		return nil, err
	}
	return &UserLogin{
		Email:    email,
		UserName: username,
	}, nil
}

func (uc *UserUsecase) Login(ctx context.Context, email, password string) (*UserLogin, error) {
	u, err := uc.ur.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return &UserLogin{
		Email:    u.Email,
		UserName: u.Username,
		Token:    "xxx",
	}, nil
}
