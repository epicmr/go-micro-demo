package main

import (
	"log"
	"net"
	"net/rpc"
)

type EchoService struct{}

func (p *EchoService) Echo(request string, reply *string) error {
	*reply = "hello, " + request
	return nil
}

func main() {
	rpc.RegisterName("EchoService", new(EchoService))

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("net.Listen failed. [%v]", err)
	}
	conn, err := listener.Accept()
	if err != nil {
		log.Fatalf("listener.Accept failed. [%v]", err)
	}
	rpc.ServeConn(conn)
}
