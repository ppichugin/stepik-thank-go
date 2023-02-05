package main

import (
	"fmt"
	"testing"
)

func Sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func TestMain(m *testing.M) {
	fmt.Println("Testing Sum(n1, n2, ...,nk)...")
	fmt.Println("Finished testing")

	// should be uncommented for testing
	//m.Run()
}

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
	// Should be commented, unnecessary Skip for test
	//t.Skip()
}

func TestSumMany(t *testing.T) {
	if Sum(1, 2, 3, 4, 5) != 15 {
		t.Errorf("Expected Sum(1, 2, 3, 4, 5) == 15")
	}
}
