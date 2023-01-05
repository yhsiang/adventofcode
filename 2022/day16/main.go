package main

import (
	_ "embed"
	"os"
	"strings"
)

//go:embed example
var example string

//go:embed input
var input string

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}

	strings.Split(file, "\n")
}

type Valve struct {
	FlowRate int
	Open     bool
}

// AA 0
// DD 20
// CC
// BB 13
// AA
// II
// JJ
