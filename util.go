package util

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func ByLine(input string) []string {
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

func ToInt(data []string) []int {
	var nums []int
	for _, d := range data {
		i, err := strconv.ParseInt(d, 10, 64)
		if err != nil {
			panic(err)
		}
		nums = append(nums, int(i))
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

// TODO: generic
func SumInt64(data []int64) int64 {
	var sum int64 = 0
	for _, d := range data {
		sum += d
	}
	return sum
}

func SumInt(data []int) int {
	var sum int = 0
	for _, d := range data {
		sum += d
	}
	return sum
}

func MultiplyInt(data []int) int {
	var sum int = 1
	for _, d := range data {
		sum *= d
	}
	return sum
}

func Abs(x int) int {
	return int(math.Abs(float64(x)))
}

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}
