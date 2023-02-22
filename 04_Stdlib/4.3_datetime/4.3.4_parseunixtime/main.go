package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// начало решения

// asLegacyDate преобразует время в легаси-дату
func asLegacyDate(t time.Time) string {
	nanos := t.Unix()
	millis := t.UnixNano()
	strNanos := strconv.FormatInt(nanos, 10)
	strMil := (strconv.FormatInt(millis, 10))[len(strNanos):]
	strMil = strings.TrimLeft(strMil, "0")
	strMil = strings.TrimRight(strMil, "0")
	strMilDot := "." + strMil
	if len(strMilDot) == 1 {
		strMilDot += "0"
	}
	return strNanos + strMilDot
}

var ErrInvalidDateFormat = errors.New("invalid unix date format")

// parseLegacyDate преобразует легаси-дату во время.
// Возвращает ошибку, если легаси-дата некорректная.
func parseLegacyDate(d string) (time.Time, error) {
	reg := regexp.MustCompile(`^(\d+)(\.)(\d{1,9})?$`)

	parsed := reg.FindStringSubmatch(d)
	if len(parsed) < 4 || len(parsed[3]) == 0 {
		return time.Time{}, ErrInvalidDateFormat
	}

	part1, err := strconv.ParseInt(parsed[1], 10, 64)
	if err != nil {
		return time.Time{}, ErrInvalidDateFormat
	}

	s := parsed[3] + strings.Repeat("0", 9-len(parsed[3]))

	part2, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return time.Time{}, ErrInvalidDateFormat
	}

	return time.Unix(part1, part2), nil
}

// конец решения

func main() {
	date, err := parseLegacyDate("3600.123456789")
	fmt.Println(date, err)

	legacyDate := asLegacyDate(time.Date(1970, 1, 1, 1, 0, 0, 123456789, time.UTC))
	fmt.Println("legacy date: ", legacyDate)

	legacyDate = asLegacyDate(time.Date(1970, 1, 1, 1, 0, 0, 0, time.UTC))
	fmt.Println("legacy date: ", legacyDate)

	legacyDate = asLegacyDate(time.Date(2022, 5, 24, 14, 45, 22, 951205000, time.UTC))
	fmt.Println("legacy date: ", legacyDate)

	date, err = parseLegacyDate("1653403522.951205000")
	fmt.Println(date, err)
}
