package util

import (
	"strconv"
	"strings"
)

func Read(input string) []string {
	data := strings.Split(string(input), "\n")
	return data
}

func ToInt64(data []string) []int64 {
	var nums []int64
	for _, d := range data {
		i, err := strconv.ParseInt(d, 10, 64)
		if err != nil {
			panic(err)
		}
		nums = append(nums, i)
	}
	return nums
}
