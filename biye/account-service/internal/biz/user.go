package biz

import (
	"account-service/internal/conf"
	"account-service/internal/pkg/midware/auth"
	"context"
	"errors"

	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email        string
	Username     string
	Bio          string
	Image        string
	PasswordHash string
}

type UserLogin struct {
	Email    string
	UserName string
	Token    string
	Bio      string
	Image    string
}

func hashPassword(pwd string) string {
	b, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func verifyPasswd(hashed, input string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(input)); err != nil {
		return false
	}
	return true
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}

type ProfileRepo interface {
}

type UserUsecase struct {
	ur   UserRepo
	log  *log.Helper
	jwtc *conf.JWT
}

func NewUserUsecase(ur UserRepo, jwtc *conf.JWT, logger log.Logger) *UserUsecase {
	return &UserUsecase{ur: ur, jwtc: jwtc, log: log.NewHelper(logger)}
}
func (uc *UserUsecase) generalToken(username string) string {
	return auth.GenerateToken(uc.jwtc.Secret, username)
}
func (uc *UserUsecase) Register(ctx context.Context, username, email, password string) (*UserLogin, error) {
	u := &User{
		Email:        email,
		Username:     username,
		PasswordHash: hashPassword(password),
	}
	if err := uc.ur.CreateUser(ctx, u); err != nil {
		return nil, err
	}
	return &UserLogin{
		Email:    email,
		UserName: username,
		Token:    uc.generalToken(username),
	}, nil
}

func (uc *UserUsecase) Login(ctx context.Context, email, password string) (*UserLogin, error) {
	u, err := uc.ur.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if !verifyPasswd(u.PasswordHash, password) {
		return nil, errors.New("login failed")
	}
	return &UserLogin{
		Email:    u.Email,
		UserName: u.Username,
		Bio:      u.Bio,
		Image:    u.Image,
		Token:    "xxx",
	}, nil
}
