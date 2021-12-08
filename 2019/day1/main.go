package main

import (
	_ "embed"
	"fmt"
	"os"

	util "github.com/yhsiang/adventofcode"
)

//go:embed example
var example string

//go:embed input
var input string

func count(num int) []int {
	var fuels []int
	for {
		f := num/3 - 2
		if f <= 0 {
			break
		}
		fuels = append(fuels, f)
		num = f

	}
	return fuels
}

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}
	data := util.Read(file)
	numbers := util.ToInt(data)
	var sum int
	for _, n := range numbers {
		sum += n/3 - 2
	}
	fmt.Printf("part1: %d\n", sum)
	sum = 0
	for _, n := range numbers {
		sum += util.SumInt(count(n))
	}
	fmt.Printf("part2: %d\n", sum)

}
