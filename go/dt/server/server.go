package server

import (
	"github.com/diwangtseb/essayes/go/dt/pb"
	"github.com/diwangtseb/essayes/go/dt/service"
	"google.golang.org/grpc"
)

func NewGrpcServer() *grpc.Server {
	srv := grpc.NewServer()
	service := &service.TransService{}
	pb.RegisterTransServiceServer(srv, service)
	return srv
}
