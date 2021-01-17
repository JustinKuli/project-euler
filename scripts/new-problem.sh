#!/usr/bin/env bash
set -euo pipefail

problem_name=$1

mkdir "internal/${problem_name}"
cat > "internal/${problem_name}/main.go" <<EOF
package ${problem_name}

import (
	"fmt"
)

func Solve() {
	fmt.Println("Hello world")
}
EOF
