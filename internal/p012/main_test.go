package p012

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	want := "76576500"
	got := Solve()
	if got != want {
		t.Errorf("Got %v, want %v", got, want)
	}
}

func TestNumOfDivisors(t *testing.T) {
	var tests = []struct {
		in  int
		out int
	}{
		{1, 1},
		{3, 2},
		{6, 4},
		{10, 4},
		{15, 4},
		{21, 4},
		{28, 6},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("number of divisors of %v", tt.in)
		t.Run(testname, func(t *testing.T) {
			ans := numOfDivisors(tt.in)
			if ans != tt.out {
				t.Errorf("Got %v, want %v", ans, tt.out)
			}
		})
	}
}
