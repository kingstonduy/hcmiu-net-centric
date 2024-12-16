package main

import (
	"bufio"
	"context"
	"fmt"
	bidirectional_streaming_proto "grpc-playground/bidirectional-streaming/proto"
	"io"
	"log"
	"os"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Establish gRPC connection
	var conn *grpc.ClientConn
	conn, err := grpc.NewClient(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create the gRPC client
	c := bidirectional_streaming_proto.NewCalculatorServiceClient(conn)

	// Call the Max RPC (assume Max is the RPC handling bidirectional streaming)
	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Error during RPC call: %v", err)
	}

	// Create a channel to handle user input and server response
	done := make(chan struct{})

	// Goroutine to send user input to the server
	go func() {
		defer func() {
			_ = stream.CloseSend()
		}()

		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter numbers (type 'done' to finish):")

		for {
			// Read user input
			fmt.Print("> ")
			input, err := reader.ReadString('\n')
			if err != nil {
				log.Fatalf("Error reading input: %v", err)
			}

			// Trim whitespace and check for termination
			input = input[:len(input)-1]
			if input == "done" {
				break
			}

			// Convert input to int
			num, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Invalid input, please enter a number.")
				continue
			}

			// Send the number to the server
			req := &bidirectional_streaming_proto.Request{
				Num: int32(num),
			}
			if err := stream.Send(req); err != nil {
				log.Fatalf("Failed to send number: %v", err)
			}
			// time.Sleep(500 * time.Millisecond) // Optional delay for better interaction
		}
	}()

	// Goroutine to receive server responses
	go func() {
		for {
			// Receive the server response
			resp, err := stream.Recv()
			if err != nil {
				if err == context.Canceled {
					break
				}
				if err == io.EOF {
					log.Println("Server finished sending responses.")
					break
				}
				log.Fatalf("Error receiving response: %v", err)
			}

			// Print the maximum number so far
			fmt.Printf("Current maximum: %d\n", resp.GetNum())
		}
		close(done)
	}()

	// Wait for both goroutines to finish
	<-done
	fmt.Println("Client finished.")
}
