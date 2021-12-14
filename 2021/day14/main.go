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

type Rules map[string]string

func NewRules(data string) Rules {
	var rules = make(Rules)
	for _, r := range util.ByLine(data) {
		output := strings.Split(r, " -> ")
		rules[output[0]] = output[1]
	}
	return rules
}

func insert(polymer string, rules Rules) string {
	s := 0
	e := 2
	var output string
	for s < len(polymer)-1 {
		// fmt.Println(s, e, polymer[s:e])
		n := rules[polymer[s:e]]
		if s == 0 {
			output += fmt.Sprintf("%s%s%s",
				string(polymer[s]), n, string(polymer[e-1]))
		} else {
			output += fmt.Sprintf("%s%s",
				n, string(polymer[e-1]))
		}
		s++
		e++
	}
	return output
}

func NewPairs(polymer string) (map[string]int, map[string]int) {
	pairs := make(map[string]int)
	elems := make(map[string]int)

	for i := 0; i < len(polymer)-1; i++ {
		elems[string(polymer[i])] += 1
		pairs[polymer[i:i+2]] += 1
	}
	elems[string(polymer[len(polymer)-1])] += 1
	return pairs, elems
}

type Set map[string]int

func (s Set) copy() Set {
	var n = make(Set)
	for k, v := range s {
		n[k] = v
	}
	return n
}

func fastInsert(pairs, elems Set, rules Rules) (map[string]int, map[string]int) {
	var newPairs = pairs.copy()
	var newElems = elems.copy()
	for pair, count := range pairs {
		// delete(newPairs, pair)
		newPairs[pair] -= count
		insert := rules[pair]
		// fmt.Println(insert)
		newElems[insert] += count
		newPairs[string(pair[0])+insert] += count
		newPairs[insert+string(pair[1])] += count

	}
	// fmt.Println(newPairs, newElems)
	return newPairs, newElems
}

func count(data string) int {
	var counts = make(map[string]int)
	for _, d := range strings.Split(data, "") {
		if _, ok := counts[d]; !ok {
			counts[d] = 0
		}
		counts[d] += 1
	}
	var countInts []int
	for _, v := range counts {
		countInts = append(countInts, v)
	}

	sort.Ints(countInts)

	return countInts[len(countInts)-1] - countInts[0]
}

func countElems(elems map[string]int) int {
	var countInts []int
	for _, v := range elems {
		countInts = append(countInts, v)
	}

	sort.Ints(countInts)
	return countInts[len(countInts)-1] - countInts[0]
}

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}

	inputs := strings.Split(file, "\n\n")
	polymer := inputs[0]
	rules := NewRules(inputs[1])
	output := polymer
	for i := 0; i < 10; i++ {
		output = insert(output, rules)
	}

	fmt.Printf("part1: %d\n", count(output))
	output = polymer
	pairs, elems := NewPairs(polymer)
	for i := 0; i < 40; i++ {
		pairs, elems = fastInsert(pairs, elems, rules)
	}

	fmt.Printf("part2: %d\n", countElems(elems))
}
