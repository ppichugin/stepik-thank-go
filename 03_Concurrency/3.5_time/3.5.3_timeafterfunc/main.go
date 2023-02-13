package main

import (
	"fmt"
	"math/rand"
	"time"
)

// начало решения

func delay(dur time.Duration, fn func()) func() {
	cancelChan := make(chan struct{})
	timer := time.NewTimer(dur)

	go func() {
		select {
		case <-timer.C:
			fn()
		case <-cancelChan:
			return
		}
	}()
	stop := func() {
		if timer.Stop() {
			cancelChan <- struct{}{}
			close(cancelChan)
			return
		}
	}
	return stop
}

// конец решения

func main() {
	rand.Seed(time.Now().Unix())

	work := func() {
		fmt.Println("work done")
	}

	cancel := delay(100*time.Millisecond, work)

	time.Sleep(10 * time.Millisecond)
	if rand.Float32() < 0.5 {
		cancel()
		fmt.Println("delayed function canceled")
	}
	time.Sleep(100 * time.Millisecond)
}
