package main

import (
	"context"
	"log"

	pb "github.com/adetiamarhadi/udemy-grpc-go/calculator/proto"
)

func callAvg(c pb.CalculatorServiceClient) {

	log.Println("Calling avg function")

	stream, err := c.Avg(context.Background())
	if err != nil {
		log.Fatalf("Error while opening the stream: %v\n", err)
	}

	numbers := []int32{3, 5, 9, 54, 23}
	for _, number := range numbers {

		log.Printf("Sending number: %d\n", number)

		stream.Send(&pb.AvgRequest{
			Number: number,
		})
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response: %v\n", err)
	}

	log.Printf("Avg: %f\n", res.Result)
}
