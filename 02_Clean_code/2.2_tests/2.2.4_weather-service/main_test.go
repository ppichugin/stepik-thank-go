package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// WeatherService предсказывает погоду.
type WeatherService struct{}

type WeatherSvc interface {
	Forecast() int
}

type WeatherServiceStub struct {
	Temperature int
}

// Forecast сообщает ожидаемую дневную температуру на завтра.
func (ws *WeatherService) Forecast() int {
	rand.Seed(time.Now().Unix())
	value := rand.Intn(31)
	sign := rand.Intn(2)
	if sign == 1 {
		value = -value
	}
	return value
}

func (wss *WeatherServiceStub) Forecast() int {
	return wss.Temperature
}

// Weather выдает текстовый прогноз погоды.
type Weather struct {
	service WeatherSvc
}

// Forecast сообщает текстовый прогноз погоды на завтра.
func (w Weather) Forecast() string {
	deg := w.service.Forecast()
	switch {
	case deg < 10:
		return "холодно"
	case deg >= 10 && deg < 15:
		return "прохладно"
	case deg >= 15 && deg < 20:
		return "идеально"
	case deg >= 20:
		return "жарко"
	}
	return "инопланетно"
}

type testCase struct {
	deg  int
	want string
}

var tests []testCase = []testCase{
	{-10, "холодно"},
	{0, "холодно"},
	{5, "холодно"},
	{10, "прохладно"},
	{15, "идеально"},
	{20, "жарко"},
}

func TestForecast(t *testing.T) {
	service := &WeatherServiceStub{}
	weather := Weather{service}
	for _, test := range tests {
		name := fmt.Sprintf("%v", test.deg)
		t.Run(name, func(t *testing.T) {
			service.Temperature = test.deg
			got := weather.Forecast()
			if got != test.want {
				t.Errorf("%s: got %s, want %s", name, got, test.want)
			}
		})
	}
}
