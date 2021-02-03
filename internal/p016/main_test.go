package p016

import (
	"testing"
)

func TestSolve(t *testing.T) {
	want := "1366"
	got := Solve()
	if got != want {
		t.Errorf("Got %v, want %v", got, want)
	}
}
