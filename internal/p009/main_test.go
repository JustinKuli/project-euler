package p009

import (
	"testing"
)

func TestSolve(t *testing.T) {
	want := "31875000"
	got := Solve()
	if got != want {
		t.Errorf("Got %v, want %v", got, want)
	}
}
