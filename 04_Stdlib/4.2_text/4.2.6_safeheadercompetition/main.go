package main

import (
	"bytes"
	"fmt"
	"strings"
)

// начало решения

func slugify(src string) string {
	bldrPhrase := strings.Builder{}
	bldrPhrase.Grow(len(src))

	var buf bytes.Buffer

	var isLastHyphen bool
	if src[len(src)-1] == '-' {
		isLastHyphen = true
	}

	for i, ch := range []byte(src) {
		if ch == ' ' || ch == 39 || ch == '_' ||
			(ch == '/' && (i != len(src)-1 && src[i+1] != ' ')) ||
			ch == 226 ||
			(ch == '.' && (i != len(src)-1 && src[i+1] != ' ')) ||
			(ch == '.' && (i == len(src)-1)) ||
			(ch == '#' && (i == len(src)-1)) {
			bldrPhrase.WriteByte('-')
			continue
		}
		if isAllowedChar(ch) {
			if buf.Len() > 0 {
				if i < len(src) {
					bldrPhrase.WriteByte('-')
				} else {
					bldrPhrase.Write(buf.Bytes())
					buf.Reset()
				}
			}
			if ch >= 'A' && ch <= 'Z' {
				ch += 'a' - 'A'
			}
			bldrPhrase.WriteByte(ch)
		}
	}
	if bldrPhrase.String()[bldrPhrase.Len()-1:bldrPhrase.Len()] == "-" && !isLastHyphen {
		return bldrPhrase.String()[:bldrPhrase.Len()-1]
	}

	return bldrPhrase.String()
}

func isAllowedChar(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') || (ch == '-')
}

// конец решения

func main() {
	var phrase = "A 100x Investment (2019)"
	slug := slugify(phrase)
	fmt.Println(slug)
	// a-100x-investment-2019

	phrase = "Go Is Awesome!"
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

	phrase = `Go at Google I/O`
	fmt.Println(slugify(phrase))
	// go-at-google-i-o

	phrase = `We haven’t killed 90% of all plankton`
	fmt.Println(slugify(phrase))
	// we-haven-t-killed-90-of-all-plankton

	phrase = `Go 1.18 is released!`
	fmt.Println(slugify(phrase))
	// go-1-18-is-released

	phrase = `Y1nbl5 Wn9vc/ Nna0dq L3np- Kgmmhto#`
	fmt.Println(slugify(phrase))
	// y1nbl5-wn9vc-nna0dq-l3np--kgmmhto
}
