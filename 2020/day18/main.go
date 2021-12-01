package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type NumStack struct {
	values []int
}

func (s *NumStack) push(i int) {
	s.values = append(s.values, i)
}

func (s *NumStack) pop() int {
	if len(s.values) == 0 {
		return -1
	}

	idx := len(s.values) - 1
	n := s.values[idx]
	s.values = s.values[:idx]
	return n
}

func (s *NumStack) isEmpty() bool {
	return len(s.values) == 0
}

type StrStack struct {
	values []string
}

func (s *StrStack) push(i string) {
	s.values = append(s.values, i)
}

func (s *StrStack) pop() string {
	if len(s.values) == 0 {
		return ""
	}
	idx := len(s.values) - 1
	n := s.values[idx]
	s.values = s.values[:idx]
	return n
}

func (s *StrStack) isEmpty() bool {
	return len(s.values) == 0
}

type Expression struct {
	nums NumStack
	ops  StrStack
}

func part1(str string) int {
	ss := strings.ReplaceAll(str, " ", "")
	var ops StrStack
	var nums NumStack
	for _, s := range strings.Split(ss, "") {
		// fmt.Println(s, nums.values, ops.values)
		switch s {
		case "+", "*", "(":
			ops.push(s)
		case ")":
			ops.pop()
			if len(nums.values) > 1 && ops.values[len(ops.values)-1] != "(" {
				var sum = 0
				a := nums.pop()
				b := nums.pop()
				op := ops.pop()
				switch op {
				case "*":
					sum = a * b
				case "+":
					sum = a + b
				}
				nums.push(sum)
			}
		default:
			i, _ := strconv.ParseInt(s, 10, 64)
			if len(ops.values) > 0 && ops.values[len(ops.values)-1] != "(" {
				var sum = 0
				op := ops.pop()
				a := nums.pop()
				switch op {
				case "*":
					sum = a * int(i)
				case "+":
					sum = a + int(i)
				}
				nums.push(sum)
			} else {
				nums.push(int(i))
			}
		}
	}

	return nums.pop()
}

func part2(str string) int {
	ss := strings.ReplaceAll(str, " ", "")
	var nums NumStack
	var ops StrStack

	for _, s := range strings.Split(ss, "") {
		switch s {
		case "+":
			ops.push(s)
		case "*":
			ops.push(s)
		case "(":
			ops.push(s)
		case ")":
			var tmpNum []int
			for {
				if ops.values[len(ops.values)-1] == "(" {
					ops.pop()
					break
				}
				op := ops.pop()
				a := nums.pop()
				b := nums.pop()
				if op == "+" {
					nums.push(a + b)
				} else {
					tmpNum = append(tmpNum, a)
					// tmp = append(tmp, op)
					nums.push(b)
				}
			}
			var n = nums.pop()
			for _, num := range tmpNum {
				n *= num
			}
			nums.push(n)
		default:
			i, _ := strconv.ParseInt(s, 10, 64)
			if len(nums.values) >= 1 && ops.values[len(ops.values)-1] == "+" {
				ops.pop()
				a := nums.pop()
				nums.push(a + int(i))
			} else {
				nums.push(int(i))
			}
		}
	}

	var tmpNum []int
	for len(ops.values) > 0 {
		op := ops.pop()
		a := nums.pop()
		b := nums.pop()
		if op == "+" {
			nums.push(a + b)
		} else {
			tmpNum = append(tmpNum, a)
			nums.push(b)
		}
	}

	var n = nums.pop()
	for _, num := range tmpNum {
		n *= num
	}
	nums.push(n)

	return nums.pop()
}

func main() {
	dat, err := ioutil.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")

	// Part1
	// var sum int
	// for _, line := range lines {
	// 	n := part1(line)
	// 	sum += n
	// }
	// fmt.Println(sum)

	var sum int
	for _, line := range lines {
		n := part2(line)
		// fmt.Println(i, n)
		sum += n
	}
	fmt.Println(sum)

}
