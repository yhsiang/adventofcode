package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

// import "io/ioutil"

type Xmas struct {
	Nums []int64
	Sums map[int64]struct{}
}

func genSums(nums []int64) map[int64]struct{} {
	var sums = make(map[int64]struct{})
	for _, i := range nums {
		for _, j := range nums {
			sums[i+j] = struct{}{}
		}
	}

	return sums
}

func NewXmas(nums []int64) *Xmas {
	return &Xmas{
		Nums: nums,
		Sums: genSums(nums),
	}
}

func (x *Xmas) Add(num int64) {
	x.Nums = append(x.Nums[1:], num)
	x.Sums = genSums(x.Nums)
}

func sum(nums []int64) (r int64) {
	for _, n := range nums {
		r += n
	}
	return
}

func main() {
	dat, err := ioutil.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")

	var nums []int64
	for _, line := range lines {
		i, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			panic(err)
		}
		nums = append(nums, i)
	}
	// PART1
	var limit = 25
	x := NewXmas(nums[0:limit])
	var target int64
	for _, n := range nums[limit:] {
		if _, ok := x.Sums[n]; !ok {
			target = n
			break
		}
		x.Add(n)
	}
	// PART2
	var start = 0
	var end = 2
	var i = 0
	var final []int
	for i < len(nums) {
		var contiguous = make(map[int64][]int)
		for end < len(nums) {
			s := sum(nums[start:end])
			if _, ok := contiguous[s]; !ok {
				contiguous[s] = append(contiguous[s], start, end)
			}
			start += 1
			end += 1
		}
		if value, ok := contiguous[target]; ok {
			final = value
			break
		}
		i += 1
		start = 0
		end = 2 + i
	}

	fmt.Println(final)
	result := nums[final[0]:final[1]]
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	fmt.Println(result[0] + result[len(result)-1])
}
