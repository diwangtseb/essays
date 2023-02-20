# directory
- [directory](#directory)
  - [build simple rpc](#build-simple-rpc)
    - [target](#target)
    - [implement the following steps](#implement-the-following-steps)
  - [ref](#ref)

## build simple rpc

### target
- client => server by self rpc

### implement the following steps
1. session
```go
package rpc

import (
	"encoding/binary"
	"net"
)

func NewSession(conn net.Conn, sessionOtps ...FuncSessionOpt) *Session {
	sessionOpt := sessionOpt{}
	for _, v := range sessionOtps {
		v.apply(&sessionOpt)
	}
	return &Session{
		conn:       conn,
		sessionOpt: sessionOpt,
	}
}

type Session struct {
	conn       net.Conn
	sessionOpt sessionOpt
}

func (s *Session) Write(data []byte) {
	headerLen := s.sessionOpt.len
	bodLen := len(data)
	allLen := headerLen + bodLen
	buffer := make([]byte, allLen)
	binary.BigEndian.PutUint32(buffer[:s.sessionOpt.len], uint32(allLen))
	copy(buffer[s.sessionOpt.len:], data)
	s.conn.Write(buffer)
}

func (s *Session) Read() []byte {
	headerLen := s.sessionOpt.len
	headBuffer := make([]byte, headerLen)
	_, err := s.conn.Read(headBuffer)
	if err != nil {
		panic(err)
	}
	lens := binary.BigEndian.Uint32(headBuffer)
	data := make([]byte, lens)
	_, err = s.conn.Read(data)
	if err != nil {
		panic(err)
	}
	return data
}

type SessionApplier interface {
	apply(*sessionOpt)
}

type sessionOpt struct {
	len int
}

type FuncSessionOpt func(*sessionOpt)

func (so FuncSessionOpt) apply(s *sessionOpt) {
	so(s)
}

// len = 4
func WithLens(len int) FuncSessionOpt {
	return func(so *sessionOpt) {
		if len != 4 {
			len = 4
		}
		so.len = len
	}
}

```
2. server
```go
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

```
3. client
```go
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

```
4. call rpc methods
```go
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

```


## ref
[code](github.com/diwangtseb/essayes/go)
