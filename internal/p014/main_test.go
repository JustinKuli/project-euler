package p014

import (
	"testing"
)

func TestSolve(t *testing.T) {
	want := "837799"
	got := Solve()
	if got != want {
		t.Errorf("Got %v, want %v", got, want)
	}
}
