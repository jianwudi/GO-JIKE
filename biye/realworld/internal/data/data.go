package data

import (
	"fmt"
	"realworld/internal/conf"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	userv1 "realworld/api/account/v1"

	"google.golang.org/grpc"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserServiceClient, NewUserRepo)

// Data .
type Data struct {
	uc userv1.AccountClient
}

// NewData .
func NewData(c *conf.Data, uc userv1.AccountClient, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{uc: uc}, cleanup, nil
}

func NewUserServiceClient(conf *conf.Data) userv1.AccountClient {
	fmt.Printf("addr:%v\n", conf.Grpccli.Addr)
	conn, err := grpc.Dial(conf.Grpccli.Addr, grpc.WithInsecure(), grpc.WithTimeout(time.Second), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	c := userv1.NewAccountClient(conn)
	return c
}
