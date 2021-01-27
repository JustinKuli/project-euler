package p012

import (
	"testing"
)

func TestSolve(t *testing.T) {
	want := "76576500"
	got := Solve()
	if got != want {
		t.Errorf("Got %v, want %v", got, want)
	}
}
