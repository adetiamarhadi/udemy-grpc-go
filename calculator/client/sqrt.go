package main

import (
	"context"
	"log"

	pb "github.com/adetiamarhadi/udemy-grpc-go/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func callSqrt(c pb.CalculatorServiceClient, n int32) {

	log.Println("Calling Sqrt function")

	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{
		Number: n,
	})

	if err != nil {

		e, ok := status.FromError(err)

		if ok {

			log.Printf("Error message from server: %s\n", e.Message())
			log.Printf("Errro code from server: %s\n", e.Code())

			if e.Code() == codes.InvalidArgument {

				log.Println("We probably sent a negative number")
				return
			}
		} else {

			log.Fatalf("A non gRPC error: %v\n", err)
		}
	}

	log.Printf("Sqrt: %f\n", res.Result)
}
