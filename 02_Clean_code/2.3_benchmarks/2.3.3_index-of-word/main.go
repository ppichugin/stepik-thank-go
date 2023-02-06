package main

import (
	"fmt"
	"strings"
)

func main() {
	words := MakeWords("in a coat of gold or a coat of red")
	index := words.Index("a")
	fmt.Println(index)

}

// Words represents a structure of words
type Words struct {
	m map[string]int
}

// MakeWords counts the first entrance of the word
func MakeWords(s string) Words {
	m := make(map[string]int, len(s))
	for i, word := range strings.Fields(s) {
		_, ok := m[word]
		if ok {
			continue
		}
		m[word] = i
	}
	return Words{m}
}

// Index returns index of the word
func (w *Words) Index(word string) int {
	index, ok := w.m[word]
	if ok {
		return index
	}
	return -1
}
