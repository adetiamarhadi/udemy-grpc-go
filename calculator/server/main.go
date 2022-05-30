package main

import (
	"log"
	"net"

	pb "github.com/adetiamarhadi/udemy-grpc-go/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var addr string = "0.0.0.0:50059"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Fail listen: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()

	pb.RegisterCalculatorServiceServer(s, &Server{})

	reflection.Register(s)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Fail Serve: %v\n", err)
	}
}
