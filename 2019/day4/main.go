package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"sort"
	"strings"

	util "github.com/yhsiang/adventofcode"
)

//go:embed example
var example string
var re = regexp.MustCompile(`1{2,}|2{2,}|3{2,}|4{2,}|5{2,}|6{2,}|7{2,}|8{2,}|9{2,}`)

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

func isPassword(input int, part2 bool) bool {
	s := fmt.Sprintf("%d", input)

	if sortString(s) != s {
		return false
	}

	if part2 {
		matches := re.FindAll([]byte(s), -1)
		var flag = false
		for _, match := range matches {
			if len(match) == 2 {
				flag = true
				break
			}
		}
		return flag

	} else {
		if !re.Match([]byte(s)) {
			return false
		}
	}

	return true
}

func main() {
	var file = example
	ranges := strings.Split(file, "-")
	from, _ := util.Int(ranges[0])
	to, _ := util.Int(ranges[1])

	var passwd []int
	for i := from; i <= to; i++ {
		if isPassword(i, false) {
			passwd = append(passwd, i)
		}
	}

	// fmt.Println(isPassword(111111, false))
	// fmt.Println(isPassword(223450, false))
	// fmt.Println(isPassword(123789, false))
	fmt.Printf("part1: %d\n", len(passwd))

	passwd = []int{}
	for i := from; i <= to; i++ {
		if isPassword(i, true) {
			passwd = append(passwd, i)
		}
	}

	fmt.Printf("part2: %d\n", len(passwd))

	// fmt.Println(isPassword(112233, true))
	// fmt.Println(isPassword(123444, true))
	// fmt.Println(isPassword(111122, true))
}
