package service

import (
	v1 "account-service/api/realworld/v1"

	"account-service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewAccountService)

type AccountService struct {
	v1.UnimplementedAccountServer

	uc  *biz.UserUsecase
	log *log.Helper
}

func NewAccountService(uc *biz.UserUsecase, logger log.Logger) *AccountService {
	return &AccountService{uc: uc, log: log.NewHelper(logger)}
}
