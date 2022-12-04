package main

import (
	"github.com/a11en4sec/program-pattern/FunctionalOptions/server"
	"time"
)

func newSerWith() {
	sb := server.ServerBuilder{}
	server := sb.Create("127.0.0.1", 8990).
		WithProtocol("udp").
		WithMaxConn(10).
		WithTimeOut(3 * time.Second).
		Build()

	_ = server
}

func newSerWithOption() {
	s1, _ := server.NewServer("127.0.0.1", 8877)
	s2, _ := server.NewServer("127.0.0.1", 8877, server.Protocal("udp"))
	s3, _ := server.NewServer("127.0.0.1", 8877, server.MaxConns(10))

	_ = s1
	_ = s2
	_ = s3

}

func main() {
	newSerWith()
	newSerWithOption()
}
