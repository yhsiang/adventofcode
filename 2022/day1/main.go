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

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}

	elves := strings.Split(file, "\n\n")
	var max int = 0
	var sumCalories []int
	for _, elf := range elves {
		calories := strings.Split(elf, "\n")
		total := util.SumInt(util.ToInt(calories))
		if total > max {
			max = total
		}
		sumCalories = append(sumCalories, total)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sumCalories)))

	// first puzzle
	fmt.Printf("%d\n", max)
	// fmt.Printf("%+v", sumCalories[0:3])
	fmt.Printf("%d\n", util.SumInt(sumCalories[0:3]))
}
