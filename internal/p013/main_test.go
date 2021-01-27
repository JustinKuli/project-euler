package p013

import (
	"testing"
)

func TestSolve(t *testing.T) {
	want := "5537376230"
	got := Solve()
	if got != want {
		t.Errorf("Got %v, want %v", got, want)
	}
}
