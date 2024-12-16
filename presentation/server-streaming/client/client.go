package main

import (
	"bufio"
	"context"
	"fmt"
	server_streaming_proto "grpc-playground/server-streaming/proto"
	"io"
	"log"
	"os"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Prompt the user for input
	fmt.Print("Enter a number for prime decomposition: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	// Convert input to an integer
	num, err := strconv.Atoi(input[:len(input)-1]) // Remove the newline character
	if err != nil {
		log.Fatalf("Invalid input, please enter a valid number: %v", err)
	}

	// Establish gRPC connection
	var conn *grpc.ClientConn
	conn, err = grpc.NewClient(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create the gRPC client
	c := server_streaming_proto.NewCalculatorServiceClient(conn)

	// Prepare the request
	message := server_streaming_proto.Request{
		Num1: int32(num), // Ensure compatibility with the proto definition
	}

	// Call the PrimeNumberDecomposition RPC
	stream, err := c.PrimeNumberDecomposition(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error during RPC call: %v", err)
	}

	// Receive the stream of prime numbers from the server
	fmt.Printf("Prime factors of %d:\n", num)
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			log.Println("Server finished streaming")
			break
		}
		if err != nil {
			log.Fatalf("Error while receiving stream: %v", err)
		}
		fmt.Printf("%d\n", resp.GetResult())
	}
}
