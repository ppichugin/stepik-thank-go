package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// начало решения

// генерит случайные слова из 5 букв
// с помощью randomWord(5)
func generate(cancel <-chan struct{}) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for {
			select {
			case <-cancel:
				return
			case out <- randomWord(5):
			}
		}
	}()
	return out
}

// выбирает слова, в которых не повторяются буквы,
// abcde - подходит
// abcda - не подходит
func takeUnique(cancel <-chan struct{}, in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for {
			select {
			case <-cancel:
				return
			case out <- checkWord(<-in):
			}
		}
	}()
	return out
}

func checkWord(str string) string {
	m := make(map[string]bool)
	for _, l := range str {
		if v := m[string(l)]; v {
			return ""
		}
		m[string(l)] = true
	}
	return str
}

type words struct {
	current  string
	reversed string
}

// переворачивает слова
// abcde -> edcba
func reverse(cancel <-chan struct{}, in <-chan string) <-chan words {
	out := make(chan words)
	go func() {
		defer close(out)
		for {
			select {
			case <-cancel:
				return
			case str := <-in:
				if str != "" {
					newWord := words{
						current:  str,
						reversed: reverseWord(str),
					}
					out <- newWord
				}
			}
		}
	}()
	return out
}

func reverseWord(word string) string {
	rns := []rune(word)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

// объединяет c1 и c2 в общий канал
func merge(cancel <-chan struct{}, c1, c2 <-chan words) <-chan words {
	out := make(chan words)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer close(out)

		go func(cancel <-chan struct{}) {
			defer wg.Done()
			for w := range c1 {
				select {
				case <-cancel:
					return
				case out <- w:
				}
			}
		}(cancel)

		go func(cancel <-chan struct{}) {
			defer wg.Done()
			for w := range c2 {
				select {
				case <-cancel:
					return
				case out <- w:
				}
			}
		}(cancel)

		wg.Wait()
	}()
	return out
}

// печатает первые n результатов
func print(cancel <-chan struct{}, in <-chan words, n int) {
	for i := 0; i < n; i++ {
		select {
		case <-cancel:
			return
		case word := <-in:
			fmt.Println(word.current, "->", word.reversed)
		}
	}
}

// конец решения

// генерит случайное слово из n букв
func randomWord(n int) string {
	const letters = "aeiourtnsl"
	chars := make([]byte, n)
	for i := range chars {
		chars[i] = letters[rand.Intn(len(letters))]
	}
	return string(chars)
}

func main() {
	cancel := make(chan struct{})
	defer close(cancel)

	c1 := generate(cancel)
	c2 := takeUnique(cancel, c1)
	c3_1 := reverse(cancel, c2)
	c3_2 := reverse(cancel, c2)
	c4 := merge(cancel, c3_1, c3_2)
	print(cancel, c4, 10)
}
