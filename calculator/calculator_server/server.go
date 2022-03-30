package main

import (
	"fmt"
	calculator "grpc_course/calculator/calculatorpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{
	calculator.UnimplementedCalculatorServiceServer
}
func main() {
	fmt.Print("hello")

	lis, err := net.Listen("tcp", "0.0.0.0:8000")

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	s := grpc.NewServer()
	calculator.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}