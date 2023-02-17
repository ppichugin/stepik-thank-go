package main

import (
	"fmt"
	"strings"
	"unicode"
)

// начало решения

// slugify возвращает "безопасный" вариант заголовка:
// только латиница, цифры и дефис
func slugify(src string) string {
	bldrPhrase := strings.Builder{}
	bldrPhrase.Grow(len(src))

	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c) && !isAllowedChar(c)
	}

	words := strings.FieldsFunc(src, f)

	for i, word := range words {
		bldrWord := strings.Builder{}
		bldrWord.Grow(len(word))

		for _, r := range word {
			if isAllowedChar(r) {
				switch {
				case unicode.IsLetter(r):
					bldrWord.WriteRune(unicode.ToLower(r))
				case unicode.IsDigit(r):
					bldrWord.WriteRune(r)
				case r == '-':
					bldrWord.WriteRune(r)
				}
			}
		}

		if bldrWord.String() != "" {
			if i == 0 {
				bldrPhrase.WriteString(bldrWord.String())
				continue
			}
			if i < len(words) {
				bldrPhrase.WriteByte('-')
				bldrPhrase.WriteString(bldrWord.String())
			}
		}
		bldrWord.Reset()
	}
	return bldrPhrase.String()
}

func isAllowedChar(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') || (ch == '-')
}

// конец решения

func main() {
	phrase := "Go Is Awesome!"
	fmt.Println(slugify(phrase))
	// go-is-awesome

	phrase = "Tabs are all we've got"
	fmt.Println(slugify(phrase))
	// tabs-are-all-we-ve-got

	phrase = "Go-Is-Awesome"
	fmt.Println(slugify(phrase))
	// go-is-awesome"

	phrase = "Go_Is_Awesome"
	fmt.Println(slugify(phrase))
	// go-is-awesome"

	phrase = `Go Talks: "Cuddle: an App Engine Demo":`
	fmt.Println(slugify(phrase))
	// go-talks-cuddle-an-app-engine-demo

	phrase = `Hello, 中国!`
	fmt.Println(slugify(phrase))
	// hello

	phrase = `Go - Is - Awesome`
	fmt.Println(slugify(phrase))
	// Go - Is - Awesome
}
