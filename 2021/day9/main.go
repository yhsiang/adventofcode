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
	rowNum  int
	colNum  int
}

func initMap(input string) *Map {
	data := util.ByLine(input)
	var output []int
	for _, d := range data {
		output = append(output, util.ToInt(strings.Split(d, ""))...)
	}
	return &Map{
		Numbers: output,
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

func exist(input []int, value int) bool {
	for _, i := range input {
		if i == value {
			return true
		}
	}
	return false
}

func (m *Map) checkPoints(index int, finded []int) (output []int) {
	var coords = []int{
		-1,
		-1 * m.rowNum,
		1,
		1 * m.rowNum,
	}

	if index == 0 {
		coords = []int{
			1,
			1 * m.rowNum,
		}
	} else if index == m.rowNum-1 {

		coords = []int{
			-1,
			1 * m.rowNum,
		}
	} else if index == (m.colNum-1)*m.rowNum {
		coords = []int{
			1,
			-1 * m.rowNum,
		}
	} else if index == (m.colNum*m.rowNum)-1 {
		coords = []int{
			-1,
			-1 * m.rowNum,
		}
	} else if index > 0 && index < m.rowNum {
		coords = []int{
			-1,
			1,
			1 * m.rowNum,
		}
	} else if index > 0 && index%m.rowNum == 0 {
		coords = []int{
			-1 * m.rowNum,
			1,
			1 * m.rowNum,
		}
	} else if index > 0 && index%m.rowNum == m.rowNum-1 {
		coords = []int{
			-1,
			-1 * m.rowNum,
			1 * m.rowNum,
		}
	} else if index > (m.colNum-1)*m.rowNum && index < (m.colNum*m.rowNum) {
		coords = []int{
			-1,
			1,
			-1 * m.rowNum,
		}
	}

	for _, c := range coords {
		coord := c + index
		if !exist(finded, coord) {
			output = append(output, c+index)
		}
	}
	return
}

func (m *Map) findBasins(lowPoints []int) int {
	var basins = make(map[int][]int)
	var finded []int
	for _, p := range lowPoints {
		var basin []int
		if !exist(basin, p) {
			basin = append(basin, p)
		}
		finded = append(finded, p)
		points := m.checkPoints(p, finded)
		for len(points) > 0 {
			var temp []int
			for _, d := range points {
				finded = append(finded, d)
				if m.Numbers[d] != 9 {
					if !exist(basin, d) {
						basin = append(basin, d)
					}
					temp = append(temp, m.checkPoints(d, finded)...)
				}
			}
			points = temp
		}
		basins[p] = basin //uniq(basin)
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
