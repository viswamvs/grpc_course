package main

import (
	"fmt"
	greet "grpc_course/greet/greetpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{
	greet.UnimplementedGreetServiceServer
}

func main() {
	fmt.Print("hello")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	 
	if err != nil {
		log.Fatalf("error %v", err)
	}

	s := grpc.NewServer()
	greet.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}

}