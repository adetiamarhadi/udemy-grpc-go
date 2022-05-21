package main

import (
	"context"
	"log"

	pb "github.com/adetiamarhadi/udemy-grpc-go/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {

	log.Printf("Greet function was invoked with %v\n", in)

	return &pb.GreetResponse{
		Result: "Hi, " + in.FirstName,
	}, nil
}
