// +build wireinject

package app

import (
	"week04/code/internal/apiserver/biz"
	"week04/code/internal/apiserver/data"
	"week04/code/internal/apiserver/server"
	"week04/code/internal/apiserver/service"

	"github.com/google/wire"
)

/* func InitApp() (*App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, NewApp))
} */

func InitApp() (*App, func(), error) {
	panic(wire.Build(data.ProviderSet, biz.ProviderSet, service.ProviderSet, server.ProviderSet, NewApp))
}
