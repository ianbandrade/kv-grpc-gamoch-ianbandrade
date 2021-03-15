package main

import (
	"context"
	"fmt"
	"key_value/client/pb"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		fmt.Printf("Did not connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewKeyValueServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var cmdPut = &cobra.Command{
		Use:   "put [key] [value]",
		Short: "Put a KeyValue object into a server",
		Long:  `Put a data with a key and value into the server with one map struct object`,
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			c.Put(ctx, &pb.KeyValue{Key: args[0], Value: args[1]})
		},
	}

	var cmdGet = &cobra.Command{
		Use:   "get [key]",
		Short: "Get an object value as from your key",
		Long:  `Get a value of an object into the map struct of the server just passing a key`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			getResponse, _ := c.Get(ctx, &pb.Key{Key: args[0]})
			fmt.Printf("Get: Key='%s' Value='%s'\n", args[0], getResponse.GetValue())
		},
	}

	var cmdGetAllKeys = &cobra.Command{
		Use:   "getAllKeys",
		Short: "Get all Keys",
		Long:  `Get all Keys of the objects into the map struct of the server`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			getAllKeysResponse, _ := c.GetAllKeys(ctx, &pb.Empty{})
			fmt.Printf("GetAllKeys: %s\n", strings.Join(getAllKeysResponse.GetKeys(), ", "))
		},
	}

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmdPut, cmdGet, cmdGetAllKeys)
	rootCmd.Execute()
}
