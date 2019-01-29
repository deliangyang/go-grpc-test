package main

import (
	"github.com/deliangyang/tt2/conections"
	pb "github.com/deliangyang/tt2/proto"
	"github.com/deliangyang/tt2/service"
	"google.golang.org/grpc"
	"net"
)

const (
	port = ":5000"
)

func main() {
	conections.ConnectDB()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	defer conections.DB.Close()
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &service.Server{})
	s.Serve(lis)

}
