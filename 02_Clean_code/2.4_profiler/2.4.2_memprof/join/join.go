package join

import (
	"sort"
	"strings"
)

// JoinWords combines words from two strings, removes duplicates
// and returns resulting words in sorted order.
func JoinWords(first, second string) []string {
	words1 := split(first)
	words2 := split(second)
	words := join(words1, words2)
	words = lower(words)
	words = sorted(words)
	words = uniq(words)
	return words
}

// split splits string into words.
func split(str string) []string {
	return strings.Fields(str)
}

// join joines two word slices into single slice.
func join(words1, words2 []string) []string {
	words := make([]string, len(words1)+len(words2))
	idx := 0
	for _, word := range words1 {
		words[idx] = word
		idx++
	}
	for _, word := range words2 {
		words[idx] = word
		idx++
	}
	return words
}

// lower converts all words to lower case.
func lower(words []string) []string {
	for idx, word := range words {
		words[idx] = strings.ToLower(word)
	}
	return words
}

// sorted sorts all words alphabetically.
func sorted(words []string) []string {
	sort.Strings(words)
	return words
}

// uniq removes duplicate words.
func uniq(words []string) []string {
	uniq := []string{}
	current := ""
	for _, word := range words {
		if word == current {
			continue
		}
		if current != "" {
			uniq = append(uniq, current)
			current = ""
		}
		current = word
	}
	if current != "" {
		uniq = append(uniq, current)
	}
	return uniq
}
