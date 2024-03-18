package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"go.klev.dev/klevs"
	"go.klev.dev/klevs/protocol"
	"google.golang.org/grpc"
)

func New(logServer *klevs.Logs) (protocol.Server, error) {
	var opts []grpc.ServerOption
	// ...
	gsrv := grpc.NewServer(opts...)

	RegisterLogsServer(gsrv, &Logs{
		Server: logServer,
	})
	RegisterMessagesServer(gsrv, &Messages{
		Logs: logServer,
	})

	return &grpcServer{gsrv}, nil
}

type grpcServer struct {
	gsrv *grpc.Server
}

func (s *grpcServer) Run(ctx context.Context) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 9283))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go func() {
		<-ctx.Done()
		s.gsrv.GracefulStop()
	}()

	return s.gsrv.Serve(lis)
}
