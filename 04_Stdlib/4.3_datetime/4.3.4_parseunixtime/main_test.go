package main

import (
	"testing"
	"time"
)

func Test_asLegacyDate(t *testing.T) {
	samples := map[time.Time]string{
		time.Date(1970, 1, 1, 1, 0, 0, 123456789, time.UTC): "3600.123456789",
		time.Date(1970, 1, 1, 1, 0, 0, 951, time.UTC):       "3600.951",
		time.Date(2022, 5, 24, 14, 45, 22, 951, time.UTC):   "1653403522.951",
		time.Date(1970, 1, 1, 1, 0, 0, 123, time.UTC):       "3600.123",
		time.Date(1970, 1, 1, 1, 0, 0, 0, time.UTC):         "3600.0",
		time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC):         "0.0",
	}
	for src, want := range samples {
		got := asLegacyDate(src)
		if got != want {
			t.Fatalf("%v: got %v, want %v", src, got, want)
		}
	}
}

func Test_parseLegacyDate(t *testing.T) {
	samples := map[string]time.Time{
		"3600.123456789":       time.Date(1970, 1, 1, 1, 0, 0, 123456789, time.UTC),
		"3600.0":               time.Date(1970, 1, 1, 1, 0, 0, 0, time.UTC),
		"0.0":                  time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
		"1.123456789":          time.Date(1970, 1, 1, 0, 0, 1, 123456789, time.UTC),
		"1653403522.951":       time.Date(2022, 5, 24, 14, 45, 22, 951000000, time.UTC),
		"1653403522.951205000": time.Date(2022, 5, 24, 14, 45, 22, 951205000, time.UTC),
	}
	for src, want := range samples {
		got, err := parseLegacyDate(src)
		if err != nil {
			t.Error(err)
			t.Fatalf("%v: unexpected error", src)
		}
		if !got.Equal(want) {
			t.Fatalf("%v: got %v, want %v", src, got, want)
		}
	}
}
