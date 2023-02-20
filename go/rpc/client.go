package rpc

import (
	"net"
	"reflect"
)

type Client struct {
	conn net.Conn
}

func NewClient(conn net.Conn) *Client {
	return &Client{
		conn: conn,
	}
}

func (c *Client) callRpc(name string, fPtr interface{}) {
	fn := reflect.ValueOf(fPtr).Elem()
	f := func(args []reflect.Value) []reflect.Value {
		inArgs := make([]interface{}, 0, len(args))
		for _, arg := range args {
			inArgs = append(inArgs, arg.Interface())
		}
		session := NewSession(c.conn, WithLens(1))
		req := RpcData{
			Name: name,
			Args: inArgs,
		}
		b, err := encode(req)
		if err != nil {
			panic(err)
		}
		session.Write(b)
		bs := session.Read()
		rspData, err := decode(bs)
		if err != nil {
			panic(err)
		}
		outArgs := make([]reflect.Value, 0, len(rspData.Args))
		for i, arg := range rspData.Args {
			if arg == nil {
				outArgs = append(outArgs, reflect.Zero(fn.Type().Out(i)))
				continue
			}
			outArgs = append(outArgs, reflect.ValueOf(arg))
		}
		return outArgs
	}
	v := reflect.MakeFunc(fn.Type(), f)
	fn.Set(v)
}
