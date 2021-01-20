package p003

import (
	"testing"
)

func TestSolve(t *testing.T) {
	want := "6857"
	got := Solve()
	if got != want {
		t.Errorf("Got %v, want %v", got, want)
	}
}
