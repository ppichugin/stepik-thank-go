package main

import (
	"sort"
	"strconv"
	"strings"
)

// начало решения

// prettify возвращает отформатированное
// строковое представление карты
func prettify(m map[string]int) string {
	if len(m) == 0 {
		return "{}"
	}

	builder := strings.Builder{}
	builder.Grow(len(m))
	isFlat := false

	sorted := sortMapToSlice(m)
	if len(sorted) <= 1 {
		isFlat = true
	}

	builder.WriteByte('{')
	for i, s := range sorted {
		value := m[s]

		// first enter
		if i == 0 && !isFlat {
			builder.WriteByte('\n')
		}

		if isFlat {
			builder.WriteString(" ")
		} else {
			builder.WriteString("    ")
		}

		builder.WriteString(s)
		builder.WriteString(": ")
		builder.WriteString(strconv.Itoa(value))

		if len(sorted) > 1 {
			builder.WriteByte(',')
		}
		if len(sorted) > 1 && i != len(sorted)-1 {
			builder.WriteByte('\n')
		}

		// last enter
		if i == len(sorted)-1 && !isFlat {
			builder.WriteByte('\n')
		}
	}
	if isFlat {
		builder.WriteString(" ")
	}
	builder.WriteByte('}')

	return builder.String()
}

func sortMapToSlice(m map[string]int) []string {
	res := make([]string, 0, len(m))
	for s := range m {
		res = append(res, s)
	}
	sort.Strings(res)

	return res
}

// конец решения
