package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/adetiamarhadi/udemy-grpc-go/calculator/proto"
)

var addr string = "localhost:50059"

func main() {

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("can't connect to %s: %v/n", addr, err)
	}

	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)

	callSum(c)
}
