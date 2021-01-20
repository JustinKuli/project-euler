#!/usr/bin/env bash
set -euo pipefail

problem_name=$1

mkdir "internal/${problem_name}"
cat > "internal/${problem_name}/main.go" <<EOF
package ${problem_name}

// Solve this problem:
//
func Solve() String {
	return "Hello world"
}
EOF

cat > "internal/${problem_name}/main_test.go" <<EOF
package ${problem_name}

import (
	"testing"
)

func TestSolve(t *testing.T) {
	want := "???"
	got := Solve()
	if got != want {
		t.Errorf("Got %v, want %v", got, want)
	}
}
EOF
