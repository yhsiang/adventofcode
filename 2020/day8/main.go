package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func handleCodes(lines []string) (finished bool, acc int) {
	var visited = make(map[int]bool, len(lines))
	for i := range lines {
		visited[i] = false
	}
	var step = 0
	for {
		if v, ok := visited[step]; ok && v {
			break
		}
		if step >= len(lines) {
			finished = true
			break
		}

		visited[step] = true
		str := lines[step]
		ops := strings.Split(str, " ")

		i, err := strconv.ParseInt(ops[1], 10, 64)
		if err != nil {
			panic(err)
		}
		// fmt.Println(ops, step, acc)
		op := ops[0]
		switch op {
		case "nop":
			step += 1
		case "acc":
			acc += int(i)
			step += 1
		case "jmp":
			step += int(i)
		}
	}

	return
}
func main() {
	dat, err := ioutil.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")
	// Part 1
	// fmt.Println(handleCodes(lines))
	for i, line := range lines {
		var codes = make([]string, len(lines))
		copy(codes, lines)
		ops := strings.Split(line, " ")
		if ops[0] == "nop" || ops[0] == "jmp" {
			if ops[0] == "nop" {
				codes[i] = strings.Replace(line, "nop", "jmp", 1)
			}

			if ops[0] == "jmp" {
				codes[i] = strings.Replace(line, "jmp", "nop", 1)
			}
		}

		finished, acc := handleCodes(codes)
		if finished {
			fmt.Println(i, acc)
			break
		}
	}
	// hint
	// https://www.reddit.com/r/adventofcode/comments/k8z1fi/2020_day8_part_2_python_so_basically_my_program/
}
