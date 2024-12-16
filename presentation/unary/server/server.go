package main

import (
	"context"
	calculator_proto "grpc-playground/unary/proto"
	"net"

	"google.golang.org/grpc"
)

// Server implements proto.CalculatorServiceServer
type Server struct {
	calculator_proto.UnimplementedCalculatorServiceServer // Embed this for compatibility
}

// Sum implements proto.CalculatorServiceServer
func (s *Server) Sum(ctx context.Context, req *calculator_proto.SumRequest) (*calculator_proto.SumResponse, error) {
	return &calculator_proto.SumResponse{
		Result: req.GetNum1() + req.GetNum2(),
	}, nil
}

// NewServer creates a new instance of Server
func NewServer() calculator_proto.CalculatorServiceServer {
	return &Server{}
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	calculator_proto.RegisterCalculatorServiceServer(grpcServer, NewServer())

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
