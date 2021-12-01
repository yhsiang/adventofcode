package main

import (
	"math"
	"strconv"
	"strings"
)

type Program2 struct {
	mask []string
	mem  map[int]int64
}

func NewProgram2() *Program2 {
	return &Program2{
		mem: make(map[int]int64),
	}
}

func (p *Program2) updateMask(str string) {
	s := reMask.FindSubmatch([]byte(strings.TrimSpace(str)))
	if len(s) != 2 {
		return
	}

	p.mask = strings.Split(string(s[1]), "")
}

func numToBinaryStr(n int64, paddLen int) []string {
	var str = make([]string, paddLen)
	for i := range str {
		str[i] = "0"
	}
	binaryStr := strings.Split(strconv.FormatInt(n, 2), "")
	for i, b := range binaryStr {
		str[paddLen-len(binaryStr)+i] = b
	}

	return str
}

func (p *Program2) storeMem(str string) {
	s := reMem.FindSubmatch([]byte(strings.TrimSpace(str)))
	if len(s) != 3 {
		return
	}
	addr, _ := strconv.ParseInt(string(s[1]), 10, 64)
	strs := numToBinaryStr(addr, 36)
	var xIndex []int
	for i := range strs {
		switch p.mask[i] {
		case "0":
		case "1":
			strs[i] = "1"
		case "X":
			strs[i] = "0"
			xIndex = append(xIndex, i)
		}
	}

	value, _ := strconv.ParseInt(string(s[2]), 10, 64)
	p.calAddrs(strs, xIndex, value)
}

func (p *Program2) calAddrs(strs []string, xIndex []int, value int64) {
	var length = len(xIndex)
	var times = int(math.Pow(2, float64(length)))
	var copied = make([]string, len(strs))
	copy(copied, strs)
	for i := 0; i < times; i++ {
		copy(copied, strs)
		for i, s := range numToBinaryStr(int64(i), length) {
			copied[xIndex[i]] = s
		}
		n, _ := strconv.ParseInt(strings.Join(copied, ""), 2, 64)
		p.mem[int(n)] = value
	}

}
func (p *Program2) sum() (s int64) {
	for _, v := range p.mem {
		s += v
	}
	return
}
