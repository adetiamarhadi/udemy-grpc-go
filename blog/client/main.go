package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/adetiamarhadi/udemy-grpc-go/blog/proto"
)

var addr string = "localhost:50055"

func main() {

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("can't connect to %s: %v/n", addr, err)
	}

	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)

	id := createBlog(c)
	readBlog(c, id)
	updateBlog(c, id)
	listBlogs(c)
	deleteBlog(c, id)
}
