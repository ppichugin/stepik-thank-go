package main

import (
	"errors"
	"fmt"
	"time"
)

var ErrCanceled error = errors.New("canceled")

// начало решения

func withRateLimit(limit int, fn func()) (handle func() error, cancel func()) {
	queue := make(chan func())
	canceled := make(chan struct{})

	rateLimit := time.Second / time.Duration(limit)
	fmt.Println("rate limit: ", rateLimit)
	ticker := time.NewTicker(1 * time.Second)

	starter := func() {
		for i := 0; i < limit; i++ {
			select {
			case <-canceled:
				return
			case f := <-queue:
				f()
			}
		}
	}

	go func() {
		defer ticker.Stop()
		defer close(queue)
		for {
			select {
			case <-canceled:
				return
			case <-ticker.C:
				go starter()
			}
		}
	}()

	handle = func() error {
		select {
		case <-canceled:
			return ErrCanceled
		case queue <- fn:
		}
		return nil
	}

	cancel = func() {
		select {
		case <-canceled:
			return
		default:
			close(canceled)
			//close(queue)
		}
	}

	return handle, cancel
}

// конец решения

func main() {
	var dots int
	work := func() {
		dots++
		fmt.Print(".")
		//time.Sleep(200 * time.Millisecond)
	}

	handle, cancel := withRateLimit(10, work)
	defer cancel()

	start := time.Now()
	const n = 2
	for i := 0; i < n; i++ {
		handle()
	}

	fmt.Println()
	fmt.Printf("%d queries took %v\n", n, time.Since(start))
	fmt.Println("dots printed:", dots)
}
