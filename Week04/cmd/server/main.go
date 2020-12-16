package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"

	"Week04/api"
	"Week04/internal/pkg/server"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	comServ := InitializeCommentService()
	s := server.NewServer(":9090")
	api.RegisterCommentServer(s.Server, comServ)

	// 监听信号
	g.Go(func() error {
		sigChan := make(chan os.Signal)

		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
		for {
			select {
			case sig := <-sigChan:
				return fmt.Errorf("receive signal: %v\n", sig)
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	})

	// 启动server
	g.Go(func() error {
		log.Println("Server started...")
		return s.Run(ctx)
	})

	if err := g.Wait(); err != nil {
		log.Fatalf("process exited %+v\n", err)
	}
}
