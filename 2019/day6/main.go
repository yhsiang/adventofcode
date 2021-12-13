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

type OrbitMap struct {
	Connections map[string]map[string]struct{}
	Orbits      map[string]int
}

func initOrbitMap(file string) *OrbitMap {
	var connections = make(map[string]map[string]struct{})
	for _, line := range util.ByLine(file) {
		l := strings.Split(line, ")")
		if _, ok := connections[l[0]]; !ok {
			connections[l[0]] = make(map[string]struct{})
		}
		connections[l[0]][l[1]] = struct{}{}
	}

	return &OrbitMap{
		Connections: connections,
		Orbits:      make(map[string]int),
	}
}

func initOrbitMapByDirection(file string) *OrbitMap {
	var connections = make(map[string]map[string]struct{})
	for _, line := range util.ByLine(file) {
		l := strings.Split(line, ")")
		if _, ok := connections[l[0]]; !ok {
			connections[l[0]] = make(map[string]struct{})
		}
		if _, ok := connections[l[1]]; !ok {
			connections[l[1]] = make(map[string]struct{})
		}
		connections[l[0]][l[1]] = struct{}{}
		connections[l[1]][l[0]] = struct{}{}
	}

	return &OrbitMap{
		Connections: connections,
		Orbits:      make(map[string]int),
	}
}

func (o *OrbitMap) count(node string, seen int) {
	nodes, ok := o.Connections[node]
	o.Orbits[node] = seen
	if !ok {
		return
	}

	for n := range nodes {
		o.count(n, seen+1)
	}
}

func (o *OrbitMap) walk(start string, end string, walked map[string]struct{}) int {
	// fmt.Println(start, end, walked)
	nodes, ok := o.Connections[start]
	if !ok {
		return 0
	}
	var newWalked = make(map[string]struct{})
	for k := range walked {
		newWalked[k] = struct{}{}
	}
	newWalked[start] = struct{}{}

	if start == end {
		return len(walked) - 2
	}

	count := 0
	for node := range nodes {
		if _, ok := walked[node]; !ok {
			count += o.walk(node, end, newWalked)
		}
	}
	return count
}

func (o *OrbitMap) sum() int {
	var sum = 0
	for _, v := range o.Orbits {
		sum += v
	}
	return sum
}

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}

	m := initOrbitMap(file)

	m.count("COM", 0)
	fmt.Printf("part1: %d\n", m.sum())

	m = initOrbitMapByDirection(file)
	var walked = make(map[string]struct{})
	fmt.Printf("part2: %d\n", m.walk("YOU", "SAN", walked))

}
