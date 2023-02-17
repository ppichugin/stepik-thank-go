package main

import (
	"fmt"
	"time"
)

func delay(duration time.Duration, fn func()) func() {
	alive := make(chan struct{})
	close(alive)

	go func() {
		time.Sleep(duration)
		select {
		case <-alive: // data race
			fn()
		default:
		}
	}()

	cancel := func() {
		alive = nil // data race
	}
	return cancel
}

func main() {
	work := func() {
		fmt.Println("work done")
	}

	cancel := delay(50*time.Millisecond, work)
	defer cancel()
	time.Sleep(100 * time.Millisecond)
}
