package server

import (
	"crypto/tls"
	"time"
)

type Server struct {
	Addr string
	Port int
	//Conf *Config
	Protocol string
	Timeout  time.Duration
	MaxCons  int
	TLS      *tls.Config
}

//type Config struct {
//	Protocol string
//	Timeout  time.Duration
//	MaxCons  int
//	TLS      *tls.Config
//}
//
//func NewServer(addr string, port int, conf *Config) (*Server, error) {
//	return &Server{
//		Addr: addr,
//		Port: port,
//		Conf: conf,
//	}, nil
//}
