package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"week04/code/internal/apiserver/app"

	"golang.org/x/sync/errgroup"
)

func main() {
	a, _, _ := app.InitApp()
	eg, ctx := errgroup.WithContext(context.Background())
	eg.Go(func() error {
		<-ctx.Done()
		a.Stop()
		fmt.Println("STOP")
		return nil
	})
	eg.Go(func() error {
		fmt.Println("run")
		return a.Run()
	})
	eg.Go(func() error {
		return hookSignal(ctx)
	})
	if err := eg.Wait(); err != nil {
		fmt.Println("err:", err.Error())
	}
	fmt.Println("DONE!")
}

func hookSignal(ctx context.Context) error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case s := <-c:
			return fmt.Errorf("get %v signal", s)
		case <-ctx.Done():
			return fmt.Errorf("other goroutine quit")
		}
	}
}
