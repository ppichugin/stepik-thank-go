package main

import (
	"fmt"
	"time"
)

// gather выполняет переданные функции одновременно
// и возвращает срез с результатами, когда они готовы
func gather(funcs []func() any) []any {
	// начало решения

	l := len(funcs)
	type result struct {
		id     int
		answer any
	}
	results := make(chan result, l)

	for i, f := range funcs {
		i, f := i, f
		go func(ch chan result) {
			a := f()
			ch <- result{
				id:     i,
				answer: a,
			}
		}(results)
	}

	// выполните все переданные функции,
	// соберите результаты в срез
	// и верните его

	answers := make([]any, l)
	for i := 0; i < l; i++ {
		a := <-results
		answers[a.id] = a.answer
	}

	return answers

	// конец решения
}

// squared возвращает функцию,
// которая считает квадрат n
func squared(n int) func() any {
	return func() any {
		time.Sleep(time.Duration(n) * 100 * time.Millisecond)
		return n * n
	}
}

func main() {
	funcs := []func() any{squared(2), squared(3), squared(4)}

	start := time.Now()
	nums := gather(funcs)
	elapsed := float64(time.Since(start)) / 1_000_000

	fmt.Println(nums)
	fmt.Printf("Took %.0f ms\n", elapsed)
}
