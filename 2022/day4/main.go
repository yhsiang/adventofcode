package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"

	util "github.com/yhsiang/adventofcode"
)

//go:embed example
var example string

//go:embed input
var input string

func toRange(str string) (int, int) {
	s := strings.Split(str, "-")
	start, _ := util.Int(s[0])
	end, _ := util.Int(s[1])
	return start, end
}

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}

	assignments := strings.Split(file, "\n")
	var containedCount int
	for _, assignment := range assignments {
		pairs := strings.Split(assignment, ",")
		s1, e1 := toRange(pairs[0])
		s2, e2 := toRange(pairs[1])

		if s2 >= s1 && e2 <= e1 {
			containedCount++
			continue
		}
		if s1 >= s2 && e1 <= e2 {
			containedCount++
			continue
		}
	}
	// first puzzle
	fmt.Printf("%d\n", containedCount)

	var noOverlapCount int
	for _, assignment := range assignments {
		pairs := strings.Split(assignment, ",")
		s1, e1 := toRange(pairs[0])
		s2, e2 := toRange(pairs[1])

		if s2 > e1 {
			noOverlapCount++
			continue
		}
		if e2 < s1 {
			noOverlapCount++
			continue
		}
	}
	fmt.Printf("%d\n", len(assignments)-noOverlapCount)
}
