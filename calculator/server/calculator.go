package main

import (
	"context"
	"log"

	pb "github.com/adetiamarhadi/udemy-grpc-go/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {

	log.Printf("sum invoked with request: %v\n", in)

	return &pb.SumResponse{
		Result: in.Number_1 + in.Number_2,
	}, nil
}
