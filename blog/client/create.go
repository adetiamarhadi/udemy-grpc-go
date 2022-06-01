package main

import (
	"context"
	"log"

	pb "github.com/adetiamarhadi/udemy-grpc-go/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {

	log.Println("Call CreateBlog function")

	blog := &pb.Blog{
		AuthorId: "Adet",
		Title:    "Pemrograman Go - Pemula",
		Content:  "Bahasa Pemrograman Go",
	}

	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog has been created: %s\n", res.Id)

	return res.Id
}
