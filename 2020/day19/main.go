package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

var reRule = regexp.MustCompile(`(\d+)\:(.+)`)

type Rules map[string]string

func parseRule(str string) Rules {
	rules := make(Rules)
	for _, s := range strings.Split(str, "\n") {
		if ss := reRule.FindSubmatch([]byte(s)); len(ss) == 3 {
			rules[string(ss[1])] = strings.TrimSpace(strings.ReplaceAll(string(ss[2]), `"`, ""))
		}
	}
	return rules
}

func part1(idx string, rules Rules) string {
	rule := rules[idx]
	var b []string
	for _, s := range strings.Split(rule, "|") {
		var a []string
		for _, ss := range strings.Split(strings.TrimSpace(s), " ") {
			// fmt.Println(ss)
			if ss == "a" || ss == "b" {
				a = append(a, ss)
			} else {
				a = append(a, part1(ss, rules))
			}
		}
		b = append(b, strings.Join(a, ""))
	}

	c := strings.Join(b, "|")

	matched, _ := regexp.Match(`\|`, []byte(c))

	if !matched {
		return c
	}
	return fmt.Sprintf("(?:%s)", c)

}
func part2(idx string, rules Rules) string {
	rule := rules[idx]
	var b []string
	if idx == "31" || idx == "42" || idx == "8" || idx == "11" {
		return rule
	}
	for _, s := range strings.Split(rule, "|") {
		var a []string
		for _, ss := range strings.Split(strings.TrimSpace(s), " ") {
			// fmt.Println(ss)
			if ss == "a" || ss == "b" {
				a = append(a, ss)
			} else {
				a = append(a, part2(ss, rules))
			}
		}
		b = append(b, strings.Join(a, ""))
	}

	c := strings.Join(b, "|")

	matched, _ := regexp.Match(`\|`, []byte(c))

	if !matched {
		return c
	}
	return fmt.Sprintf("(?:%s)", c)

}

func main() {
	dat, err := ioutil.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n\n")
	rules := parseRule(lines[0])
	// part1
	// r0 := regexp.MustCompile(fmt.Sprintf("^%s$", part1("0", rules)))
	// var sum = 0
	// for _, msg := range strings.Split(lines[1], "\n") {
	// 	if r0.Match([]byte(msg)) {
	// 		sum += 1
	// 	}
	// }
	// fmt.Println(sum)

	// part2
	// rules["8"] = `42 | 42 8`
	// rules["11"] = `42 31 | 42 11 31`
	rules["31"] = part1("31", rules)
	rules["42"] = part1("42", rules)
	rules["8"] = fmt.Sprintf("%s+", rules["42"])
	var s []string
	for i := 1; i < 5; i++ {
		s = append(s, fmt.Sprintf("%s{%d}%s{%d}", rules["42"], i, rules["31"], i))
	}
	rules["11"] = fmt.Sprintf("(?:%s)", strings.Join(s, "|"))
	// fmt.Print(part2("0", rules))
	r0 := regexp.MustCompile(fmt.Sprintf("^%s$", part2("0", rules)))
	var sum = 0
	for _, msg := range strings.Split(lines[1], "\n") {
		if r0.Match([]byte(msg)) {
			sum += 1
		}
	}
	fmt.Println(sum)
}
