package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	*grpc.Server

	addr string
}

func NewServer(addr string) *Server {
	serv := grpc.NewServer()

	return &Server{
		Server: serv,
		addr:   addr,
	}
}

func (s *Server) Run(ctx context.Context) error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("recovered: %v\n", r)
			}
		}()
		<- ctx.Done()
		s.GracefulStop()
		fmt.Println("grpc server graceful stopped!")
	}()

	return s.Server.Serve(lis)
}