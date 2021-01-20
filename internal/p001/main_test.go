package p001

import (
	"testing"
)

func TestSolve(t *testing.T) {
	want := "233168"
	got := Solve()
	if got != want {
		t.Errorf("Got %v, want %v", got, want)
	}
}
