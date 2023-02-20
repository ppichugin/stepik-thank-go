package main

import (
	"log"
	"testing"
	"time"
)

func Test(t *testing.T) {
	t1 := MakeTimeOfDay(17, 45, 22, time.UTC)
	t2 := MakeTimeOfDay(20, 3, 4, time.UTC)

	if t1.Equal(t2) {
		t.Errorf("%v should not be equal to %v", t1, t2)
	}

	before, _ := t1.Before(t2)
	if !before {
		t.Errorf("%v should be before %v", t1, t2)
	}

	after, _ := t1.After(t2)
	if after {
		t.Errorf("%v should NOT be after %v", t1, t2)
	}

	/*
		08:59:59 UTC to 09:01:01 UTC
	*/
	t1 = MakeTimeOfDay(8, 59, 59, time.UTC)
	t2 = MakeTimeOfDay(9, 1, 1, time.UTC)

	after, err := t1.After(t2)
	if err != nil {
		log.Println(err)
		t.Errorf("After(): unexpected error comparing")
	}
	if after {
		t.Errorf("%v should NOT be equal to %v", t1, t2)
	}
}
