package main

import (
	"context"
	"fmt"
	greet "grpc_course/greet/greetpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{
	greet.UnimplementedGreetServiceServer
}
func (*server) Greet(ctx context.Context,req  *greet.GreetRequest) (*greet.GreetResponse, error){
	fmt.Printf("Greet function was invoked with %v", req)
	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()
	result := "Hello " + firstName + " " + lastName
	res := &greet.GreetResponse{
		Result: result,
	}
	return res,nil
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