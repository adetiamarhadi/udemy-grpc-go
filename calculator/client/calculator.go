package main

import (
	"context"
	"log"

	pb "github.com/adetiamarhadi/udemy-grpc-go/calculator/proto"
)

func callSum(c pb.CalculatorServiceClient) {

	log.Println("Calling sum function.")

	res, err := c.Sum(context.Background(), &pb.SumRequest{
		Number_1: 10,
		Number_2: 13,
	})

	if err != nil {
		log.Fatalf("Get error: %v\n", err)
	}

	log.Printf("Sum Result: %d\n", res.Result)
}
