package main

import (
	"context"
	"fmt"

	"go.klev.dev/klevs/api"
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

	resp, err := client.List(context.TODO(), &api.LogsListRequest{})
	if err != nil {
		panic(err)
	}
	fmt.Println("resp:", resp)
}
