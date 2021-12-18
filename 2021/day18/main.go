package main

import (
	_ "embed"
	"fmt"
	"math"
	"os"

	util "github.com/yhsiang/adventofcode"
)

//go:embed example
var example string

//go:embed input
var input string

// learn from https://getsturdy.com/advent-of-code-2021-uoeIDQk/changes/fb7d026d-d6be-4ec0-919b-8bf00a72fb94
type Pair struct {
	Value *int
	Left  *Pair
	Right *Pair
}

func parse(str string, pc int) (*Pair, int) {
	if pc > len(str) {
		return nil, -1
	}

	if str[pc] == '[' {
		l, npc := parse(str, pc+1)
		r, npc := parse(str, npc+1)
		return &Pair{Left: l, Right: r}, npc + 1
	}
	val, _ := util.Int(string(str[pc]))
	return &Pair{Value: &val}, pc + 1
}

var lastNum *Pair
var addToNext *int

func work(p *Pair, dep int, explode bool, split bool, transform bool) bool {
	if dep == 0 {
		lastNum = nil
		addToNext = nil
	}

	if explode && dep >= 4 && !transform && p.Value == nil && p.Left.Value != nil && p.Right.Value != nil {
		addToNext = p.Right.Value
		if lastNum != nil {
			*lastNum.Value += *p.Left.Value
		}
		i := 0
		*p = Pair{Value: &i}
		return true
	}

	if p.Value == nil {
		if work(p.Left, dep+1, explode, split, transform) {
			transform = true
		}
		if work(p.Right, dep+1, explode, split, transform) {
			transform = true
		}
		return transform
	}

	if addToNext != nil {
		*p.Value += *addToNext
		addToNext = nil
	}

	lastNum = p
	if split && !transform && *p.Value >= 10 {
		l := int(math.Floor(float64(*p.Value) / 2))
		r := int(math.Ceil(float64(*p.Value) / 2))
		*p = Pair{Left: &Pair{Value: &l}, Right: &Pair{Value: &r}}
		return true
	}
	return transform
}

func r(p *Pair) bool {
	if work(p, 0, true, false, false) {
		return true
	}
	return work(p, 0, false, true, false)
}

func reduce(p *Pair) *Pair {
	for {
		did := r(p)
		if !did {
			return p
		}
	}
}

func (p *Pair) print() string {
	var output string
	if p.Left != nil {
		output += "[" + p.Left.print()
	} else if p.Value != nil {
		output += fmt.Sprintf("%d", *p.Value)
	}
	if p.Right != nil {
		output += "," + p.Right.print() + "]"
	}
	return output

}

func mag(p *Pair) int {
	if p.Value != nil {
		return *p.Value
	}
	return mag(p.Left)*3 + mag(p.Right)*2
}

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}

	lines := util.ByLine(file)

	var prev *Pair
	for _, s := range lines {
		line, _ := parse(s, 0)
		if prev == nil {
			prev = line
		} else {
			prev = reduce(&Pair{Left: prev, Right: line})
		}
	}

	prev = reduce(prev)
	fmt.Printf("part1: %d\n", mag(prev))

	var largest int
	for _, v := range lines {
		for _, vv := range lines {
			l, _ := parse(v, 0)
			r, _ := parse(vv, 0)
			m := mag(reduce(&Pair{Left: l, Right: r}))
			largest = util.Max(largest, m)
		}
	}
	fmt.Printf("part2: %d\n", largest)

}
