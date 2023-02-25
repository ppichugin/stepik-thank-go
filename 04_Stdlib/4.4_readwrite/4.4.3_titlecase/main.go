package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Print(" ")
		w := scanner.Text()
		word := []rune(w)
		upper := strings.ToUpper(string(word[0]))
		lower := strings.ToLower(string(word[1:]))
		fmt.Print(upper + lower)
	}
	fmt.Println()
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
