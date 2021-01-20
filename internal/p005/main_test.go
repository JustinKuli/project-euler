package p005

import (
	"testing"
)

func TestSolve(t *testing.T) {
	want := "232792560"
	got := Solve()
	if got != want {
		t.Errorf("Got %v, want %v", got, want)
	}
}
