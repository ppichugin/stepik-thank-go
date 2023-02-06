package main

import (
	"strings"
)

func main() {

}

// MatchContains Библиотечная
func MatchContains(pattern string, src string) bool {
	return strings.Contains(src, pattern)
}

// MatchContainsCustom Самописная
func MatchContainsCustom(pattern string, src string) bool {
	if pattern == "" {
		return true
	}
	if len(pattern) > len(src) {
		return false
	}
	pat_len := len(pattern)
	for idx := 0; idx < len(src)-pat_len+1; idx++ {
		if src[idx:idx+pat_len] == pattern {
			return true
		}
	}
	return false
}
