package main

import (
	_ "embed"
	"fmt"
	"os"
	"sort"
	"strings"

	util "github.com/yhsiang/adventofcode"
)

//go:embed example
var example string

//go:embed input
var input string

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

// 2-length: 1
// 3-length: 7
// 4-length: 4
// 5-length: 2,3,5
// 6-length: 6,0,9
// 7-length: 8
// Algorithm
// 1. 1, 7, 4, 8
// 2. 6,0,9 -> contains 4: 9 -> contains 7: 0
// 3. 2,3,5 -> 6 contains: 5 -> contains 7: 3
func analyze(digits []string) map[string]int {
	var segments = make(map[int]string)
	var _235 []string
	var _609 []string
	for _, d := range digits {
		switch len(d) {
		case 2:
			segments[1] = d
		case 3:
			segments[7] = d
		case 4:
			segments[4] = d
		case 5:
			_235 = append(_235, d)
		case 6:
			_609 = append(_609, d)
		case 7:
			segments[8] = d
		}
	}

	for _, d := range _609 {
		if contains(d, segments[4]) && contains(d, segments[7]) {
			segments[9] = d
		} else if contains(d, segments[7]) {
			segments[0] = d
		} else {
			segments[6] = d
		}
	}

	for _, d := range _235 {
		if contains(segments[6], d) {
			segments[5] = d
		} else if contains(d, segments[7]) {
			segments[3] = d
		} else {
			segments[2] = d
		}

	}

	var seg = make(map[string]int)
	for k, v := range segments {
		seg[sortString(v)] = k
	}
	return seg
}

func contains(str string, compare string) bool {
	var maps = make(map[string]int)
	for _, s := range compare {
		maps[string(s)] = 0
	}
	for _, s := range str {
		if _, ok := maps[string(s)]; !ok {
			continue
		}
		maps[string(s)] = 1
	}
	var sum = 0
	for _, s := range maps {
		sum += s
	}

	return sum == len(compare)

}

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}
	lines := util.ByLine(file)
	var uniq int
	for _, line := range lines {
		digits := strings.Split(line, " | ")
		remain := strings.Split(digits[1], " ")
		for _, digit := range remain {
			if (len(digit) >= 2 && len(digit) <= 4) || len(digit) == 7 {
				uniq++
			}
		}
	}
	fmt.Printf("part1: %d\n", uniq)
	var data []string
	for _, line := range lines {
		digits := strings.Split(line, " | ")
		front := strings.Split(digits[0], " ")
		mappings := analyze(front)
		back := strings.Split(digits[1], " ")
		var fourDigits string
		for _, d := range back {
			fourDigits += fmt.Sprintf("%d", mappings[sortString(d)])
		}
		data = append(data, fourDigits)
	}

	fmt.Printf("part2: %d\n", util.SumInt(util.ToInt(data)))
}
