package main

import (
	"context"
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

	//fmt.Printf("client created: %f", c)
	doUnary(c)
}

func doUnary(c greet.GreetServiceClient){
	req := &greet.GreetRequest{
		Greeting: &greet.Greeting{
			FirstName: "viswa",
			LastName: "mohan",
		},
		
	}
	res, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Printf("Response: %v", res.Result)
}