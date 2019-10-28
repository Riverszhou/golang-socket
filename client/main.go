package main

import (
	"io"
	"log"
	"net"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	l, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Panic(err)
	}

	for {
		client, err := l.Accept()
		if err != nil {
			log.Panic(err)
		}

		go handleClientRequest(client)
	}
}

func handleClientRequest(client net.Conn) {
	if client == nil {
		return
	}
	defer client.Close()
	server, err := net.Dial("tcp", net.JoinHostPort("127.0.0.1", "8081"))
	if err != nil {
		log.Println(err)
		return
	}
	defer server.Close()
	go io.Copy(server, client)
	io.Copy(client, server)
}
