package main

import (
	"net"

	"github.com/diwangtseb/essayes/go/dt/server"
)

func main() {
	srv := server.NewGrpcServer()
	lis, _ := net.Listen("tcp", "0.0.0.0:1234")
	if err := srv.Serve(lis); err != nil {
		panic(err)
	}
}
