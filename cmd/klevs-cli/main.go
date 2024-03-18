package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	api "go.klev.dev/klevs/protocol/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// ...
	conn, err := grpc.Dial("localhost:9283", opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	logs := api.NewLogsClient(conn)
	messages := api.NewMessagesClient(conn)

	if err := rootCmd(logs, messages).Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func rootCmd(logs api.LogsClient, messages api.MessagesClient) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "klevs-cli",
		Short: "klevs cli client",
	}

	rootCmd.AddCommand(logsCmd(logs))
	rootCmd.AddCommand(messagesCmd(messages))

	return rootCmd
}
