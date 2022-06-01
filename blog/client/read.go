package main

import (
	"context"
	"log"

	pb "github.com/adetiamarhadi/udemy-grpc-go/blog/proto"
)

func readBlog(c pb.BlogServiceClient, id string) {

	log.Printf("Calling ReadBlog function with id: %s\n", id)

	req := &pb.BlogId{Id: id}

	res, err := c.ReadBlog(context.Background(), req)
	if err != nil {
		log.Printf("Error happened while reading: %v\n", err)
	}

	log.Printf("blog with id %s: %v\n", id, res)
}
