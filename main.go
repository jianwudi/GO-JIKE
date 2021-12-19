package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return server(ctx)
	})

	g.Go(func() error {
		return hookSignal(ctx)
	})
	if err := g.Wait(); err != nil {
		fmt.Println("err:", err.Error())
	}

	fmt.Println("DONE!")
}

func server(ctx context.Context) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "hello world!")
	})
	s := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func(ctx context.Context) {
		<-ctx.Done()
		fmt.Println("shutdown")
		s.Shutdown(context.Background())
	}(ctx)
	return s.ListenAndServe()
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
