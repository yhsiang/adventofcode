package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Rule struct {
	Character string
	Min       int
	Max       int
}

func NewRule(str string) (*Rule, error) {
	d1 := strings.Split(str, " ")
	d2 := strings.Split(d1[0], "-")
	min, err := strconv.ParseInt(d2[0], 10, 64)
	if err != nil {
		return nil, err
	}

	max, err := strconv.ParseInt(d2[1], 10, 64)
	if err != nil {
		return nil, err
	}

	return &Rule{
		Character: d1[1],
		Min:       int(min),
		Max:       int(max),
	}, nil
}

func (r *Rule) Match(str string) bool {
	var mappings = make(map[string]int, 26)
	chars := strings.Split(str, "")
	for _, char := range chars {
		if _, ok := mappings[char]; !ok {
			mappings[char] = 0
		}
		mappings[char] += 1
	}

	if v, ok := mappings[r.Character]; ok {
		return v >= r.Min && v <= r.Max
	}

	return false
}

type RuleP struct {
	Character string
	Pos1      int
	Pos2      int
}

func NewRuleP(str string) (*RuleP, error) {
	d1 := strings.Split(str, " ")
	d2 := strings.Split(d1[0], "-")
	min, err := strconv.ParseInt(d2[0], 10, 64)
	if err != nil {
		return nil, err
	}

	max, err := strconv.ParseInt(d2[1], 10, 64)
	if err != nil {
		return nil, err
	}

	return &RuleP{
		Character: d1[1],
		Pos1:      int(min) - 1,
		Pos2:      int(max) - 1,
	}, nil
}

func (r *RuleP) Match(str string) bool {
	chars := strings.Split(strings.TrimSpace(str), "")
	if chars[r.Pos1] == r.Character && chars[r.Pos2] != r.Character {
		return true
	}

	if chars[r.Pos1] != r.Character && chars[r.Pos2] == r.Character {
		return true
	}

	return false
}

func main() {
	// a, err := NewRuleP("10-12 d")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(a.Match("ddvddnmdnlvdddqdcqph"))
	dat, err := ioutil.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")

	var valid = 0
	for _, line := range lines {
		r := strings.Split(line, ":")
		rule, err := NewRuleP(r[0])
		if err != nil {
			continue
		}

		b := rule.Match(strings.TrimSpace(r[1]))
		// fmt.Println(b, r[1])

		if b {
			valid += 1
		}
	}
	fmt.Println(valid)
}
