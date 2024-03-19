package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"go.klev.dev/klevs/protocol/grpc"
	"google.golang.org/protobuf/encoding/prototext"
)

func logsCmd(logs grpc.LogsClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logs",
		Short: "interact with logs",
	}

	cmd.AddCommand(logsListCmd(logs))
	cmd.AddCommand(logsCreateCmd(logs))
	cmd.AddCommand(logsDeleteCmd(logs))

	return cmd
}

func logsListCmd(logs grpc.LogsClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list logs",
	}

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		resp, err := logs.List(cmd.Context(), &grpc.LogsListRequest{})
		if err != nil {
			return err
		}

		fmt.Println(prototext.Format(resp))
		return nil
	}

	return cmd
}

func logsCreateCmd(logs grpc.LogsClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create <name>",
		Short: "create new log",
		Args:  cobra.ExactArgs(1),
	}

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		resp, err := logs.Create(cmd.Context(), &grpc.LogsCreateRequest{
			Name: args[0],
		})
		if err != nil {
			return err
		}

		fmt.Println(prototext.Format(resp))
		return nil
	}

	return cmd
}

func logsDeleteCmd(logs grpc.LogsClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete <name>",
		Short: "delete a log",
		Args:  cobra.ExactArgs(1),
	}

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		resp, err := logs.Delete(cmd.Context(), &grpc.LogsDeleteRequest{
			Name: args[0],
		})
		if err != nil {
			return err
		}

		fmt.Println(prototext.Format(resp))
		return nil
	}

	return cmd
}
