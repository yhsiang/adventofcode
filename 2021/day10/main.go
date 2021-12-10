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

type SyntaxChecker struct {
	Stack []string
}

var part1Score = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

var part2Score = map[string]int{
	")": 1,
	"]": 2,
	"}": 3,
	">": 4,
}

var pairSymbols = map[string]string{
	")": "(",
	"]": "[",
	"}": "{",
	">": "<",
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
}

func (s *SyntaxChecker) add(symbol string) (string, error) {
	switch symbol {
	case "{", "(", "[", "<":
		s.Stack = append(s.Stack, symbol)
	case "}", ")", "]", ">":
		lastIdx := len(s.Stack) - 1
		last := s.Stack[lastIdx]
		expected := pairSymbols[symbol]
		if expected == last {
			s.Stack = s.Stack[:lastIdx]
		} else {
			return symbol, fmt.Errorf("expected %s, but found %s instead", pairSymbols[last], symbol)
		}
	}
	return "", nil
}

func (s *SyntaxChecker) check(line string) (string, []string) {
	for _, symbol := range strings.Split(line, "") {
		errSymbol, err := s.add(symbol)
		if err != nil {
			return errSymbol, nil
		}
	}

	var completed []string
	for i := len(s.Stack) - 1; i >= 0; i-- {
		completed = append(completed, pairSymbols[s.Stack[i]])
	}
	return "", completed
}

func count(input []string) int {
	var total = 0
	for _, i := range input {
		total = total*5 + part2Score[i]
	}
	return total
}

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}

	lines := util.ByLine(file)

	var part1Scores []int
	var part2Scores []int
	for _, line := range lines {
		checker := &SyntaxChecker{}
		s1, s2 := checker.check(line)
		if s1 != "" {
			part1Scores = append(part1Scores, part1Score[s1])
		} else {
			part2Scores = append(part2Scores, count(s2))
		}
	}

	fmt.Printf("part1: %d\n", util.SumInt(part1Scores))
	sort.Ints(part2Scores)
	fmt.Printf("part2: %d\n", part2Scores[len(part2Scores)/2])

}
