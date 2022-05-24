package main

import (
	"context"
	"io"
	"log"

	pb "github.com/adetiamarhadi/udemy-grpc-go/greet/proto"
)

func callGreetManyTimes(c pb.GreetServiceClient) {

	log.Println("call greet many times service")

	req := &pb.GreetRequest{
		FirstName: "Adet",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes: %v\n", err)
	}

	for {

		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("Greet Many Times: %s\n", msg.Result)
	}
}
