package service

import (
	"context"
	"github.com/deliangyang/tt2/model"
	pb "github.com/deliangyang/tt2/proto"
	"strconv"
)

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	id, _ := strconv.Atoi(in.Name)
	product, err := model.GetUserInfoByUserId(id)
	return &pb.HelloReply{Message: product.Name}, err
}