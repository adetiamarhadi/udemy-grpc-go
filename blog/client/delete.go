package main

import (
	"context"
	"log"

	pb "github.com/adetiamarhadi/udemy-grpc-go/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, id string) {

	log.Printf("Call DeleteBlog with id: %s\n", id)

	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})
	if err != nil {
		log.Fatalf("Error while deleting: %v\n", err)
	}

	log.Printf("Blog with id %s successfully deleted\n", id)
}
