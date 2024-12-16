package main

import (
	"fmt"
	client_streaming_proto "grpc-playground/client-streaming/proto"
	"io"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"
)

// Server implements proto.CalculatorServiceServer
type Server struct {
	client_streaming_proto.UnimplementedCalculatorServiceServer
}

// Average implements client_streaming_proto.CalculatorServiceServer.
func (s *Server) Average(stream client_streaming_proto.CalculatorService_AverageServer) error {
	var sum float64 = 0
	var cnt int = 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			res := sum / float64(cnt)

			resp := &client_streaming_proto.Response{
				Num: fmt.Sprintf("%f", res),
			}
			log.Printf("send %v", resp)
			return stream.SendAndClose(resp)
		} else if err != nil {
			panic(err)
		}
		log.Printf("receive %v", req)

		cnt++
		n, _ := strconv.Atoi(req.Num)
		sum += float64(n)
	}
}

// NewServer creates a new instance of Server
func NewServer() client_streaming_proto.CalculatorServiceServer {
	return &Server{}
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	client_streaming_proto.RegisterCalculatorServiceServer(grpcServer, NewServer())

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
