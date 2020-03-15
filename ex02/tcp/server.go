package main

import (
	"log"
	"net"
)

func catHandler(conn net.Conn) {
	for {
		message := make([]byte, 4096)
		length, err := conn.Read(message)
		if err != nil {
			break
		}
		if length > 0 {
			_, err = conn.Write(message)
			if err != nil {
				break
			}
		}
	}
	conn.Close()
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
  defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
		}
		go catHandler(conn)
	}
}
