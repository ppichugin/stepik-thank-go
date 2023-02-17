package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// начало решения

type Total struct {
	variable int32
}

func (t *Total) Increment() {
	atomic.AddInt32(&t.variable, 1)
}

func (t *Total) Value() int {
	return int(atomic.LoadInt32(&t.variable))
}

// конец решения

func main() {
	var wg sync.WaitGroup

	var total Total

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				total.Increment()
			}
		}()
	}

	wg.Wait()
	fmt.Println("total", total.Value())
}
