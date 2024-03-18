package main

import (
	"fmt"

	"github.com/klev-dev/klevdb"
	"github.com/spf13/cobra"
	"go.klev.dev/klevs/protocol/grpc"
)

func messagesCmd(messages grpc.MessagesClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "messages",
		Short: "interact with messages",
	}

	cmd.AddCommand(messagesPublishCmd(messages))
	cmd.AddCommand(messagesNextOffsetCmd(messages))
	cmd.AddCommand(messagesConsumeCmd(messages))
	cmd.AddCommand(messagesGetByOffsetCmd(messages))

	return cmd
}

func messagesPublishCmd(messages grpc.MessagesClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "publish <name>",
		Short: "publish new message",
		Args:  cobra.ExactArgs(1),
	}

	keyStr := cmd.Flags().String("key-string", "", "key to publish, as string")
	keyBase64 := cmd.Flags().BytesBase64("key-base64", nil, "key to publish, as base64")
	valueStr := cmd.Flags().String("value-string", "", "value to publish, as string")
	valueBase64 := cmd.Flags().BytesBase64("value-base64", nil, "value to publish, as base64")

	cmd.MarkFlagsMutuallyExclusive("key-string", "key-base64")
	cmd.MarkFlagsMutuallyExclusive("value-string", "value-base64")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		msg := &grpc.PublishMessage{}

		switch {
		case cmd.Flags().Changed("key-string"):
			msg.Key = []byte(*keyStr)
		case cmd.Flags().Changed("key-base64"):
			msg.Key = *keyBase64
		}

		switch {
		case cmd.Flags().Changed("value-string"):
			msg.Value = []byte(*valueStr)
		case cmd.Flags().Changed("value-base64"):
			msg.Value = *valueBase64
		}

		resp, err := messages.Publish(cmd.Context(), &grpc.PublishRequest{
			Name:     args[0],
			Messages: []*grpc.PublishMessage{msg},
		})
		if err != nil {
			return err
		}

		fmt.Println("resp:", resp)
		return nil
	}

	return cmd
}

func messagesNextOffsetCmd(messages grpc.MessagesClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "next-offset <name>",
		Short: "get next offset of log",
		Args:  cobra.ExactArgs(1),
	}

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		resp, err := messages.NextOffset(cmd.Context(), &grpc.NextOffsetRequest{
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

func messagesConsumeCmd(messages grpc.MessagesClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "consume <name>",
		Short: "consume messages",
		Args:  cobra.ExactArgs(1),
	}

	offset := cmd.Flags().Int64("offset", klevdb.OffsetOldest, "offset to start (default oldest)")
	maxCount := cmd.Flags().Int64("max-count", 32, "max number of message to return (defaults to 32)")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		resp, err := messages.Consume(cmd.Context(), &grpc.ConsumeRequest{
			Name:     args[0],
			Offset:   *offset,
			MaxCount: *maxCount,
		})
		if err != nil {
			return err
		}

		fmt.Println("resp:", resp)
		return nil
	}

	return cmd
}

func messagesGetByOffsetCmd(messages grpc.MessagesClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-by-offset <name>",
		Short: "get message by offset",
		Args:  cobra.ExactArgs(1),
	}

	offset := cmd.Flags().Int64("offset", klevdb.OffsetOldest, "offset to get (default oldest)")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		resp, err := messages.GetByOffset(cmd.Context(), &grpc.GetByOffsetRequest{
			Name:   args[0],
			Offset: *offset,
		})
		if err != nil {
			return err
		}

		fmt.Println("resp:", resp)
		return nil
	}

	return cmd
}
