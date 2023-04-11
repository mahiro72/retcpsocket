package main

import (
	"log"
	"net"
)

func main() {
	log.Println("server init")

	// syscall (socket, bind, listen)
	ln,err := net.Listen("tcp",":8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		// syscall (accept)
		conn,err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		buf := make([]byte,1024)
		_,err = conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		conn.Close()

		log.Println("receive message: ",string(buf))
	}
}