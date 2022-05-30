package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/adetiamarhadi/udemy-grpc-go/greet/proto"
)

var addr string = "localhost:50051"

func main() {

	opts := []grpc.DialOption{}

	tls := true
	if tls {

		certFile := "ssl/ca.crt"

		creds, err := credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			log.Fatalf("Error while loading CA trust certificate: %v\n", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	// conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	callGreet(c)
	// callGreetManyTimes(c)
	// callLongGreet(c)
	// callGreetEveryone(c)
	// callGreetWithDeadline(c, 1*time.Second)
}
