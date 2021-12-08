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

type Coordinate struct {
	X int64
	Y int64
}

func initCoordinates(input string) Coordinate {
	// 0,9
	coordinates := strings.Split(input, ",")

	x, _ := util.Int64(coordinates[0])
	y, _ := util.Int64(coordinates[1])

	return Coordinate{
		X: x,
		Y: y,
	}
}

type Line struct {
	From Coordinate
	To   Coordinate
}

func initLine(input string) *Line {
	data := strings.Split(input, "->")
	from := initCoordinates(strings.TrimSpace(data[0]))
	to := initCoordinates(strings.TrimSpace(data[1]))

	return &Line{
		From: from,
		To:   to,
	}
}

type Diagram struct {
	Lines   []*Line
	Map     [][]int
	Hashmap map[string]int
}

func initDiagram(input []string) *Diagram {
	var lines []*Line
	for _, d := range input {
		lines = append(lines, initLine(d))
	}

	return &Diagram{
		Lines:   lines,
		Hashmap: make(map[string]int),
	}
}

// debug
func (d *Diagram) print() {
	for _, rows := range d.Map {
		for _, col := range rows {
			fmt.Printf("%d", col)
		}
		fmt.Printf("\n")
	}
}

func signum(from int64, to int64) int {
	if from == to {
		return 0
	} else if to-from > 0 {
		return 1
	} else {
		return -1
	}
}

func (d *Diagram) draw(isPart2 bool) {
	for _, line := range d.Lines {
		if !isPart2 {
			if line.From.X == line.To.X || line.From.Y == line.To.Y {
				d.realDraw(line)
			}
		} else {
			d.realDraw(line)
		}
	}
}

func (d *Diagram) realDraw(line *Line) {
	dx := signum(line.From.X, line.To.X)
	dy := signum(line.From.Y, line.To.Y)
	for i, j := line.From.X, line.From.Y; i != line.To.X+int64(dx) || j != line.To.Y+int64(dy); i, j = i+int64(dx), j+int64(dy) {
		coord := fmt.Sprintf("%d,%d", i, j)
		d.Hashmap[coord] += 1
	}
}

func (d *Diagram) overlap() int {
	var overlap = 0
	for _, d := range d.Hashmap {
		if d >= 2 {
			overlap++
		}
	}
	return overlap
}

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}

	data := util.ByLine(file)

	diagram1 := initDiagram(data)
	diagram1.draw(false)
	fmt.Printf("part1: %d\n", diagram1.overlap())

	diagram2 := initDiagram(data)
	diagram2.draw(true)
	fmt.Printf("part2: %d\n", diagram2.overlap())

}
