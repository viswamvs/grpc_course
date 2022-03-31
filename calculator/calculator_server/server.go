package main

import (
	"fmt"
	calculator "grpc_course/calculator/calculatorpb"
	"log"
	"net"
	"context"
	"google.golang.org/grpc"
)

type server struct{
	calculator.UnimplementedCalculatorServiceServer
}

func (*server) Sum(ctx context.Context, req *calculator.SumRequest) (*calculator.SumResponse, error){
	fmt.Printf("Received Sum RPC: %v", req)
	a := req.Num1
	b := req.Num2
	sum := a + b

	res := &calculator.SumResponse{
		Result: sum,
	}

	return res, nil
}

func main() {
	fmt.Print("Calculator Running")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	s := grpc.NewServer()
	calculator.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}