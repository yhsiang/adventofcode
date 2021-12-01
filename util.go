package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func Read(input string) []string {
	dat, err := ioutil.ReadFile(input)
	if err != nil {
		panic(err)
	}

	data := strings.Split(string(dat), "\n")
	return data
}

func GetInput() string {
	if len(os.Args) == 2 {
		return fmt.Sprintf("./%s", os.Args[1])
	}

	return "./example"
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
