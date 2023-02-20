package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

func main() {
	t1 := MakeTimeOfDay(17, 45, 22, time.UTC)
	t2 := MakeTimeOfDay(20, 3, 4, time.UTC)

	if t1.Equal(t2) {
		log.Printf("%v should not be equal to %v\n", t1, t2)
	}

	before, _ := t1.Before(t2)
	if !before {
		log.Printf("%v should be before %v\n", t1, t2)
	}
}

// начало решения

// TimeOfDay описывает время в пределах одного дня
type TimeOfDay struct {
	hour, min, sec int
	loc            *time.Location
}

// Hour возвращает часы в пределах дня
func (t TimeOfDay) Hour() int {
	return t.hour
}

// Minute возвращает минуты в пределах часа
func (t TimeOfDay) Minute() int {
	return t.min
}

// Second возвращает секунды в пределах минуты
func (t TimeOfDay) Second() int {
	return t.sec
}

// String возвращает строковое представление времени
// в формате чч:мм:сс TZ (например, 12:34:56 UTC)
func (t TimeOfDay) String() string {
	return fmt.Sprintf("%02d:%02d:%02d %s", t.hour, t.min, t.sec, t.loc.String())
}

// Equal сравнивает одно время с другим.
// Если у t и other разные локации - возвращает false.
func (t TimeOfDay) Equal(other TimeOfDay) bool {
	if !t.isComparable(other) {
		return false
	}
	cur, oth := convToTime(t, other)
	return cur.Equal(oth)
}

var ErrTimeLocation = errors.New("different locations")

// Before возвращает true, если время t предшествует other.
// Если у t и other разные локации - возвращает ошибку.
func (t TimeOfDay) Before(other TimeOfDay) (bool, error) {
	if !t.isComparable(other) {
		return false, ErrTimeLocation
	}
	cur, oth := convToTime(t, other)
	if cur.Before(oth) {
		return true, nil
	}
	return false, nil
}

func (t TimeOfDay) isComparable(other TimeOfDay) bool {
	return t.loc.String() == other.loc.String()
}

func convToTime(c TimeOfDay, o TimeOfDay) (time.Time, time.Time) {
	current := time.Date(2023, 01, 01, c.hour, c.min, c.sec, 0, c.loc)
	other := time.Date(2023, 01, 01, o.hour, o.min, o.sec, 0, o.loc)
	return current, other
}

// After возвращает true, если время t идет после other.
// Если у t и other разные локации - возвращает ошибку.
func (t TimeOfDay) After(other TimeOfDay) (bool, error) {
	if !t.isComparable(other) {
		return false, ErrTimeLocation
	}
	cur, oth := convToTime(t, other)
	if cur.After(oth) {
		return true, nil
	}
	return false, nil
}

// MakeTimeOfDay создает время в пределах дня
func MakeTimeOfDay(hour, min, sec int, loc *time.Location) TimeOfDay {
	return TimeOfDay{
		hour: hour,
		min:  min,
		sec:  sec,
		loc:  loc,
	}
}

// конец решения
