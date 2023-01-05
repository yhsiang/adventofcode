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

type Set map[string]int

func NewSet(data string) (Set, string, string) {
	var set = make(Set)
	var rows = strings.Split(data, "\n")
	// var cols = rows[0]
	start := ""
	goal := ""
	for i, d := range rows {
		for j, m := range d {
			c := fmt.Sprintf("%d,%d", i, j)
			switch m {
			case 'S':
				set[c] = 0
				start = c
			case 'E':
				set[c] = 27
				goal = c
			default:
				set[c] = int(m) - 96
			}

		}
	}
	return set, start, goal
}

var coords = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func neighbors(node string) []string {
	x, y := util.Coord(node)
	var output []string
	for _, c := range coords {
		coord := fmt.Sprintf("%d,%d", x+c[0], y+c[1])
		output = append(output, coord)
	}
	return output
}

type State struct {
	Node string
	Cost int
}

func createDist(s Set) Set {
	var newSet = make(Set)
	for k := range s {
		newSet[k] = int(^uint(0) >> 1)
	}
	return newSet
}

func (s Set) aStar(source string, goal string) int {
	var queue = []State{
		{source, 0},
	}
	var dist = createDist(s)
	dist[source] = 0
	for len(queue) > 0 {
		// should be lowest, min-heap will be better for O(logN)
		state := queue[0]
		queue = queue[1:]

		if state.Node == goal {
			return state.Cost
		}

		if state.Cost > dist[state.Node] {
			continue
		}

		for _, node := range neighbors(state.Node) {
			v, ok := s[node]
			if !ok || v > s[state.Node]+1 {
				continue
			}

			next := State{
				Node: node,
				Cost: state.Cost + 1, // steps
			}

			if next.Cost < dist[next.Node] {
				queue = append(queue, next)
				dist[next.Node] = next.Cost
			}
		}
	}

	return int(^uint(0) >> 1)
}

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}

	s, b, c := NewSet(file)
	cost := s.aStar(b, c)
	// first puzzle
	fmt.Printf("%d\n", cost)

	choices := []string{}
	for k, v := range s {
		if v == 1 || v == 0 {
			choices = append(choices, k)
		}
	}

	var costs = []int{}
	for _, choice := range choices {
		cost := s.aStar(choice, c)
		costs = append(costs, cost)
	}

	sort.Ints(costs)
	fmt.Printf("%+v\n", costs[0])

}
