package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/spf13/pflag"
	"go.klev.dev/klevs/api"
	"go.klev.dev/klevs/server"

	"google.golang.org/grpc"
)

func main() {
	flags := pflag.NewFlagSet(os.Args[0], pflag.ContinueOnError)
	dataDir := flags.String("data-dir", "/var/lib/klevs", "root data directory")
	if err := flags.Parse(os.Args[1:]); err != nil {
		if errors.Is(err, pflag.ErrHelp) {
			os.Exit(0)
		}
		panic(err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 9283))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	// ...
	grpcServer := grpc.NewServer(opts...)

	api.RegisterLogsServer(grpcServer, &server.Logs{
		DataDir: *dataDir,
	})
	api.RegisterMessagesServer(grpcServer, &server.Messages{
		DataDir: *dataDir,
	})

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
