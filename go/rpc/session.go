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
