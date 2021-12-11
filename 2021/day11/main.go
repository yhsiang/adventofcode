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

type Octopuses struct {
	Energies map[string]int
	Rows     int
	Cols     int
}

func initOctopuses(file string) *Octopuses {
	lines := util.ByLine(file)
	var energies = make(map[string]int)
	for i, l := range lines {
		for j, v := range strings.Split(l, "") {
			n, _ := util.Int64(v)
			energies[fmt.Sprintf("%d,%d", i, j)] = int(n)
		}
	}

	return &Octopuses{
		Energies: energies,
		Rows:     len(lines),
		Cols:     len(lines[0]),
	}
}

var coords = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
	{1, 1},
	{-1, 1},
	{-1, -1},
	{1, -1},
}

func (m *Octopuses) flash(x int, y int, flashed map[string]struct{}) {
	for _, c := range coords {
		coord := fmt.Sprintf("%d,%d", x+c[0], y+c[1])
		if _, ok := m.Energies[coord]; ok {
			m.Energies[coord] += 1
			if m.Energies[coord] > 9 {
				if _, ok := flashed[coord]; !ok {
					flashed[coord] = struct{}{}
					m.flash(x+c[0], y+c[1], flashed)
				}
			}
		}
	}
}

func (m *Octopuses) model(steps int, part2 bool) int {
	var flashes int
	var i int
	for {
		i++
		var flashed = make(map[string]struct{})
		for coord := range m.Energies {
			r, c := util.Coord(coord)
			m.Energies[coord] += 1
			if m.Energies[coord] > 9 {
				if _, ok := flashed[coord]; !ok {
					flashed[coord] = struct{}{}
					m.flash(r, c, flashed)
				}
			}
		}
		for k := range flashed {
			m.Energies[k] = 0
		}
		flashes += len(flashed)
		if part2 && len(flashed) == len(m.Energies) {
			break
		}
		if !part2 && i == steps {
			break
		}
	}
	if part2 {
		return i
	}
	return flashes
}

func (m *Octopuses) print() {
	for r := 0; r < m.Rows; r++ {
		var s string
		for c := 0; c < m.Cols; c++ {
			coord := fmt.Sprintf("%d,%d", r, c)
			s += fmt.Sprintf("%d", m.Energies[coord])

		}
		fmt.Println(s)
	}
	fmt.Println()
}

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}

	m := initOctopuses(file)
	fmt.Printf("part1: %d\n", m.model(100, false))
	m = initOctopuses(file)
	fmt.Printf("part2: %d\n", m.model(100, true))
}
