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

func sim(init_fishes []int64, round int) int {
	fishes := make([]int64, len(init_fishes))
	copy(fishes, init_fishes)
	for i := 1; i <= round; i++ {
		for i := range fishes {
			if fishes[i] == 0 {
				fishes = append(fishes, 8)
				fishes[i] = 7
			}
			fishes[i] -= 1
		}
	}
	return len(fishes)
}

func fastSim(init_fishes []int64, round int) int64 {
	// [9]int64{} != make([]int64, 9)
	var fishes = make([]int64, 9)
	for _, d := range init_fishes {
		fishes[d] += 1
	}

	for i := 1; i <= round; i++ {
		fishes[7] += fishes[0]
		fishes = append(fishes[1:], fishes[0])
	}

	return util.SumInt64(fishes)
}

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}

	data := util.ToInt64(strings.Split(file, ","))

	fmt.Printf("part1: %d\n", sim(data, 80)) // 210 still works
	fmt.Printf("part2: %d\n", fastSim(data, 256))
}
