package main

import (
	"context"
	"log"
	"net"

	pb "github.com/adetiamarhadi/udemy-grpc-go/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var collection *mongo.Collection

var addr string = "0.0.0.0:50055"

type Server struct {
	pb.BlogServiceServer
}

func main() {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		log.Fatal(err)
	}

	if err = client.Connect(context.Background()); err != nil {
		log.Fatal(err)
	}

	collection = client.Database("blogdb").Collection("blog")

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Fail listen: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()

	pb.RegisterBlogServiceServer(s, &Server{})

	reflection.Register(s)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Fail Serve: %v\n", err)
	}
}
