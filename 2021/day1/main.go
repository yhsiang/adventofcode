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

	var windows = append([]int64{0}, nums[0:2]...)
	var three_nums []int64
	for _, d := range nums[2:] {
		windows = append(windows[1:], d)
		three_nums = append(three_nums, util.SumInt64(windows))
	}
	fmt.Printf("part2: %d\n", count(three_nums))
}
