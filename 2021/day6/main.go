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
	fishes := init_fishes
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
	var fishes = make([]int64, 9)
	for _, d := range init_fishes {
		fishes[d] += 1
	}

	for i := 1; i <= round; i++ {
		new_fish := append(fishes[1:], fishes[0])
		new_fish[6] += fishes[0]
		fishes = new_fish
	}

	return util.Sum(fishes)
}

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}

	data := util.ToInt64(strings.Split(file, ","))

	fmt.Printf("part1: %d\n", sim(data, 80))
	fmt.Printf("part2: %d\n", fastSim(data, 256))
}
