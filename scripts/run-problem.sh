#!/usr/bin/env bash
set -euo pipefail

problem_name=$1
cat > main.go <<EOF
package main

import (
	"fmt"

	"github.com/JustinKuli/project-euler/internal/${problem_name}"
)

func main() {
	fmt.Println(${problem_name}.Solve())
}
EOF

go run main.go
