package dikv

import (
	"context"

	"github.com/DiKV/api/dikv/proto"
)

type Server struct {
	proto.UnimplementedDiKVServer
}

func (s *Server) Put(ctx context.Context, request *proto.PutRequest) (*proto.PutReply, error) {
	return &proto.PutReply{
		Code:    200,
		Message: "Success",
	}, nil
}

func (s *Server) Get(ctx context.Context, request *proto.GetRequest) (*proto.GetReply, error) {
	return &proto.GetReply{Code: 200, Message: "Success"}, nil
}

func New() *Server {
	return &Server{}
}
