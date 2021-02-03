package p015

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	want := "137846528820"
	got := Solve()
	if got != want {
		t.Errorf("Got %v, want %v", got, want)
	}
}

func TestBoxPaths(t *testing.T) {
	var tests = []struct {
		x, y, ans int
	}{
		{1, 1, 2},
		{2, 2, 6},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Box %v x %v", tt.x, tt.y)
		t.Run(testname, func(t *testing.T) {
			ans := boxPaths(tt.x, tt.y)
			if ans != tt.ans {
				t.Errorf("Got %v, want %v", ans, tt.ans)
			}
		})
	}
}
