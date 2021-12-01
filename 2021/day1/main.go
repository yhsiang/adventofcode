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

func count(nums []int64) int {
	var increased int = 0
	// fmt.Printf("%+v", nums)
	for i, d := range nums {
		if i == 0 {
			continue
		}
		if d-nums[i-1] > 0 {
			increased += 1
		}
	}
	return increased
}

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}

	data := util.Read(file)
	var nums = util.ToInt64(data)
	fmt.Printf("part1: %d\n", count(nums))

	var windows []int64
	var three_nums []int64
	for _, d := range nums {
		windows = append(windows, d)
		if len(windows) == 3 {
			three_nums = append(three_nums, windows[0]+windows[1]+windows[2])
			windows = windows[1:]
			// fmt.Printf("%+v", windows)
		}
	}
	// fmt.Printf("%+v", three_nums)
	fmt.Printf("part2: %d\n", count(three_nums))
}
