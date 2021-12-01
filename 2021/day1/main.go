package main

import (
	"fmt"

	util "github.com/yhsiang/adventofcode"
)

func main() {
	data := util.Read(util.GetInput())
	var nums = util.ToInt64(data)

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
	fmt.Printf("part1: %d\n", increased)

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
	var increased2 int = 0
	// fmt.Printf("%+v", nums)
	for i, d := range three_nums {
		if i == 0 {
			continue
		}
		if d-three_nums[i-1] > 0 {
			increased2 += 1
		}
	}
	fmt.Printf("part2: %d\n", increased2)

}
