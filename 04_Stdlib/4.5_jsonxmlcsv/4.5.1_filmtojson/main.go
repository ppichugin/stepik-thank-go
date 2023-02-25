package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

// начало решения

// Duration описывает продолжительность фильма
type Duration time.Duration

func (d Duration) MarshalJSON() ([]byte, error) {
	minutes := time.Duration(d).Minutes()
	hours := int(minutes / 60)
	reminderMinutes := int(minutes) % 60
	var hStr, mStr string
	b := make([]byte, 0, 8)
	b = append(b, '"')
	if hours > 0 {
		hStr = fmt.Sprintf("%dh", hours)
		b = append(b, []byte(hStr)...)
	}
	if reminderMinutes > 0 {
		mStr = fmt.Sprintf("%dm", reminderMinutes)
		b = append(b, []byte(mStr)...)
	}
	b = append(b, '"')
	return b, nil
}

// Rating описывает рейтинг фильма
type Rating int

func (r Rating) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer

	if r < 0 {
		return []byte{}, errors.New("rating must be greater than 0")
	}

	buf.WriteByte('"')
	stars := int(r)
	for i := 0; i < stars; i++ {
		buf.WriteString("★")
	}
	for i := 0; i < 5-stars; i++ {
		buf.WriteString("☆")
	}
	buf.WriteByte('"')

	return buf.Bytes(), nil
}

// Movie описывает фильм
type Movie struct {
	Title    string   `json:"Title"`
	Year     int      `json:"Year"`
	Director string   `json:"Director"`
	Genres   []string `json:"Genres"`
	Duration Duration `json:"Duration"`
	Rating   Rating   `json:"Rating"`
}

// MarshalMovies кодирует фильмы в JSON.
//   - если indent = 0 - использует json.Marshal
//   - если indent > 0 - использует json.MarshalIndent
//     с отступом в указанное количество пробелов.
func MarshalMovies(indent int, movies ...Movie) (string, error) {
	var b []byte
	var err error

	if indent == 0 {
		b, err = json.Marshal(movies)
	} else {
		b, err = json.MarshalIndent(movies, "", strings.Repeat(" ", indent))
	}
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// конец решения

func main() {
	m1 := Movie{
		Title:    "Interstellar",
		Year:     2014,
		Director: "Christopher Nolan",
		Genres:   []string{"Adventure", "Drama", "Science Fiction"},
		Duration: Duration(2*time.Hour + 49*time.Minute),
		Rating:   5,
	}
	m2 := Movie{
		Title:    "Sully",
		Year:     2016,
		Director: "Clint Eastwood",
		Genres:   []string{"Drama", "History"},
		Duration: Duration(time.Hour + 36*time.Minute),
		Rating:   4,
	}

	b, err := MarshalMovies(4, m1, m2)
	fmt.Println(err)
	// nil
	fmt.Println(string(b))
	/*
		[
		    {
		        "Title": "Interstellar",
		        "Year": 2014,
		        "Director": "Christopher Nolan",
		        "Genres": [
		            "Adventure",
		            "Drama",
		            "Science Fiction"
		        ],
		        "Duration": "2h49m",
		        "Rating": "★★★★★"
		    },
		    {
		        "Title": "Sully",
		        "Year": 2016,
		        "Director": "Clint Eastwood",
		        "Genres": [
		            "Drama",
		            "History"
		        ],
		        "Duration": "1h36m",
		        "Rating": "★★★★☆"
		    }
		]
	*/
}
