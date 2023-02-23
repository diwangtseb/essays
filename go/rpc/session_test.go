package rpc

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestRpcSession(t *testing.T) {
	go func() {
		lis, err := net.Listen("tcp", ":1234")
		if err != nil {
			panic(err)
		}
		conn, err := lis.Accept()
		if err != nil {
			panic(err)
		}
		session := NewSession(conn, WithLens(1))
		session.Write([]byte("world"))
		time.Sleep(time.Second * 1)
	}()
	go func() {
		time.Sleep(time.Second * 2)
		conn, err := net.Dial("tcp", ":1234")
		if err != nil {
			panic(err)
		}
		session := NewSession(conn, WithLens(1))
		r := session.Read()
		fmt.Println(string(r))
	}()
	time.Sleep(time.Second * 5)
}
