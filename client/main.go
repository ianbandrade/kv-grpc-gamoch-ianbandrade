package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"key_value/client/pb"
	"strings"
	"time"
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

	_, _ = c.Put(ctx, &pb.KeyValue{Key: "foo", Value: "bar"})

	getResponse, _ := c.Get(ctx, &pb.Key{Key: "foo"})
	fmt.Printf("Get Response: %s\n", getResponse.GetValue())

	getAllKeysResponse, _ := c.GetAllKeys(ctx, &pb.Empty{})
	fmt.Printf("GetAllKeys Response: %s\n", strings.Join(getAllKeysResponse.GetKeys(), ", "))
}
