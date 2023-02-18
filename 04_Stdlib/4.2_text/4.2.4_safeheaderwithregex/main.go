package main

import (
	"fmt"
	"regexp"
	"strings"
)

// начало решения

// slugify возвращает "безопасный" вариант заголовка:
// только латиница, цифры и дефис
func slugify(src string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9-]+`)
	processed := re.ReplaceAllString(src, "-")
	processed = strings.ToLower(processed)
	return strings.Trim(processed, "-")
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
