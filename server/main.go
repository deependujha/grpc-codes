package main

import (
	"context"
	"flag"
	"fmt"
	proto "grpc-server/proto"
	"log"
	"net"

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
	fmt.Println("AddTwoNum")
	num1 := in.A
	num2 := in.B
	return &proto.NumResp{C: num1 + num2}, nil
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
