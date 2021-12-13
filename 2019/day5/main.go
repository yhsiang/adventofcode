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

func getOpCode(i int) (int, []int) {

	opcode := i % 100
	var modes []int
	i = i / 100
	modes = append(modes, i%10)
	i = i / 10
	modes = append(modes, i%10)

	return opcode, modes
}

func getP(index int, modes []int, programIndex int, program []int) int {
	if modes[index] == 1 {
		return program[programIndex]
	}
	return program[program[programIndex]]
}

func intcode(data []int, input int, part2 bool) string {
	var program = make([]int, len(data))
	copy(program, data)
	var next int
	var output string
	for next < len(data) {
		code := program[next]
		op, modes := getOpCode(code)
		if op == 99 {
			break
		}
		if op == 1 {
			c := program[next+3]
			program[c] = getP(0, modes, next+1, program) + getP(1, modes, next+2, program)
			next += 4
		}
		if op == 2 {
			c := program[next+3]
			program[c] = getP(0, modes, next+1, program) * getP(1, modes, next+2, program)
			next += 4
		}
		if op == 3 {
			// input
			a := program[next+1]
			program[a] = input
			next += 2
		}
		if op == 4 {
			//output
			output += fmt.Sprintf("%d", getP(0, modes, next+1, program))
			next += 2
		}

		if part2 && op == 5 {
			// jump-if-true
			if getP(0, modes, next+1, program) != 0 {
				next = getP(1, modes, next+2, program)
				// 	input[next+3] = p2
			} else {
				next += 3
			}
		}
		if part2 && op == 6 {
			// jump-if-false
			if getP(0, modes, next+1, program) == 0 {
				next = getP(1, modes, next+2, program)
			} else {
				next += 3
			}
		}
		if part2 && op == 7 {
			// less than
			c := program[next+3]
			if getP(0, modes, next+1, program) < getP(1, modes, next+2, program) {
				program[c] = 1
			} else {
				program[c] = 0
			}

			next += 4
		}
		if part2 && op == 8 {
			// equals
			c := program[next+3]
			if getP(0, modes, next+1, program) == getP(1, modes, next+2, program) {
				program[c] = 1
			} else {
				program[c] = 0
			}

			next += 4
		}
	}

	return output
}

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}

	data := util.ToInt(strings.Split(file, ","))
	fmt.Printf("part1: %s\n", intcode(data, 1, false))
	fmt.Printf("part2: %s\n", intcode(data, 5, true))
}
