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

func expectedSymbol(symbol string) string {
	switch symbol {
	case "}":
		return "{"
	case ")":
		return "("
	case ">":
		return "<"
	case "]":
		return "["
	case "[":
		return "]"
	case "<":
		return ">"
	case "(":
		return ")"
	case "{":
		return "}"
	}
	return ""
}

func (s *SyntaxChecker) add(symbol string) (string, error) {
	switch symbol {
	case "{", "(", "[", "<":
		s.Stack = append(s.Stack, symbol)
	case "}", ")", "]", ">":
		lastIdx := len(s.Stack) - 1
		last := s.Stack[lastIdx]
		expected := expectedSymbol(symbol)
		if expected == last {
			s.Stack = s.Stack[:lastIdx]
		} else {
			return symbol, fmt.Errorf("expected %s, but found %s instead", expectedSymbol(last), symbol)
		}
	}
	return "", nil
}

func (s *SyntaxChecker) add2(symbol string) {
	switch symbol {
	case "{", "(", "[", "<":
		s.Stack = append(s.Stack, symbol)
	case "}", ")", "]", ">":
		lastIdx := len(s.Stack) - 1
		last := s.Stack[lastIdx]
		expected := expectedSymbol(symbol)
		if expected == last {
			s.Stack = s.Stack[:lastIdx]
		}
	}
}

var part1Score = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

func (s *SyntaxChecker) check(line string) string {
	for _, symbol := range strings.Split(line, "") {
		errSymbol, err := s.add(symbol)
		if err != nil {
			// errSymbols = append(errSymbols, errSymbol)
			return errSymbol
		}
	}
	return ""
}

func (s *SyntaxChecker) check2(line string) []string {
	for _, symbol := range strings.Split(line, "") {
		s.add2(symbol)
	}
	var completed []string
	// for _, s := range s.Stack {
	for i := len(s.Stack) - 1; i >= 0; i-- {
		completed = append(completed, expectedSymbol(s.Stack[i]))
	}

	return completed
}

func (s *SyntaxChecker) reset() {
	s.Stack = []string{}
}

var part2Score = map[string]int{
	")": 1,
	"]": 2,
	"}": 3,
	">": 4,
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

	checker := &SyntaxChecker{}
	var errSymbols []string
	var remains []string
	for _, line := range lines {
		symbol := checker.check(line)
		if symbol != "" {
			errSymbols = append(errSymbols, symbol)
		} else {
			remains = append(remains, line)
		}
		checker.reset()
	}

	var sum = 0
	for _, s := range errSymbols {
		sum += part1Score[s]
	}
	fmt.Printf("part1: %d\n", sum)
	var scores []int
	for _, r := range remains {
		scores = append(scores, count(checker.check2(r)))
		checker.reset()
	}
	sort.Ints(scores)
	fmt.Printf("part2: %d\n", scores[len(scores)/2])

}
