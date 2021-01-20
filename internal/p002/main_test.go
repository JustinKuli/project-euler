package p002

import (
	"testing"
)

func TestSolve(t *testing.T) {
	want := "4613732"
	got := Solve()
	if got != want {
		t.Errorf("Got %v, want %v", got, want)
	}
}
