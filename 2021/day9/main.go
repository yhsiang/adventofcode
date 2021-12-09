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
	Numbers []int
	HashMap map[int]int
	rowNum  int
	colNum  int
}

func initMap(input string) *Map {
	data := util.ByLine(input)
	var output []int
	var hashmap = make(map[int]int)
	for _, d := range data {
		output = append(output, util.ToInt(strings.Split(d, ""))...)
	}
	for i, d := range output {
		hashmap[i] = d
	}
	return &Map{
		Numbers: output,
		HashMap: hashmap,
		rowNum:  len(data[0]),
		colNum:  len(data),
	}
}

func (m *Map) findLowPoints() (int, []int) {
	// -1, -1 * rowNum, 1, 1* rowNum
	w := -1
	n := -1 * m.rowNum
	e := 1
	s := 1 * m.rowNum
	var lowPoints []int
	var riskLevels = 0
	for i, d := range m.Numbers {
		wc := w + i
		nc := n + i
		ec := e + i
		sc := s + i
		if wc >= 0 && wc < len(m.Numbers) && d > m.Numbers[wc] {
			continue
		}
		if nc >= 0 && nc < len(m.Numbers) && d > m.Numbers[nc] {
			continue
		}
		if ec >= 0 && ec < len(m.Numbers) && d > m.Numbers[ec] {
			continue
		}
		if sc >= 0 && sc < len(m.Numbers) && d > m.Numbers[sc] {
			continue
		}
		if d != 9 {
			riskLevels += 1 + d
			lowPoints = append(lowPoints, i)
		}
	}

	return riskLevels, lowPoints
}

func exist(input map[int]struct{}, value int) bool {
	_, ok := input[value]
	return ok
}

func (m *Map) checkPoints(index int, finded map[int]struct{}) map[int]struct{} {
	output := make(map[int]struct{})
	var coords = []int{
		-1,
		-1 * m.rowNum,
		1,
		1 * m.rowNum,
	}
	switch {
	case index >= 0 && index <= m.rowNum-1:
		coords = []int{
			-1,
			1,
			1 * m.rowNum,
		}
	case index >= (m.colNum-1)*m.rowNum && index <= m.colNum*m.rowNum-1:
		coords = []int{
			-1,
			1,
			-1 * m.rowNum,
		}
	case index%m.rowNum == 0:
		coords = []int{
			-1 * m.rowNum,
			1,
			1 * m.rowNum,
		}
	case index%m.rowNum == m.rowNum-1:
		coords = []int{
			-1,
			-1 * m.rowNum,
			1 * m.rowNum,
		}
	}

	for _, c := range coords {
		coord := c + index
		if !exist(finded, coord) && coord >= 0 && coord < len(m.Numbers) {
			output[coord] = struct{}{}
		}
	}

	return output
}

func (m *Map) findBasins(lowPoints []int) int {
	var basins = make(map[int]map[int]struct{})
	var finded = make(map[int]struct{})
	for _, p := range lowPoints {
		var basin = make(map[int]struct{})
		basin[p] = struct{}{}
		finded[p] = struct{}{}
		points := m.checkPoints(p, finded)
		for len(points) > 0 {
			var temp = make(map[int]struct{})
			for d := range points {
				finded[d] = struct{}{}
				if v, ok := m.HashMap[d]; ok && v != 9 {
					basin[d] = struct{}{}
					for k := range m.checkPoints(d, finded) {
						temp[k] = struct{}{}
					}
				}
			}
			points = temp
		}
		basins[p] = basin
	}

	var basinLens []int
	for _, v := range basins {
		basinLens = append(basinLens, len(v))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(basinLens)))
	return util.MultiplyInt(basinLens[0:3])
}

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}

	m := initMap(file)
	riskLevel, points := m.findLowPoints()
	fmt.Printf("part1: %d\n", riskLevel)
	fmt.Printf("part2: %d\n", m.findBasins(points))

}
