package util

import (
	"fmt"
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

func Int64(data string) (num int64, err error) {
	num, err = strconv.ParseInt(data, 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func BinToInt64(data string) (num int64, err error) {
	num, err = strconv.ParseInt(data, 2, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
