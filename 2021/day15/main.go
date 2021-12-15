package main

import (
	_ "embed"
	"fmt"
	"os"
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
	var cols = rows[0]
	for i, d := range rows {
		for j, m := range d {
			c := fmt.Sprintf("%d,%d", i, j)
			n, _ := util.Int(string(m))
			set[c] = n
		}
	}
	return set, "0,0", fmt.Sprintf("%d,%d", len(rows)-1, len(cols)-1)
}

func ExpandedSet(data string) (Set, string, string) {
	var set = make(Set)
	var maze [][]int
	var rows = strings.Split(data, "\n")
	var cols = rows[0]
	for _, row := range rows {
		maze = append(maze, util.ToInt(strings.Split(row, "")))
	}

	var rowsNum = 5 * len(rows)
	var colsNum = 5 * len(cols)
	for i := 0; i < rowsNum; i++ {
		for j := 0; j < colsNum; j++ {
			c := fmt.Sprintf("%d,%d", i, j)
			x := i / len(rows)
			y := j / len(cols)
			cost := maze[i%len(rows)][j%len(cols)] + x + y
			if cost > 9 {
				cost -= 9
			}
			set[c] = cost
		}
	}

	return set, "0,0", fmt.Sprintf("%d,%d", rowsNum-1, colsNum-1)
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
	// fmt.Println(output)
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

func (s Set) dijkstra(source string, goal string) int {
	var queue = []State{
		{source, 0},
	}
	var dist = createDist(s)
	var cost int
	for {
		state := queue[0]
		queue = queue[1:]

		if state.Node == goal && len(queue) == 0 {
			cost = state.Cost
			break
		}

		if state.Cost > dist[state.Node] {
			continue
		}

		for _, node := range neighbors(state.Node) {
			v, ok := s[node]
			if !ok {
				continue
			}
			next := State{
				Node: node,
				Cost: state.Cost + v,
			}
			if next.Cost < dist[next.Node] {
				queue = append(queue, next)
				dist[next.Node] = next.Cost
			}
		}
		// fmt.Println(dist)
	}

	return cost
}

func print(s Set, row, col int) {
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			coord := fmt.Sprintf("%d,%d", i, j)
			fmt.Printf("%d", s[coord])
		}
		fmt.Println()
	}
}

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}

	s, start, end := NewSet(file)
	fmt.Printf("part1: %d\n", s.dijkstra(start, end))
	s, start, end = ExpandedSet(file)
	// print(s, 50, 50)
	fmt.Printf("part2: %d\n", s.dijkstra(start, end))
}
