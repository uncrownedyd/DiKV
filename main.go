package main

import (
	"log"
	"net"

	"github.com/DiKV/api/dikv"
	"github.com/DiKV/api/dikv/proto"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Panicf("failed to listen 12345: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterDiKVServer(s, dikv.New())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
