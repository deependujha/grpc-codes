package main

import (
	"context"
	"flag"
	"fmt"
	proto "grpc-server/proto"
	"io"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	proto.UnimplementedDeepSrvServer
}

func (s *server) AddTwoNum(ctx context.Context, in *proto.NumReq) (*proto.NumResp, error) {
	fmt.Printf("AddTwoNum: Got number: %d + %d\n", in.A, in.B)
	num1 := in.A
	num2 := in.B
	return &proto.NumResp{C: num1 + num2}, nil
}

func (s *server) AddAllTheseNum(stream proto.DeepSrv_AddAllTheseNumServer) error {
	var curr_sum int32
	fmt.Println(strings.Repeat("=", 60))
	for {
		num, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&proto.NumResp{C: curr_sum})
		}
		if err != nil {
			return err
		}
		curr_sum += num.A
		fmt.Printf("Got number: %d; And curr sum: %d\n", num.A, curr_sum)
	}
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	proto.RegisterDeepSrvServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
