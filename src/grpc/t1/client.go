package t1

//client.go

import (
	"log"
	"os"
	"context"
	"google.golang.org/grpc"
	pb "grpc/t1/helloworld"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	defer conn.Close()
	//c := pb.NewGreeterClient(conn)
	c2 := pb.NewMhcTestClient(conn)

	name := defaultName
	if len(os.Args) >1 {
		name = os.Args[1]
	}
	//r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	//if err != nil {
	//	log.Fatal("could not greet: %v", err)
	//}
	//
	//log.Printf("Greeting: %s", r.Message)

	r, err := c2.Test1(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatal("could not test1: %v", err)
	}

	log.Printf("Greeting: %s", r.Message)
}