package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"google.golang.org/grpc"

	pb "grpc-cat/cat"
)

func ping(addr, text string, wg *sync.WaitGroup) {
	defer wg.Done()

	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewCatClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = c.Reply(ctx, &pb.Package{Message: text})
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	args := os.Args[1:]
	threads, err := strconv.Atoi(args[0])
	addr := args[1]
	text := args[2]

	if err != nil {
		log.Fatal(err)
	}
	start := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go ping(addr, text, &wg)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("%s\n", elapsed)
}
