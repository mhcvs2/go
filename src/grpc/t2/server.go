package main

// server.go

import (
	"log"
	"net"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "grpc/t2/helloworld"
	"fmt"
)

const (
	port = ":50051"
)

type server struct {}

func (s *server) Test1(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println(in.Name)
	return &pb.HelloReply{Message: "test1 " + in.Name}, nil
}


func main() {
	fmt.Println("start")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMhcTestServer(s, &server{})
	s.Serve(lis)
}