package main

import (
	"fmt"
	"time"
)

func delay(duration time.Duration, fn func()) func() {
	canceled := false

	go func() {
		time.Sleep(duration)
		if !canceled { // data race
			fn()
		}
	}()

	cancel := func() {
		canceled = true // data race
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
