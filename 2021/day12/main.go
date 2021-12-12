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

type Graph struct {
	Nodes map[string][]string
}

func initGraph(input string) *Graph {
	lines := util.ByLine(input)
	nodes := make(map[string][]string)
	for _, line := range lines {
		s := strings.Split(line, "-")
		if _, ok := nodes[s[0]]; !ok {
			nodes[s[0]] = []string{}
		}
		if _, ok := nodes[s[1]]; !ok {
			nodes[s[1]] = []string{}
		}
		nodes[s[0]] = append(nodes[s[0]], s[1])
		nodes[s[1]] = append(nodes[s[1]], s[0])

	}

	return &Graph{
		Nodes: nodes,
	}
}

type Set map[string]struct{}

func (s Set) copy() Set {
	a := make(Set)
	for k := range s {
		a[k] = struct{}{}
	}
	return a
}

func (g *Graph) paths(node string, visited Set, dup string, part2 bool) int {
	if node == "end" {
		return 1
	}

	if part2 && node == "start" && len(visited) > 0 {
		return 0
	}

	if _, ok := visited[node]; ok && strings.ToLower(node) == node {
		if part2 && dup == "" {
			dup = node
		} else {
			return 0
		}
	}
	newVisit := visited.copy()
	newVisit[node] = struct{}{}
	count := 0
	for _, dest := range g.Nodes[node] {
		count += g.paths(dest, newVisit, dup, part2)
	}
	return count
}

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}

	m := initGraph(file)
	var visited = make(Set)
	fmt.Printf("part1: %d\n", m.paths("start", visited, "", false))
	visited = make(Set)
	fmt.Printf("part2: %d\n", m.paths("start", visited, "", true))
}
