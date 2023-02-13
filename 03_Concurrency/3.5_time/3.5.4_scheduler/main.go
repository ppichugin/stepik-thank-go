package main

import (
	"fmt"
	"time"
)

// начало решения

func schedule(dur time.Duration, fn func()) func() {
	ticker := time.NewTicker(dur)
	canceled := make(chan struct{})

	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-canceled:
				return
			case <-ticker.C:
				fn()
			}
		}
	}()

	return func() {
		select {
		case <-canceled:
			return
		default:
			ticker.Stop()
			close(canceled)
		}
	}
}

// конец решения

func main() {
	work := func() {
		at := time.Now()
		fmt.Printf("%s: work done\n", at.Format("15:04:05.000"))
	}

	cancel := schedule(50*time.Millisecond, work)
	defer cancel()

	// хватит на 5 тиков
	time.Sleep(260 * time.Millisecond)
}
