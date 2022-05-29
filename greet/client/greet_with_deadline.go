package main

import (
	"context"
	"log"
	"time"

	pb "github.com/adetiamarhadi/udemy-grpc-go/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func callGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {

	log.Println("Calling GreetWithDeadline function")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	defer cancel()

	req := &pb.GreetRequest{
		FirstName: "Adet",
	}

	res, err := c.GreetWithDeadline(ctx, req)

	if err != nil {

		e, ok := status.FromError(err)

		if ok {

			if e.Code() == codes.DeadlineExceeded {

				log.Println("Deadline exceeded!")

				return
			} else {

				log.Fatalf("Unexpected gRPC error: %v\n", err)
			}
		} else {

			log.Fatalf("A non gRPC error: %v\n", err)
		}
	}

	log.Printf("GreetWithDeadline: %s\n", res.Result)
}
