package service

import (
	v1 "account-service/api/account/v1"
	"fmt"
	"io"

	"account-service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	opentracing "github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
	config "github.com/uber/jaeger-client-go/config"
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

func initJaeger(service string) (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
		ServiceName: service,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:          true,
			CollectorEndpoint: "http://192.168.72.128:14268/api/traces",
		},
	}
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	return tracer, closer
}
