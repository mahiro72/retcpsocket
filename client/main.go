package main

import (
	"log"
	"net"
	"os"
)

func main(){
	log.Println("client init")

	if len(os.Args) != 2 {
		panic("err: args length is not 2")
	}
	msg := os.Args[1]

	conn,err := net.Dial("tcp","localhost:8080")
	if err != nil {
		panic(err)
	}

	_,err = conn.Write([]byte(msg))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("send message success")
}