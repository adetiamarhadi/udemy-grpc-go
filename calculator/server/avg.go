package main

import (
	"io"
	"log"

	pb "github.com/adetiamarhadi/udemy-grpc-go/calculator/proto"
)

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {

	log.Println("Avg Function invoked")

	var sum int32 = 0
	count := 0

	for {

		req, err := stream.Recv()

		if err == io.EOF {

			return stream.SendAndClose(&pb.AvgResponse{
				Result: float64(sum) / float64(count),
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		log.Printf("Receiving number: %d\n", req.Number)

		sum += req.Number
		count++
	}
}
