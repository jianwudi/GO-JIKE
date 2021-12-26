package app

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

type App struct {
	grpcSvr *grpc.Server
}

func NewApp(grpcSvr *grpc.Server) (app *App, err error) {
	app = &App{
		grpcSvr: grpcSvr,
	}
	return app, nil
}

func (app *App) Run() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	return app.grpcSvr.Serve(lis)
}

func (app *App) Stop(){
	app.grpcSvr.GracefulStop()
}
