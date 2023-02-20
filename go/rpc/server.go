package rpc

import (
	"net"
	"reflect"
)

type Server struct {
	addr  string
	funcs map[string]reflect.Value
}

func NewServer(addr string) *Server {
	return &Server{
		addr:  addr,
		funcs: map[string]reflect.Value{},
	}
}

func (s *Server) Register(name string, f interface{}) {
	s.funcs[name] = reflect.ValueOf(f)
}

func (s *Server) Run() {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := lis.Accept()
		if err != nil {
			panic(err)
		}
		session := NewSession(conn, WithLens(1))
		bs := session.Read()
		data, err := decode(bs)
		if err != nil {
			return
		}
		f, ok := s.funcs[data.Name]
		if !ok {
			return
		}
		inArgs := make([]reflect.Value, 0, len(data.Args))
		for _, arg := range data.Args {
			inArgs = append(inArgs, reflect.ValueOf(arg))
		}
		out := f.Call(inArgs)
		outArgs := make([]interface{}, 0, len(out))
		for _, o := range out {
			outArgs = append(outArgs, o.Interface())
		}
		respData := RpcData{
			Name: data.Name,
			Args: outArgs,
		}
		rbs, err := encode(respData)
		if err != nil {
			return
		}
		session.Write(rbs)
	}
}

type User struct {
	Name string
}

func Quser() User {
	return User{
		Name: "rsp --",
	}
}
