package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
)

// начало решения

// RandomReader создает читателя, который возвращает случайные байты,
// но не более max штук
func RandomReader(max int) io.Reader {
	return &randomReader{max: max}
}

type randomReader struct {
	max int
}

func (r *randomReader) Read(p []byte) (int, error) {
	if r.max <= 0 {
		return 0, io.EOF
	}
	if len(p) > r.max {
		p = p[:r.max]
	}
	n, err := rand.Read(p)
	r.max -= n
	if err != nil {
		return n, err
	}
	return n, nil
}

// конец решения

func main() {
	rand.Seed(0)

	rnd := RandomReader(5)
	rd := bufio.NewReader(rnd)
	for {
		b, err := rd.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d ", b)
	}
	fmt.Println()
	// 1 148 253 194 250
}
