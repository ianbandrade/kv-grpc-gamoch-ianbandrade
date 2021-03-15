package main

import (
	"context"
	"fmt"
	"key_value/server/pb"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedKeyValueServiceServer
}

var storage = make(map[string]string)

func (s *server) Get(_ context.Context, in *pb.Key) (*pb.Value, error) {
	return &pb.Value{Value: storage[in.Key], Defined: storage[in.Key] != ""}, nil
}

func (s *server) Put(_ context.Context, in *pb.KeyValue) (*pb.Empty, error) {
	storage[in.Key] = in.Value

	return &pb.Empty{}, nil
}

func (s *server) GetAllKeys(_ context.Context, _ *pb.Empty) (*pb.StoredKeys, error) {
	var keys []string

	for k := range storage {
		keys = append(keys, k)
	}

	return &pb.StoredKeys{Keys: keys}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		fmt.Printf("Failed to listen: %v\n", err)
	}

	s := grpc.NewServer()

	fmt.Printf("Server listening on localhost%s\n", port)

	pb.RegisterKeyValueServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v\n", err)
	}
}
