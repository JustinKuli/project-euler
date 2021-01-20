package p004

import (
	"testing"
)

func TestSolve(t *testing.T) {
	want := "906609"
	got := Solve()
	if got != want {
		t.Errorf("Got %v, want %v", got, want)
	}
}
