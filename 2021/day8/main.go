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

// 2-length: 1
// 3-length: 7
// 4-length: 4
// 5-length: 2,3,5
// 6-length: 6,0,9
// 7-length: 8
// Algorithm
// 1. 1, 7, 4, 8
// 2. 6,0,9 -> contains 4: 9 -> contains 7: 0
// 3. 2,3,5 -> 6 contains: 5 -> contains 7: 3

type Set map[string]struct{}

func (s Set) intersect(b Set) Set {
	var c = make(Set)
	for k := range s {
		if _, ok := b[k]; ok {
			c[k] = struct{}{}
		}
	}
	return c
}

func analyze(front []string, back []string) (output string) {
	var fronts = make(map[int]Set)
	for _, d := range front {
		var set = make(Set)
		for _, s := range strings.Split(d, "") {
			set[s] = struct{}{}
		}
		fronts[len(d)] = set
	}
	var backs []Set
	for _, d := range back {
		var set = make(Set)
		for _, s := range strings.Split(d, "") {
			set[s] = struct{}{}
		}
		backs = append(backs, set)
	}

	for _, s := range backs {
		a := len(s)
		b := len(s.intersect(fronts[4]))
		c := len(s.intersect(fronts[2]))
		switch {
		case a == 2:
			output += "1"
		case a == 3:
			output += "7"
		case a == 4:
			output += "4"
		case a == 7:
			output += "8"
		case a == 5 && b == 2:
			output += "2"
		case a == 5 && b == 3 && c == 1:
			output += "5"
		case a == 5 && b == 3 && c == 2:
			output += "3"
		case a == 6 && b == 4:
			output += "9"
		case a == 6 && b == 3 && c == 1:
			output += "6"
		case a == 6 && b == 3 && c == 2:
			output += "0"
		}
	}

	return
}

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}
	lines := util.ByLine(file)
	var uniq int
	for _, line := range lines {
		digits := strings.Split(line, " | ")
		remain := strings.Split(digits[1], " ")
		for _, digit := range remain {
			if (len(digit) >= 2 && len(digit) <= 4) || len(digit) == 7 {
				uniq++
			}
		}
	}
	fmt.Printf("part1: %d\n", uniq)
	var data []string
	for _, line := range lines {
		digits := strings.Split(line, " | ")
		front := strings.Split(digits[0], " ")
		back := strings.Split(digits[1], " ")
		data = append(data, analyze(front, back))
	}

	fmt.Printf("part2: %d\n", util.SumInt(util.ToInt(data)))
}
