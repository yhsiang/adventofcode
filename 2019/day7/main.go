package main

import (
	_ "embed"
	"fmt"
	"os"
	"sort"
	"strings"

	util "github.com/yhsiang/adventofcode"
)

//go:embed example
var example string

//go:embed input
var input string

type IntCode struct {
	Name       string
	Input      int
	Program    []int
	Phase      int
	Next       int
	InputCount int
}

func initIntCode(name string, data []int, input int, phase int) *IntCode {
	program := make([]int, len(data))
	copy(program, data)
	return &IntCode{
		Name:       name,
		Input:      input,
		Program:    program,
		Phase:      phase,
		Next:       0,
		InputCount: 0,
	}
}

func (ic *IntCode) getOpCode(next int) (int, []int) {
	i := ic.Program[next]
	opcode := i % 100
	var modes []int
	i = i / 100
	modes = append(modes, i%10)
	i = i / 10
	modes = append(modes, i%10)

	return opcode, modes
}

func (ic *IntCode) halted() bool {
	return ic.Program[ic.Next] == 99
}

func (ic *IntCode) getP(index int, modes []int, program []int, programIndex int) int {
	if modes[index] == 1 {
		return program[programIndex]
	}
	return program[program[programIndex]]
}

func (ic *IntCode) run() int {
	next := ic.Next
	program := ic.Program
	var value int
loop:
	for next < len(ic.Program) {
		op, modes := ic.getOpCode(next)
		switch {
		case op == 99:
			break loop
		case op == 1:
			c := program[next+3]
			program[c] = ic.getP(0, modes, program, next+1) + ic.getP(1, modes, program, next+2)
			next += 4
		case op == 2:
			c := program[next+3]
			program[c] = ic.getP(0, modes, program, next+1) * ic.getP(1, modes, program, next+2)
			next += 4
		case op == 3:
			// input
			var choice int
			if ic.InputCount == 0 {
				choice = ic.Phase
				ic.InputCount += 1
			} else {
				choice = ic.Input
			}
			a := program[next+1]
			program[a] = choice
			next += 2
		case op == 4:
			value = ic.getP(0, modes, program, next+1)
			next += 2
			ic.Next = next
			break loop
		case op == 5:
			// jump-if-true
			if ic.getP(0, modes, program, next+1) != 0 {
				next = ic.getP(1, modes, program, next+2)
			} else {
				next += 3
			}
		case op == 6:
			// jump-if-false
			if ic.getP(0, modes, program, next+1) == 0 {
				next = ic.getP(1, modes, program, next+2)
			} else {
				next += 3
			}
		case op == 7:
			// less than
			c := program[next+3]
			if ic.getP(0, modes, program, next+1) < ic.getP(1, modes, program, next+2) {
				program[c] = 1
			} else {
				program[c] = 0
			}
			next += 4
		case op == 8:
			// equals
			c := program[next+3]
			if ic.getP(0, modes, program, next+1) == ic.getP(1, modes, program, next+2) {
				program[c] = 1
			} else {
				program[c] = 0
			}
			next += 4
		}
	}
	return value

}

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}

	data := util.ToInt(strings.Split(file, ","))
	var output []int
	util.Perm([]rune("01234"), func(a []rune) {
		phases := util.RuneToInt(a)
		input := 0
		for i, phase := range phases {
			ic := initIntCode(fmt.Sprintf("%d", i), data, input, phase)
			input = ic.run()
		}
		output = append(output, input)
	})

	sort.Ints(output)
	fmt.Printf("part1: %d\n", output[len(output)-1])

	output = []int{}
	util.Perm([]rune("56789"), func(r []rune) {
		phases := util.RuneToInt(r)
		a := initIntCode("A", data, 0, phases[0])
		b := initIntCode("B", data, 0, phases[1])
		c := initIntCode("C", data, 0, phases[2])
		d := initIntCode("D", data, 0, phases[3])
		e := initIntCode("E", data, 0, phases[4])

		for !e.halted() {
			b.Input = a.run()
			c.Input = b.run()
			d.Input = c.run()
			e.Input = d.run()
			a.Input = e.run()
		}
		output = append(output, a.Input)
	})
	sort.Ints(output)
	fmt.Printf("part2: %d\n", output[len(output)-1])

}
