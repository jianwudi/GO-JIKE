package data

import (
	"context"
	"week04/code/internal/apiserver/biz"
)

type greeterRepo struct {
	data *Data
}

func NewGreeterRepo(data *Data) biz.GreeterRepo {
	return &greeterRepo{data: data}
}
func (r *greeterRepo) CreateGreeter(ctx context.Context, g *biz.Greeter) error {
	return nil
}
func (r *greeterRepo) UpdateGreeter(ctx context.Context, g *biz.Greeter) error {
	return nil
}
