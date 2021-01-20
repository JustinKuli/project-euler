package p010

import (
	"testing"
)

func TestSolve(t *testing.T) {
	want := "142913828922"
	got := Solve()
	if got != want {
		t.Errorf("Got %v, want %v", got, want)
	}
}
