package main

import (
	"context"
	"google.golang.org/grpc"
	pb "github.com/deliangyang/tt2/proto"
	"net"
)

const (
	port = ":5000"
)

type server struct {}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "hello world"}, nil
}

func main() {
	// fmt.Println(go_test.Hi("hello world"))
	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	s.Serve(lis)

}
