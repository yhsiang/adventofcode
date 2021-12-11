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

type Map struct {
	HashMap map[string]int
	Rows    int
	Cols    int
}

func initMap(input string) *Map {
	lines := util.ByLine(input)
	var maps = make(map[string]int)
	for i, l := range lines {
		for j, v := range strings.Split(l, "") {
			n, _ := util.Int64(v)
			maps[fmt.Sprintf("%d,%d", i, j)] = int(n)
		}
	}

	return &Map{
		HashMap: maps,
		Rows:    len(lines),
		Cols:    len(lines[0]),
	}
}

func (m *Map) print() {
	for r := 0; r < m.Rows; r++ {
		var s string
		for c := 0; c < m.Cols; c++ {
			coord := fmt.Sprintf("%d,%d", r, c)
			s += fmt.Sprintf("%d", m.HashMap[coord])

		}
		fmt.Println(s)
	}
	fmt.Println()
}

var coords = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func (m *Map) lowPoints() (int, [][]int) {
	var risk int
	var points [][]int
	for r := 0; r < m.Rows; r++ {
		for c := 0; c < m.Cols; c++ {
			current := m.HashMap[fmt.Sprintf("%d,%d", r, c)]
			var low = true
			for _, move := range coords {
				coord := fmt.Sprintf("%d,%d", r+move[0], c+move[1])
				v, ok := m.HashMap[coord]
				if !ok {
					continue
				}
				if current > v {
					low = false
					break
				}
			}
			if low && current != 9 {
				risk += 1 + current
				points = append(points, []int{r, c})
			}
		}
	}

	return risk, points
}

func (m *Map) check(point []int, finded map[string]struct{}) (output [][]int) {
	for _, c := range coords {
		x, y := point[0]+c[0], point[1]+c[1]
		coord := fmt.Sprintf("%d,%d", x, y)
		v, ok := m.HashMap[coord]
		_, ok2 := finded[coord]
		if !ok2 && ok && v != 9 {
			finded[coord] = struct{}{}
			output = append(output, []int{x, y})
		}
	}

	return
}

func (m *Map) basins(points [][]int) int {
	var basins = make(map[string]map[string]struct{})
	for _, p := range points {
		var basin = make(map[string]struct{})
		var finded = make(map[string]struct{})
		coord := fmt.Sprintf("%d,%d", p[0], p[1])
		basin[coord] = struct{}{}
		finded[coord] = struct{}{}
		points := m.check(p, finded)
		for len(points) > 0 {
			target := points[0]
			points = append(points[1:], m.check(target, finded)...)
		}
		basins[coord] = finded
	}

	var sizes []int
	for _, b := range basins {
		sizes = append(sizes, len(b))
	}

	sort.Ints(sizes)
	n := len(sizes)
	return sizes[n-1] * sizes[n-2] * sizes[n-3]
}

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}

	m := initMap(file)
	riskLevel, points := m.lowPoints()
	fmt.Printf("part1: %d\n", riskLevel)
	fmt.Printf("part2: %d\n", m.basins(points))
}
