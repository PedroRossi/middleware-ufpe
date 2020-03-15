package main

import (
	"os"
	"fmt"
	"log"
	"net"
	"sync"
  "time"
	"strconv"
)

func ping(protocol, addr, text string, wg *sync.WaitGroup) {
	defer wg.Done()

	conn, err := net.Dial(protocol, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Fprintf(conn, text)
	buffer := make([]byte, 1024)
	_, err = conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
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
  start := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go ping(protocol, addr, text, &wg)
	}
	wg.Wait()
  elapsed := time.Since(start)
  fmt.Printf("%s\n", elapsed)
}
