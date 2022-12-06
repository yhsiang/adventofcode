package main

import (
	_ "embed"
	"fmt"
	"os"
	"regexp"
	"strings"

	util "github.com/yhsiang/adventofcode"
)

//go:embed example
var example string

//go:embed input
var input string

func NewStacks(str string) []string {
	lines := strings.Split(str, "\n")
	lastLine := lines[len(lines)-1]
	numbers := strings.Split(lastLine, "   ")
	stackLen, _ := util.Int(strings.Trim(numbers[len(numbers)-1], " "))

	crates := util.ReverseStringSlice(lines[:len(lines)-1])
	// fmt.Printf("%+v\n", lines[:len(lines)-1])
	// fmt.Printf("%+v\n", crates)
	var stacks []string
	for i := 1; i < stackLen*4-2; i += 4 {
		var stack string
		for _, line := range crates {
			// 10 3*4 - 2 (// 34 9*4 -2)
			// 1,5,9,13
			stack += string(line[i])
		}
		stacks = append(stacks, stack)
	}
	// fmt.Printf("%+v", stacks)
	// ZN
	// MCD
	// p
	return stacks
}

var re = regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

func ParseCommand(command string) (int, int, int) {
	// move 1 from 2 to 1
	matched := re.FindSubmatch([]byte(command))
	// fmt.Printf("%+v\n", matched)
	size, _ := util.Int(fmt.Sprintf("%s", matched[1]))
	from, _ := util.Int(fmt.Sprintf("%s", matched[2]))
	to, _ := util.Int(fmt.Sprintf("%s", matched[3]))
	return size, from - 1, to - 1
}

func Move(stacks []string, command string) []string {
	size, from, to := ParseCommand(command)
	// fmt.Printf("%d, %d, %d \n", size, from, to)
	fromStr := stacks[from]
	toStr := stacks[to]
	var extract string
	var remain string
	for i := len(fromStr) - 1; i >= 0; i-- {
		if string(fromStr[i]) == " " {
			continue
		}
		if len(extract) == size {
			remain += string(fromStr[i])
		} else {
			extract += string(fromStr[i])
		}
	}
	// fmt.Printf("%s, %s\n", extract, remain)

	stacks[from] = util.Reverse(remain)
	for _, str := range extract {
		if strings.Contains(toStr, " ") {
			toStr = strings.Replace(toStr, " ", string(str), 1)
		} else {
			toStr += string(str)
		}
	}
	stacks[to] = toStr
	return stacks
}

func Move9001(stacks []string, command string) []string {
	size, from, to := ParseCommand(command)
	// fmt.Printf("%d, %d, %d \n", size, from, to)
	fromStr := stacks[from]
	toStr := stacks[to]
	var extract string
	var remain string
	for i := len(fromStr) - 1; i >= 0; i-- {
		if string(fromStr[i]) == " " {
			continue
		}
		if len(extract) == size {
			remain += string(fromStr[i])
		} else {
			extract += string(fromStr[i])
		}
	}
	// fmt.Printf("%s, %s\n", extract, remain)

	stacks[from] = util.Reverse(remain)
	for _, str := range util.Reverse(extract) {
		if strings.Contains(toStr, " ") {
			toStr = strings.Replace(toStr, " ", string(str), 1)
		} else {
			toStr += string(str)
		}
	}
	stacks[to] = toStr
	return stacks
}

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}

	inputs := strings.Split(file, "\n\n")
	commands := strings.Split(inputs[1], "\n")
	stacks := NewStacks(inputs[0])
	for _, command := range commands {
		stacks = Move(stacks, command)
	}

	var result string
	for _, stack := range stacks {
		result += string(stack[len(stack)-1])
	}

	// first puzzle
	fmt.Printf("%s\n", result)

	stacks = NewStacks(inputs[0])
	for _, command := range commands {
		// fmt.Printf("%+v\n", stacks)
		stacks = Move9001(stacks, command)
	}

	result = ""
	for _, stack := range stacks {
		result += string(stack[len(stack)-1])
	}
	fmt.Printf("%s\n", result)
}
