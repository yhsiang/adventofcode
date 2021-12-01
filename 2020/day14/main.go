package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var reMask = regexp.MustCompile(`^mask = ([X01]+)`)
var reMem = regexp.MustCompile(`^mem\[(\d+)\] = (\d+)`)

type Program struct {
	mask map[int]int64
	mem  map[int64]int64
}

func NewProgram() *Program {
	return &Program{
		mem: make(map[int64]int64),
	}
}

func (p *Program) updateMask(str string) {
	// binary := strings.ReplaceAll(str, "X", "0")
	s := reMask.FindSubmatch([]byte(strings.TrimSpace(str)))
	if len(s) != 2 {
		return
	}
	var mask = make(map[int]int64)
	for i, m := range strings.Split(string(s[1]), "") {
		switch m {
		case "0":
			mask[i] = 0
		case "1":
			mask[i] = 1
		case "X":
		}
	}
	p.mask = mask
}

// func stringToBinarySlice(str string) (r []int64) {
// 	for _, b := range strings.Split(str, "") {
// 		n, _ := strconv.ParseInt(b, 10, 64)
// 		r = append(r, n)
// 	}
// 	return r
// }

func numToBinarySlice(n int64) []int64 {
	var nums = make([]int64, 36)
	binaryStr := strings.Split(strconv.FormatInt(n, 2), "")
	for i, b := range binaryStr {
		n, _ := strconv.ParseInt(b, 10, 64)
		nums[36-len(binaryStr)+i] = n
	}

	return nums
}

func binarySliceToNum(nums []int64) int64 {
	var str string
	for _, n := range nums {
		switch n {
		case 0:
			str += "0"
		case 1:
			str += "1"
		}
	}

	n, _ := strconv.ParseInt(str, 2, 64)
	return n
}

func (p *Program) storeMem(str string) {
	s := reMem.FindSubmatch([]byte(strings.TrimSpace(str)))
	if len(s) != 3 {
		return
	}
	addr, _ := strconv.ParseInt(string(s[1]), 10, 64)

	num, _ := strconv.ParseInt(string(s[2]), 10, 64)

	nums := numToBinarySlice(num)

	for k, v := range p.mask {
		nums[k] = v
	}

	p.mem[addr] = binarySliceToNum(nums)

}

func (p *Program) sum() (s int64) {
	for _, v := range p.mem {
		s += v
	}
	return
}

func main() {
	dat, err := ioutil.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")

	// Part1
	// p := NewProgram()
	// for _, line := range lines {
	// 	if reMask.Match([]byte(line)) {
	// 		p.updateMask(line)
	// 	} else {
	// 		p.storeMem(line)
	// 	}
	// }
	// fmt.Println(p.sum())

	p := NewProgram2()
	// p.updateMask("mask = 000000000000000000000000000000X1001X")
	// p.storeMem("mem[42] = 100")
	// p.updateMask("mask = 00000000000000000000000000000000X0XX")
	// p.storeMem("mem[26] = 1")
	// p.updateMask("mask = 0X10X1101111001X1X100X1X00011100XX11")
	// p.storeMem("mem[60365] = 24782")
	// fmt.Println(p.mem)
	for _, line := range lines {
		// fmt.Println(line)
		if reMask.Match([]byte(line)) {
			p.updateMask(line)
		} else {
			p.storeMem(line)
		}
	}
	fmt.Println(p.sum())
}
