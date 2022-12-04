package server

import (
	"crypto/tls"
	"time"
)

// Option 函数式编程
type Option func(server *Server)

func Protocal(p string) Option {
	return func(s *Server) {
		s.Protocol = p
	}
}

func Timeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}

func MaxConns(maxconns int) Option {
	return func(s *Server) {
		s.MaxCons = maxconns
	}
}

func TLS(tls *tls.Config) Option {
	return func(s *Server) {
		s.TLS = tls
	}
}

func NewServer(addr string, port int, options ...func(server *Server)) (*Server, error) {
	srv := Server{
		Addr:     addr,
		Port:     port,
		Protocol: "tcp",
		Timeout:  30 * time.Second,
		MaxCons:  1000,
		TLS:      nil,
	}

	for _, option := range options {
		option(&srv)
	}

	return &srv, nil
}
