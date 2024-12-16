package main

import (
	server_streaming_proto "grpc-playground/server-streaming/proto"
	"net"

	"google.golang.org/grpc"
)

// Server implements proto.CalculatorServiceServer
type Server struct {
	server_streaming_proto.UnimplementedCalculatorServiceServer
}

// PrimeNumberDecomposition implements server_streaming_proto.CalculatorServiceServer.
func (s *Server) PrimeNumberDecomposition(req *server_streaming_proto.Request, stream server_streaming_proto.CalculatorService_PrimeNumberDecompositionServer) error {
	n := (int)(req.GetNum1())
	if n <= 1 {
		return nil
	}

	divisor := 2

	// Iterate through possible divisors
	for n > 1 {
		for n%divisor == 0 {
			stream.Send(&server_streaming_proto.Response{Result: int32(divisor)})
			n /= divisor
		}
		divisor++
		// Optimization: Stop checking beyond the square root of the number
		if divisor*divisor > n && n > 1 {
			stream.Send(&server_streaming_proto.Response{Result: int32(n)})
			break
		}
	}

	return nil
}

// NewServer creates a new instance of Server
func NewServer() server_streaming_proto.CalculatorServiceServer {
	return &Server{}
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	server_streaming_proto.RegisterCalculatorServiceServer(grpcServer, NewServer())

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
