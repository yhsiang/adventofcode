package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	dat, err := ioutil.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")

	var nums = []int64{0}
	for _, line := range lines {
		i, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			panic(err)
		}
		nums = append(nums, i)
	}

	// var diffs = map[int64]int64{
	// 	1: 0,
	// 	2: 0,
	// 	3: 1,
	// }

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	// PART 1
	// var previous int64
	// for _, n := range nums {
	// 	// fmt.Println(previous, n, n-previous)
	// 	diffs[n-previous] += 1
	// 	previous = n

	// }
	// fmt.Println(diffs[1] * diffs[3])

	var table [][]int64
	var ctn []int64
	for i, n := range nums {
		if i == 0 {
			ctn = append(ctn, n)
			continue
		}

		if (n - nums[i-1]) == 1 {
			ctn = append(ctn, n)
		} else {
			if len(ctn) > 0 {
				table = append(table, ctn)
				ctn = []int64{}
			}
			ctn = append(ctn, n)
		}
	}

	if len(ctn) > 0 {
		table = append(table, ctn)
		ctn = []int64{}
	}

	// for _, t := range table {
	// 	fmt.Println(t, len(t))
	// }
	var pow2 float64 = 0
	var pow7 float64
	for _, t := range table {
		if len(t) >= 5 {
			pow7 += 1
		}

		if len(t) == 4 {
			pow2 += 2
		}

		if len(t) == 3 {
			pow2 += 1
		}
	}
	// [a,b] base 1 (2^0)
	// [a,b,c] base 2 (2^1)
	// [a,b,c,d] base 4 (2^2)
	// [a,b,c,d,e] base 7 (1+2+4)
	// [a,b,c,d,e,f] base 13 (2+4+7)

	fmt.Println(pow2, pow7)
	fmt.Println(int64(math.Pow(2, pow2) * math.Pow(7, pow7)))
	// hint
	// https://www.reddit.com/r/adventofcode/comments/ka9pc3/2020_day_10_part_2_suspicious_factorisation/
}
