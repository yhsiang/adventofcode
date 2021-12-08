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

func intcode(data []int) []int {
	var input = make([]int, len(data))
	copy(input, data)
	var next int
	for {
		code := input[next]
		if code == 99 {
			break
		}
		if code == 1 {
			a := input[next+1]
			b := input[next+2]
			c := input[next+3]
			input[c] = input[a] + input[b]
			next += 3
		}
		if code == 2 {
			a := input[next+1]
			b := input[next+2]
			c := input[next+3]
			input[c] = input[a] * input[b]
			next += 3
		}
		next++

	}

	return input
}

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}
	data := util.ToInt(strings.Split(file, ","))
	data[1] = 12
	data[2] = 2
	fmt.Printf("part1: %d\n", intcode(data)[0])

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data); j++ {
			data[1] = i
			data[2] = j
			if intcode(data)[0] == 19690720 {
				fmt.Printf("part2: %d\n", 100*i+j)
			}
		}
	}
}
