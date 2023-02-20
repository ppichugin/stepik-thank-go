package main

import "time"

// начало решения

func isLeapYear(year int) bool {
	return time.Date(year, 12, 31, 23, 59, 0, 0, time.Local).YearDay() == 366
	// classic solution:
	// return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// конец решения
