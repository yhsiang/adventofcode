package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var reRule = regexp.MustCompile(`^([\w\s]+): (\d+)-(\d+) or (\d+)-(\d+)$`)
var reDeparture = regexp.MustCompile(`^departure \w+$`)

type Rule struct {
	Name string
	Min  []int64
	Max  []int64
}

type RuleSlice []Rule

func NewRule(str string) RuleSlice {
	strs := strings.Split(str, "\n")
	var rules []Rule
	for _, rule := range strs {
		var r = Rule{}
		matches := reRule.FindSubmatch([]byte(strings.TrimSpace(rule)))
		if len(matches) != 6 {
			return nil
		}
		r.Name = string(matches[1])
		for _, value := range matches[2:4] {
			i, _ := strconv.ParseInt(string(value), 10, 64)
			r.Min = append(r.Min, i)
		}
		for _, value := range matches[4:6] {
			i, _ := strconv.ParseInt(string(value), 10, 64)
			r.Max = append(r.Max, i)
		}
		rules = append(rules, r)
	}

	return rules
}

func (r Rule) Match(i int64) bool {
	if (i >= r.Min[0] && i <= r.Min[1]) || (i >= r.Max[0] && i <= r.Max[1]) {
		return true
	}
	return false
}

func (rs RuleSlice) flat() [][]int64 {
	var i [][]int64
	for _, r := range rs {
		i = append(i, r.Min)
		i = append(i, r.Max)
	}
	return i
}

func check(i int64, numbers [][]int64) bool {
	for _, n := range numbers {
		if i >= n[0] && i <= n[1] {
			return true
		}
	}
	return false
}

func verify(str string, numbers [][]int64) (invalid []int64) {
	tickets := strings.Split(str, "\n")
	if tickets[0] != "nearby tickets:" {
		return
	}
	for _, ticket := range tickets[1:] {
		for _, t := range strings.Split(ticket, ",") {
			i, _ := strconv.ParseInt(t, 10, 64)
			if !check(i, numbers) {
				invalid = append(invalid, i)
			}
		}
	}

	return
}

func verify2(str string, numbers [][]int64) (valids []string) {
	tickets := strings.Split(str, "\n")
	if tickets[0] != "nearby tickets:" {
		return
	}
	for _, ticket := range tickets[1:] {
		var valid = true
		for _, t := range strings.Split(ticket, ",") {
			i, _ := strconv.ParseInt(t, 10, 64)
			if !check(i, numbers) {
				valid = false
			}
		}
		if valid {
			valids = append(valids, ticket)
		}
	}

	return
}

func sum(numbers []int64) (s int64) {
	for _, v := range numbers {
		s += v
	}
	return s
}

func match(numbers []int64, rs RuleSlice) []Rule {
	var rules []Rule
	for _, r := range rs {
		var n int
		for _, num := range numbers {
			if r.Match(num) {
				n++
			}
		}
		if n == len(numbers) {
			rules = append(rules, r)
		}
	}

	return rules
}

func combine(tickets []string) [][]int64 {
	var combinations = make([][]int64, len(strings.Split(strings.TrimSpace(tickets[0]), ",")))
	for _, ticket := range tickets {
		for i, t := range strings.Split(strings.TrimSpace(ticket), ",") {
			n, _ := strconv.ParseInt(t, 10, 64)

			combinations[i] = append(combinations[i], n)
		}
	}
	return combinations
}

func main() {
	dat, err := ioutil.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n\n")
	rs := NewRule(lines[0])

	numbers := rs.flat()

	// Part1
	// result := verify(lines[2], numbers)
	// fmt.Println(sum(result))

	tickets := verify2(lines[2], numbers)
	// fmt.Println(combine([]string{
	// 	"3,9,18",
	// 	"15,1,5",
	// 	"5,14,9",
	// 	"4,1,2",
	// }))

	reversed := combine(tickets)
	var visited = make(map[string][]int)
	for i, r := range reversed {
		rules := match(r, rs)
		for _, a := range rules {
			visited[a.Name] = append(visited[a.Name], i)
		}
	}

	var indexes [][]int
	for _, v := range visited {
		indexes = append(indexes, v)
	}

	sort.Slice(indexes, func(i, j int) bool {
		return len(indexes[i]) < len(indexes[j])
	})

	var fields = make(map[int]string)
	for _, indexes := range indexes {
		for k, v := range visited {
			if len(v) == len(indexes) {
				for _, index := range indexes {
					if _, ok := fields[index]; !ok {
						fields[index] = k
					}
				}
			}
		}
	}
	var targetIndex []int
	for k, f := range fields {
		if reDeparture.Match([]byte(f)) {
			fmt.Println(k, f)
			targetIndex = append(targetIndex, k)
		}
	}

	myTicket := strings.Split(lines[1], "\n")[1]
	num := strings.Split(myTicket, ",")
	var product int64 = 1
	for _, i := range targetIndex {
		k, _ := strconv.ParseInt(num[i], 10, 64)
		product *= k
	}
	fmt.Println(product)
}
