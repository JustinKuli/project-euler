#!/usr/bin/env bash
set -euo pipefail

problem_name=$1

mkdir "internal/${problem_name}"
cat > "internal/${problem_name}/main.go" <<EOF
package ${problem_name}

import (
	"fmt"
)

// Solve this problem:
//
func Solve() {
	fmt.Println("Hello world")
}
EOF

cat > "internal/${problem_name}/main_test.go" <<EOF
package ${problem_name}

import (
	"fmt"
	"testing"
)

func add(a, b int) int {
	return a+b
}

func TestAdd(t *testing.T) {
	var tests = []struct {
		a, b int
		out int
	}{
		{1, 2, 3},
		{2, 2, 4},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Add %v and %v", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := add(tt.a, tt.b)
			if ans != tt.out {
				t.Errorf("Got %v, want %v", ans, tt.out)
			}
		})
	}
}
EOF
