package main

import (
	"context"
	pb "grpc/proto/greeter"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	var opts []grpc.DialOption
	// TODO: change WithInsecure opt
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.NewClient("localhost:50051", opts...)
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Braulio"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", r.GetMessage())
}
