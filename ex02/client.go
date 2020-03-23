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

func ping(protocol, addr string, n int, wg *sync.WaitGroup) {
	defer wg.Done()

	start := time.Now()
	conn, err := net.Dial(protocol, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte(strconv.Itoa(n)))
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 4)
	_, err = conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	elapsed := time.Since(start)
	fmt.Printf("%s\n", elapsed)
}

func main() {
	args := os.Args[1:]
	threads, err := strconv.Atoi(args[0])
	protocol := args[1]
	addr := args[2]
	n, err := strconv.Atoi(args[3])
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go ping(protocol, addr, n, &wg)
	}
	wg.Wait()
}
