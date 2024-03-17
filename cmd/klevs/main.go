package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/spf13/pflag"
	"go.klev.dev/klevs/protocol/grpc"
	"go.klev.dev/klevs/server"
	"golang.org/x/sync/errgroup"
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

	logServer, err := server.New(*dataDir)
	if err != nil {
		panic(err)
	}

	gServer, err := grpc.New(logServer)
	if err != nil {
		panic(err)
	}

	var g, ctx = errgroup.WithContext(context.Background())
	g.Go(func() error {
		return gServer.Run(ctx)
	})

	if err := g.Wait(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
