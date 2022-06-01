package main

import (
	"context"
	"log"

	pb "github.com/adetiamarhadi/udemy-grpc-go/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {

	log.Println("Call UpdateBlog function")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Ayu",
		Title:    "HTML5 & CSS Tips Tricks",
		Content:  "Modern Website",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		log.Fatalf("Error happened while updating: %v\n", err)
	}

	log.Println("Blog was updated!")
}
