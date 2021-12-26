package data

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

type Data struct {
}

func NewData() (*Data, func(), error) {
	cleanup := func() {
	}
	return &Data{}, cleanup, nil
}
