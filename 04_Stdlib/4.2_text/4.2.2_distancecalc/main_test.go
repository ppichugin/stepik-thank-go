package main

import "testing"

func Test(t *testing.T) {
	directions := []string{
		"100m to intersection",
		"turn right",
		"straight 300m",
		"enter motorway",
		"straight 5km",
		"exit motorway",
		"500m straight",
		"turn sharp left",
		"continue 100m to destination",
		"tldr 1.6km",
	}
	//const want = 6000
	const want = 7600
	got := calcDistance(directions)
	if got != want {
		t.Errorf("%v: got %v, want %v", directions, got, want)
	}
}
