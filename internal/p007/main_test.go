package p007

import (
	"testing"
)

func TestSolve(t *testing.T) {
	want := "104743"
	got := Solve()
	if got != want {
		t.Errorf("Got %v, want %v", got, want)
	}
}
