package p006

import (
	"testing"
)

func TestSolve(t *testing.T) {
	want := "25164150"
	got := Solve()
	if got != want {
		t.Errorf("Got %v, want %v", got, want)
	}
}
