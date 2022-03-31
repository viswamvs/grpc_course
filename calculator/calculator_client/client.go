package main

import (
	"context"
	"fmt"
	calculator "grpc_course/calculator/calculatorpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("calculator client ")

	conn, err := grpc.Dial("localhost:50051",  grpc.WithInsecure())

	if err != nil {
		log.Fatalf("couldn't connect: %v", err)
	}

	defer conn.Close()
	c := calculator.NewCalculatorServiceClient(conn)

	//fmt.Printf("client created: %f", c)
	doUnary(c)
}

func doUnary(c calculator.CalculatorServiceClient){
	req := &calculator.SumRequest{
		Num1: 5,
		Num2: 5,
		
	}
	res, err := c.Sum(context.Background(), req)

	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Printf("Response: %v", res.Result)
}