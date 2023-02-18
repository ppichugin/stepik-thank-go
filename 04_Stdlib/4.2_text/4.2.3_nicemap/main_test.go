package main

import "testing"

func Test(t *testing.T) {
	testcases := []struct {
		name string
		m    map[string]int
		want string
	}{
		{
			name: "multi line map",
			m:    map[string]int{"one": 1, "two": 2, "three": 3},
			want: "{\n    one: 1,\n    three: 3,\n    two: 2,\n}",
		},
		{
			name: "one line map",
			m:    map[string]int{"answer": 42},
			want: "{ answer: 42 }",
		},
		{
			name: "empty map",
			m:    map[string]int{},
			want: "{}",
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			got := prettify(testcase.m)
			want := testcase.want
			if got != want {
				t.Errorf("%v\ngot:\n%v\n\nwant:\n%v", testcase.m, got, want)
			}
		})
	}
}
