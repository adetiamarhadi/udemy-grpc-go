package main

import (
	"context"
	"io"
	"log"

	pb "github.com/adetiamarhadi/udemy-grpc-go/calculator/proto"
)

func callPrimes(c pb.CalculatorServiceClient) {

	log.Println("primes was invoked")

	req := &pb.PrimeRequest{
		Number: 12390392840,
	}

	stream, err := c.Primes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling primes: %v\n", err)
	}

	for {

		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v\n", err)
		}

		log.Printf("Primes: %d\n", res.Result)
	}
}
