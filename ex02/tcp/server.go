package main

import (
	"log"
	"net"
	"strconv"
)

func fib(n int) int {
	if n <= 0 {
		return -1
	} else if n <= 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func handler(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 4)
	length, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	if length > 0 {
		n, err := strconv.Atoi(string(buf[:length]))
		if err != nil {
			log.Fatal(err)
		}
		f := fib(n)
		r := strconv.Itoa(f)
		_, err = conn.Write([]byte(r))
		if err != nil {
			log.Fatal(err)
		}
	}
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
		go handler(conn)
	}
}
