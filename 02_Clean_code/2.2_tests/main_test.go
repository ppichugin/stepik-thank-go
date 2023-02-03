package main

import (
	"testing"
)

func TestSumZero(t *testing.T) {
	if Sum() != 0 {
		t.Errorf("Expected Sum() == 0")
	}
}

func TestSumOne(t *testing.T) {
	if Sum(1) != 1 {
		t.Errorf("Expected Sum(1) == 1")
	}
}

func TestSumPair(t *testing.T) {
	if Sum(1, 2) != 3 {
		t.Errorf("Expected Sum(1, 2) == 3")
	}
}

func TestSumMany(t *testing.T) {
	if Sum(1, 2, 3, 4, 5) != 15 {
		t.Errorf("Expected Sum(1, 2, 3, 4, 5) == 15")
	}
}
