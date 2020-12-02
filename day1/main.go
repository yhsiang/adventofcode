package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	dat, err := ioutil.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	data := strings.Split(string(dat), "\n")
	var nums []int64
	for _, d := range data {
		i, err := strconv.ParseInt(d, 10, 64)
		if err != nil {
			panic(err)
		}
		nums = append(nums, i)
	}

	// var t1, t2, t3 int64
	// for _, n := range nums {
	// 	t1 = 2020 - n
	// 	for _, i := range nums {
	// 		if t1 != i {
	// 			continue
	// 		}
	// 		fmt.Printf("%d %d\n", t1, i)
	// 		t2 = i
	// 	}
	// 	if t1 == t2 {
	// 		t3 = n
	// 		break
	// 	}
	// }
	// fmt.Println(t1 * t3)
	var a, b, c int64
	var found = false
	for i1, n1 := range nums {
		for i2, n2 := range nums {
			if i1 == i2 {
				continue
			}
			for i3, n3 := range nums {
				if i2 == i3 {
					continue
				}
				if n1+n2+n3 != 2020 {
					continue
				}
				// fmt.Println(n1, n2, n3)
				found = true
				a = n1
				b = n2
				c = n3
				break
			}
			if found {
				break
			}
		}
	}
	fmt.Println(a * b * c)
}
