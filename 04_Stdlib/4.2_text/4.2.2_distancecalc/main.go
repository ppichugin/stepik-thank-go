package main

import (
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

// начало решения

// calcDistance возвращает общую длину маршрута в метрах
func calcDistance(directions []string) int {
	var distance int

loop:
	for _, direction := range directions {
		for _, word := range strings.Fields(direction) {
			char, _ := utf8.DecodeRuneInString(word)
			if !unicode.IsDigit(char) {
				continue
			}
			if strings.HasSuffix(word, "km") {
				value, err := strconv.ParseFloat(strings.TrimSuffix(word, "km"), 64)
				if err == nil {
					distance += int(value * 1000)
				}
				//fmt.Printf("direction: %s, km: %f\n", direction, value)
				continue loop
			}
			if strings.HasSuffix(word, "m") {
				value, err := strconv.Atoi(strings.TrimSuffix(word, "m"))
				if err == nil {
					distance += value
				}
				//fmt.Printf("direction: %s, meters: %d\n", direction, value)
				continue loop
			}
		}
	}

	return distance
}

// конец решения

func main() {

}
