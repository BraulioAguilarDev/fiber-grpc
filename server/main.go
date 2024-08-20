package main

import (
	"context"
	"fmt"
	pb "grpc/proto/greeter"
	"log"
	"net"
	"time"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

const (
	HOST          = "localhost"
	TCP_ADDRESS   = ":50051"
	FIBER_ADDRESS = ":3000"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: fmt.Sprintf("Hello %s!", req.Name),
	}, nil
}

func startGRPCServer() {
	lis, err := net.Listen("tcp", TCP_ADDRESS)
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	srv := grpc.NewServer()
	pb.RegisterGreeterServer(srv, &server{})
	log.Printf("server listening at %v", lis.Addr())

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	fmt.Println("Server started")

	go startGRPCServer()

	app := fiber.New()
	app.Get("/hello/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")

		var opts []grpc.DialOption
		opts = append(opts, grpc.WithInsecure())
		conn, err := grpc.NewClient(fmt.Sprintf("%s%s", HOST, TCP_ADDRESS), opts...)
		if err != nil {
			return err
		}
		defer conn.Close()

		client := pb.NewGreeterClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		r, err := client.SayHello(ctx, &pb.HelloRequest{Name: name})
		if err != nil {
			return err
		}

		return c.SendString(r.GetMessage())
	})

	if err := app.Listen(FIBER_ADDRESS); err != nil {
		log.Fatalf("could not start fiber server")
	}
}
