package main

import (
	"log"
	"net"
)

func main() {
	server := newServer()
	go server.run()

	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}
	defer listener.Close()

	log.Printf("starting server on :8888")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("unable to accept connection: %s", err.Error())
			continue
		}

		client := newClient(conn, server.commands)
		go client.readInput()
	}
}
