package main

import (
	"bufio"
	"context"
	"fmt"
	calculator_proto "grpc-playground/unary/proto"
	"log"
	"os"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Prompt the user for input numbers
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter first number (Num1): ")
	input1, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	num1, err := strconv.Atoi(input1[:len(input1)-1]) // Remove the newline character
	if err != nil {
		log.Fatalf("Invalid input for Num1, please enter a valid number: %v", err)
	}

	fmt.Print("Enter second number (Num2): ")
	input2, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	num2, err := strconv.Atoi(input2[:len(input2)-1]) // Remove the newline character
	if err != nil {
		log.Fatalf("Invalid input for Num2, please enter a valid number: %v", err)
	}

	// Establish gRPC connection
	var conn *grpc.ClientConn
	conn, err = grpc.NewClient(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create the gRPC client
	c := calculator_proto.NewCalculatorServiceClient(conn)

	// Prepare the request
	message := calculator_proto.SumRequest{
		Num1: int32(num1),
		Num2: int32(num2),
	}

	// Call the Sum RPC
	response, err := c.Sum(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error during RPC call: %v", err)
	}

	// Display the result
	fmt.Printf("Response from Server: %d + %d = %d\n", num1, num2, response.GetResult())
}
