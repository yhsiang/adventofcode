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

type Fish struct {
	State int64
}

func initFishes(data []int64) (fishes []*Fish) {
	for _, d := range data {
		fishes = append(fishes, &Fish{
			State: d,
		})
	}
	return
}

func (t *Fish) run() {
	if t.State == 0 {
		t.State = 7
	}
	t.State -= 1
}

func simulate(fishes []*Fish, round int) int {
	for i := 1; i <= round; i++ {
		for _, t := range fishes {
			if t.State == 0 {
				fishes = append(fishes, &Fish{
					State: 8,
				})
			}
			t.run()
		}
		// fmt.Printf("%d: %d\n", i, len(timers))
		// print(timers)
	}
	return len(fishes)
}

func fastSim(init_fishes []int64, round int) int64 {
	var fishes = [9]int64{0}
	for _, d := range init_fishes {
		fishes[d] += 1
	}
	// fmt.Printf("%v\n", fishes)

	for i := 1; i <= round; i++ {
		temp := fishes[0]
		for j := 1; j < 9; j++ {
			// fmt.Printf("%d, %d\n", j-1, j)
			fishes[j-1] = fishes[j]
		}
		fishes[6] += temp
		fishes[8] = temp
		// fmt.Printf("%d %v\n", i, fishes)
	}

	var sum int64 = 0
	for _, d := range fishes {
		sum += d
	}

	return sum
}

func print(fishes []*Fish) {
	for _, t := range fishes {
		fmt.Printf("%d", t.State)
	}
	fmt.Printf("\n")
}

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}

	data := util.ToInt64(strings.Split(file, ","))
	fishes := initFishes(data)

	fmt.Printf("part1: %d\n", simulate(fishes, 80))
	fmt.Printf("part2: %d\n", fastSim(data, 256))

}
