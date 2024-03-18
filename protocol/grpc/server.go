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

func New(logs *klevs.Logs) (protocol.Server, error) {
	var opts []grpc.ServerOption
	// ...
	server := grpc.NewServer(opts...)

	RegisterLogsServer(server, &Logs{
		Logs: logs,
	})
	RegisterMessagesServer(server, &Messages{
		Logs: logs,
	})

	return &grpcServer{server}, nil
}

type grpcServer struct {
	server *grpc.Server
}

func (s *grpcServer) Run(ctx context.Context) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 9283))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go func() {
		<-ctx.Done()
		s.server.GracefulStop()
	}()

	return s.server.Serve(lis)
}
