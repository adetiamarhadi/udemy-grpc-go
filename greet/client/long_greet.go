package main

import (
	"context"
	"log"
	"time"

	pb "github.com/adetiamarhadi/udemy-grpc-go/greet/proto"
)

func callLongGreet(c pb.GreetServiceClient) {

	log.Println("Call Long Greet Function")

	reqs := []*pb.GreetRequest{
		{FirstName: "Adet"},
		{FirstName: "Marhadi"},
		{FirstName: "Joe"},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling Long Greet: %v\n", err)
	}

	for _, req := range reqs {

		log.Printf("Sending request: %v\n", req)

		stream.Send(req)

		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from Long Greet: %v\n", err)
	}

	log.Printf("Result: %v\n", res.Result)
}
