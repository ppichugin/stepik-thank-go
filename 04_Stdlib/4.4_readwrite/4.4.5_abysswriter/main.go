package main

import (
	"fmt"
	"io"
	"strings"
)

// начало решения

// AbyssWriter пишет данные в никуда,
// но при этом считает количество записанных байт
type AbyssWriter struct {
	written int
}

func (w *AbyssWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	w.written += n
	return n, nil
}

// Total возвращает общее количество записанных байт
func (w *AbyssWriter) Total() int {
	return w.written
}

// NewAbyssWriter создает новый AbyssWriter
func NewAbyssWriter() *AbyssWriter {
	aw := AbyssWriter{}
	return &aw
}

// конец решения

func main() {
	r := strings.NewReader("go is awesome")
	w := NewAbyssWriter()
	written, err := io.Copy(w, r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("written %d bytes\n", written)
	fmt.Println(written == int64(w.Total()))
}
