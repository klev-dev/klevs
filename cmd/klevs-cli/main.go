package main

import (
	"context"
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

	client := api.NewLogsClient(conn)

	rootCmd := &cobra.Command{
		Use:   "klevs-cli",
		Short: "klevs cli client",
	}

	rootCmd.AddCommand(logsCmd(client))

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func logsCmd(logs api.LogsClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logs",
		Short: "interact with logs",
	}

	cmd.AddCommand(logsListCmd(logs))
	cmd.AddCommand(logsCreateCmd(logs))

	return cmd
}

func logsListCmd(logs api.LogsClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list logs",
	}

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		resp, err := logs.List(context.TODO(), &api.LogsListRequest{})
		if err != nil {
			return err
		}

		fmt.Println("resp:", resp)
		return nil
	}

	return cmd
}

func logsCreateCmd(logs api.LogsClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create <name>",
		Short: "create new log",
		Args:  cobra.ExactArgs(1),
	}

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		resp, err := logs.Create(context.TODO(), &api.LogsCreateRequest{
			Name: args[0],
		})
		if err != nil {
			return err
		}

		fmt.Println("resp:", resp)
		return nil
	}

	return cmd
}
