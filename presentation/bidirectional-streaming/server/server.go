package main

import (
	bidirectional_streaming_proto "grpc-playground/bidirectional-streaming/proto"
	"io"
	"net"

	"google.golang.org/grpc"
)

// Server implements proto.CalculatorServiceServer
type Server struct {
	bidirectional_streaming_proto.UnimplementedCalculatorServiceServer
}

// Max implements bidirectional_streaming_proto.CalculatorServiceServer.
func (s *Server) Max(stream bidirectional_streaming_proto.CalculatorService_MaxServer) error {
	max := int32(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		} else if err != nil {
			panic(err)
		}
		num := req.GetNum()
		if num > max {
			max = num
		}
		stream.Send(&bidirectional_streaming_proto.Response{Num: max})
	}
}

// NewServer creates a new instance of Server
func NewServer() bidirectional_streaming_proto.CalculatorServiceServer {
	return &Server{}
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	bidirectional_streaming_proto.RegisterCalculatorServiceServer(grpcServer, NewServer())

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
