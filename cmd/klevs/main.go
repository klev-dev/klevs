package main

import (
	"fmt"
	"log"
	"net"

	"go.klev.dev/klevs/api"
	"go.klev.dev/klevs/server"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 9283))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	// ...
	grpcServer := grpc.NewServer(opts...)

	api.RegisterLogsServer(grpcServer, &server.Logs{})

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
