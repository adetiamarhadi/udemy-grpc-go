package main

import (
	"context"
	"io"
	"log"

	pb "github.com/adetiamarhadi/udemy-grpc-go/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func listBlogs(c pb.BlogServiceClient) {

	log.Println("Calling ListBlogs function")

	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Error while calling list blogs function: %v\n", err)
	}

	for {

		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Something wrong happened: %v\n", err)
		}

		log.Printf("blog: %v\n", res)
	}
}
