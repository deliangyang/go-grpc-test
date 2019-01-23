package main

import (
	"context"
	"github.com/deliangyang/tt2/amath"
	"google.golang.org/grpc"
	pb "github.com/deliangyang/tt2/proto"
	"log"
)

const (
	address = "0.0.0.0:5000"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: amath.Name})
	if err != nil {
		panic(err)
	}
	log.Println(r.Message)

}
