package main

import (
	"log"
	"net"
	"net/rpc"
)

type Listener int

func (l *Listener) GetLine(line []byte, ack *bool) error {
	log.Println(string(line))
	return nil
}

func main() {
	addy, err := net.ResolveTCPAddr("tcp", "0.0.0.0:42586")
	if err != nil {
		log.Fatal(err)
	}

	inbound, err := net.ListenTCP("tcp", addy)
	if err != nil {
		log.Fatal(err)
	}
	
	log.Println("Success to listen!")

	listener := new(Listener)
	rpc.Register(listener)
	rpc.Accept(inbound)
}
