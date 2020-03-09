package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"math/rand"
	"time"
)

type Car struct {
	id  int
	dir bool
}

type Bridge struct {
	current *Car
}

func (b *Bridge) setCurrent(car *Car) {
	b.current = car
}

func (b *Bridge) hasCar() bool {
	return b.current != nil
}

func (b *Bridge) popCurrent() *Car {
	var aux = b.current
	b.current = nil
	return aux
}

var (
	ctx = context.TODO()
	sem = semaphore.NewWeighted(int64(1))
)

func consume(bridge *Bridge) {
	for {
		if err := sem.Acquire(ctx, 1); err != nil {
			fmt.Printf("Failed to acquire semaphore: %v", err)
			continue
		}
		if bridge.hasCar() {
			c := bridge.popCurrent()
			fmt.Printf("Consumed: %d, %t\n", c.id, c.dir)
		}
		sem.Release(1)
	}
}

func produce(bridge *Bridge, dir bool) {
	for {
		if err := sem.Acquire(ctx, 1); err != nil {
			fmt.Printf("Failed to acquire semaphore: %v\n", err)
			continue
		}
		if !bridge.hasCar() {
			c := &Car{id: rand.Intn(1000) + 1, dir: dir}
			bridge.setCurrent(c)
			fmt.Printf("Produced: %d, %t\n", c.id, c.dir)
		}
		sem.Release(1)
		time.Sleep(1 * time.Millisecond)
	}
}

// set arg in semaphore to 3 to act like there is no semaphore
func main() {
	bridge := &Bridge{}

	go consume(bridge)
	go produce(bridge, false)
	go produce(bridge, true)

	for {
	}
}
