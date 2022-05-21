package main

import (
	"context"
	"log"

	pb "github.com/adetiamarhadi/udemy-grpc-go/greet/proto"
)

func callGreet(c pb.GreetServiceClient) {

	log.Println("calling greet function")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Adet",
	})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}
