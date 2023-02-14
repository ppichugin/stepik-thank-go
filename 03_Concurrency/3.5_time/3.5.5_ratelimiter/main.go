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
	finished := make(chan struct{}, 1)

	rateLimit := time.Second / time.Duration(limit)
	fmt.Println("rate limit: ", rateLimit)
	throttler := time.Tick(rateLimit)

	starter := func() {
		for i := 0; i < limit; i++ {
			select {
			case <-canceled:
				return
			case <-throttler:
				select {
				case f := <-queue:
					go f()
				case <-canceled:
					return
				}
			}
		}
		finished <- struct{}{}
	}

	go func() {
		defer close(queue)
		for {
			select {
			case <-canceled:
				return
			case <-finished:
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
		}
	}

	finished <- struct{}{}

	return handle, cancel
}

// конец решения

func main() {
	var dots int
	work := func() {
		dots++
		fmt.Print(".")
		//time.Sleep(500 * time.Millisecond)
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
