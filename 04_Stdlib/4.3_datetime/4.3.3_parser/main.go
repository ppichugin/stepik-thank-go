package main

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"
)

// начало решения

// Task описывает задачу, выполненную в определенный день
type Task struct {
	Date  time.Time
	Dur   time.Duration
	Title string
}

// ParsePage разбирает страницу журнала
// и возвращает задачи, выполненные за день
func ParsePage(src string) ([]Task, error) {
	lines := strings.Split(src, "\n")
	date, err := parseDate(lines[0])
	if err != nil {
		return nil, err
	}
	tasks, err := parseTasks(date, lines[1:])
	if err != nil {
		return nil, err
	}

	sortTasks(tasks)
	return tasks, nil
}

var ErrTimeIncorrect = errors.New("incorrect time format")
var ErrListTasks = errors.New("incorrect tasks' list")

// parseDate разбирает дату в формате дд.мм.гггг
func parseDate(src string) (time.Time, error) {
	parsed, err := time.Parse("02.01.2006", src)
	if err != nil {
		return time.Time{}, ErrTimeIncorrect
	}
	return parsed, nil
}

// parseTasks разбирает задачи из записей журнала
func parseTasks(date time.Time, lines []string) ([]Task, error) {
	m := make(map[string]Task, len(lines))
	reg := regexp.MustCompile("(\\d+:\\d+) - (\\d+:\\d+) (.+)")
	for _, s := range lines {
		split := reg.FindStringSubmatch(s)
		if len(split) < 4 {
			return nil, ErrListTasks
		}

		start, err := time.Parse("15:04", split[1])
		if err != nil {
			return nil, err
		}

		end, err := time.Parse("15:04", split[2])
		if err != nil {
			return nil, err
		}

		if start.After(end) || start.Equal(end) {
			return nil, ErrListTasks
		}

		task := Task{
			Date:  date,
			Dur:   end.Sub(start),
			Title: split[3],
		}

		var dur time.Duration
		if t, ok := m[task.Title]; ok {
			dur = t.Dur + task.Dur
		} else {
			dur = task.Dur
		}
		task.Dur = dur
		m[task.Title] = task
	}
	return mapToSlice(m), nil
}

func mapToSlice(m map[string]Task) []Task {
	res := make([]Task, 0, len(m))
	for _, task := range m {
		res = append(res, task)
	}
	return res
}

// sortTasks упорядочивает задачи по убыванию длительности
func sortTasks(tasks []Task) {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Dur > tasks[j].Dur
	})
}

// конец решения
// ::footer

func main() {
	page := `15.04.2022
8:00 - 8:30 Завтрак
8:30 - 9:30 Оглаживание кота
9:30 - 10:00 Интернеты
10:00 - 14:00 Напряженная работа
14:00 - 14:45 Обед
14:45 - 15:00 Оглаживание кота
15:00 - 19:00 Напряженная работа
19:00 - 19:30 Интернеты
19:30 - 22:30 Безудержное веселье
22:30 - 23:00 Оглаживание кота`

	entries, err := ParsePage(page)
	if err != nil {
		panic(err)
	}
	fmt.Println("Мои достижения за", entries[0].Date.Format("2006-01-02"))
	for _, entry := range entries {
		fmt.Printf("- %v: %v\n", entry.Title, entry.Dur)
	}

	// ожидаемый результат
	/*
		Мои достижения за 2022-04-15
		- Напряженная работа: 8h0m0s
		- Безудержное веселье: 3h0m0s
		- Оглаживание кота: 1h45m0s
		- Интернеты: 1h0m0s
		- Обед: 45m0s
		- Завтрак: 30m0s
	*/
}
