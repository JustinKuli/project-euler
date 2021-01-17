#!/usr/bin/env bash
set -euo pipefail

problem_name=$1
cat > main.go <<EOF
package main

import (
	"github.com/JustinKuli/project-euler/internal/${problem_name}"
)

func main() {
	${problem_name}.Solve()
}
EOF

go run main.go
