package main

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"

	util "github.com/yhsiang/adventofcode"
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

	data := util.Read(file)
	// fmt.Print("%v", data)
	var depth = 0
	var position = 0
	for _, d := range data {
		command := strings.Split(d, " ")
		// fmt.Printf("%s\n", command[0])
		n, _ := strconv.Atoi(command[1])

		switch command[0] {
		case "forward":
			position += n
		case "down":
			depth += n
		case "up":
			depth -= n
		}
	}
	fmt.Printf("part1: %d\n", position*depth)

	depth = 0
	position = 0
	var aim = 0
	for _, d := range data {
		command := strings.Split(d, " ")
		// fmt.Printf("%s\n", command[0])
		n, _ := strconv.Atoi(command[1])

		switch command[0] {
		case "forward":
			position += n
			depth += aim * n
		case "down":
			aim += n
		case "up":
			aim -= n
		}
		// fmt.Printf("%d, %d, %d\n", aim, position, depth)
	}
	fmt.Printf("part2: %d", position*depth)
}
