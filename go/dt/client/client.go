package client

import (
	"github.com/diwangtseb/essayes/go/dt/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewTransClient(addr string) pb.TransServiceClient {
	opt := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(addr, opt)
	if err != nil {
		panic(err)
	}
	return pb.NewTransServiceClient(conn)
}
