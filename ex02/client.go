package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

func ping(clientN int, protocol, addr, text string, wg *sync.WaitGroup) {
	defer wg.Done()

	conn, err := net.Dial(protocol, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for i := 0; i < 10000; i++ {
		start := time.Now()

		fmt.Fprintf(conn, text)
		buffer := make([]byte, 1024)
		_, err = conn.Read(buffer)
		if err != nil {
			log.Fatal(err)
		}

		elapsed := time.Since(start)
		fmt.Printf("%d:%s\n", clientN, elapsed)
	}
}

func main() {
	args := os.Args[1:]
	threads, err := strconv.Atoi(args[0])
	protocol := args[1]
	addr := args[2]
	text := args[3]

	if err != nil {
		log.Fatal(err)
	}
	var wg sync.WaitGroup
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go ping(i, protocol, addr, text, &wg)
	}
	wg.Wait()
}
