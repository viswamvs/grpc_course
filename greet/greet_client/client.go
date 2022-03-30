package main

import (
	"fmt"
	greet "grpc_course/greet/greetpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Print("client  ")

	conn, err := grpc.Dial("localhost:50051",  grpc.WithInsecure())

	if err != nil {
		log.Fatalf("couldn't connect: %v", err)
	}

	defer conn.Close()
	c := greet.NewGreetServiceClient(conn)

	fmt.Printf("client created: %f", c)
}