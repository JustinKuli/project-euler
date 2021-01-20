package p008

import (
	"testing"
)

func TestSolve(t *testing.T) {
	want := "23514624000"
	got := Solve()
	if got != want {
		t.Errorf("Got %v, want %v", got, want)
	}
}
