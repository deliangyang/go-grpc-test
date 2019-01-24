package main

import (
	"context"
	"fmt"
	"github.com/deliangyang/tt2/conections"
	"github.com/deliangyang/tt2/model"
	pb "github.com/deliangyang/tt2/proto"
	"google.golang.org/grpc"
	"net"
	"strconv"
)

const (
	port = ":5000"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	id, _ := strconv.Atoi(in.Name)
	product, err := model.GetUserInfoByUserId(id)
	fmt.Println(product)
	return &pb.HelloReply{Message: product.Name}, err
}

func main() {
	// fmt.Println(go_test.Hi("hello world"))
	conections.ConnectDB()
	var count int
	conections.DB.Model(&model.Product{}).Count(&count)
	product := model.Product{
		UserId: count + 1,
		Name:   "hello world",
	}
	conections.DB.Create(&product)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	defer conections.DB.Close()
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	s.Serve(lis)

}
