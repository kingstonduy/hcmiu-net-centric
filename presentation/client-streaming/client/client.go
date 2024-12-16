package main

import (
	"context"
	"fmt"
	client_streaming_proto "grpc-playground/client-streaming/proto"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Define a list of numbers to send to the server
	numbers := []int32{10, 20, 30, 40, 50}

	// Establish gRPC connection
	var conn *grpc.ClientConn
	conn, err := grpc.NewClient(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create the gRPC client
	c := client_streaming_proto.NewCalculatorServiceClient(conn)

	// Call the Average RPC
	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("Error during RPC call: %v", err)
	}

	// Send numbers to the server
	for _, num := range numbers {
		log.Printf("Sending number: %d", num)

		numStr := fmt.Sprintf("%d", num)
		req := &client_streaming_proto.Request{
			Num: numStr,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Failed to send number: %v", err)
		}
		time.Sleep(500 * time.Millisecond) // Optional: Simulate delay between sends
	}

	// Close the stream and receive the response
	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	// Print the calculated average
	log.Printf("The average of the numbers is: %s", response.GetNum())
}
