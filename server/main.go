package main

import (
	"context"
	pb "grpc-test/server/pb/greeting"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedGreetorServer
}

func (s *server) SayHello(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	log.Printf("Received request to greet '%s'\n", in.GetName())
	return &pb.Response{
		Greeting: "Hello, " + in.GetName(),
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":5051")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterGreetorServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
