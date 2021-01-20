package p011

import (
	"testing"
)

func TestSolve(t *testing.T) {
	want := "70600674"
	got := Solve()
	if got != want {
		t.Errorf("Got %v, want %v", got, want)
	}
}
