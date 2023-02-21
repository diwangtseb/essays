package server

import (
	"net"
	"testing"
)

func TestServer(t *testing.T) {
	srv := NewGrpcServer()
	lis, _ := net.Listen("tcp", ":1234")
	srv.Serve(lis)
}
