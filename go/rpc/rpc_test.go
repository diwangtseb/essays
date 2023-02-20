package rpc

import (
	"encoding/gob"
	"fmt"
	"net"
	"testing"
	"time"
)

func TestRpc(t *testing.T) {
	s := NewServer(":1234")
	s.Register("quser", Quser)
	gob.Register(User{})
	go s.Run()
	time.Sleep(time.Second * 1)
	go func() {
		conn, err := net.Dial("tcp", ":1234")
		if err != nil {
			panic(err)
		}
		c := NewClient(conn)
		var q func() User
		c.callRpc("quser", &q)
		r := q()
		fmt.Println(r)
	}()
	time.Sleep(time.Second * 5)
}
