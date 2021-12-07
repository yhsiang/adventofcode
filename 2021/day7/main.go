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

// generate table first
func seqArr(num int) []int {
	var sum = 0
	var output = make([]int, num+1)
	for i := 1; i <= num; i++ {
		sum += i
		output[i] = sum
	}
	return output
}

func seq(num int) int {
	var sum = 0
	for i := 1; i <= num; i++ {
		sum += i
	}
	return sum
}

func fastSeq(num int) int {
	return num * (num + 1) / 2
}

func positions(input []int, max int, part2 bool) int {
	var pos = []int{}
	// var seqs = seqArr(max)
	for position := 0; position < max; position++ {
		var fuels = []int{}
		for i := range input {
			p := util.Abs(input[i] - position)
			if part2 {
				fuels = append(fuels, fastSeq(p))
			} else {
				fuels = append(fuels, p)
			}
		}
		pos = append(pos, util.SumInt(fuels))
	}
	min, _ := util.MinMax(pos)
	return min
}

// will be median in sorted array
func mostCommon(input []int) int {
	var freq = make(map[int]int)
	for _, d := range input {
		freq[d] += 1
	}
	var maxKey int
	var maxValue = 0
	for key, value := range freq {
		if value > maxValue {
			maxValue = value
			maxKey = key
		}
	}
	return maxKey
}

func cal(input []int, max int, part2 bool) int {
	var pos = []int{}
	var seqs = seqArr(max)
	target := mostCommon(input)
	var fuels = []int{}
	for i := range input {
		p := util.Abs(input[i] - target)
		if part2 {
			fuels = append(fuels, seqs[p])
		} else {
			fuels = append(fuels, p)
		}
	}
	pos = append(pos, util.SumInt(fuels))
	min, _ := util.MinMax(pos)
	return min
}

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}
	data := util.ToInt(strings.Split(file, ","))
	_, max := util.MinMax(data)

	fmt.Printf("part1: %d\n", positions(data, max, false))
	fmt.Printf("part2: %d\n", positions(data, max, true))

	// fmt.Printf("part1: %d\n", cal(data, max, false))

}
