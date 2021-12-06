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
	// Direction string
	// Moves     int64
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
	Lines []*Line
	Map   [][]int
}

func initDiagram(input []string, max int) *Diagram {
	var lines []*Line
	for _, d := range input {
		lines = append(lines, initLine(d))
	}
	var output [][]int
	for i := 0; i < max; i++ {
		var temp = []int{}
		for j := 0; j < max; j++ {
			temp = append(temp, 0)
		}
		output = append(output, temp)
	}
	return &Diagram{
		Lines: lines,
		Map:   output,
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
	if to-from > 0 {
		return 1
	} else {
		return -1
	}
}

func (d *Diagram) move(isPart2 bool) {
	for _, line := range d.Lines {
		if line.From.X == line.To.X {
			dy := signum(line.From.Y, line.To.Y)
			for i, j := line.From.Y, line.To.Y; i != j+int64(dy); i += int64(dy) {
				d.Map[i][line.From.X] += 1
			}
		}
		if line.From.Y == line.To.Y {
			dx := signum(line.From.X, line.To.X)
			for i, j := line.From.X, line.To.X; i != j+int64(dx); i += int64(dx) {
				d.Map[line.From.Y][i] += 1
			}
		}
		if isPart2 {
			if line.From.X != line.To.X && line.From.Y != line.To.Y {
				dx := signum(line.From.X, line.To.X)
				dy := signum(line.From.Y, line.To.Y)
				for i, j := line.From.X, line.From.Y; i != line.To.X+int64(dx) && j != line.To.Y+int64(dy); i, j = i+int64(dx), j+int64(dy) {
					d.Map[j][i] += 1
				}
			}
		}
	}
}

func (d *Diagram) overlap() int {
	var overlap = 0
	for _, rows := range d.Map {
		for _, col := range rows {
			if col >= 2 {
				overlap++
			}
		}
	}
	return overlap
}

func main() {
	var file = example
	var max = 10
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
		max = 1000
	}

	data := util.Read(file)

	diagram1 := initDiagram(data, max)
	diagram1.move(false)
	fmt.Printf("part1: %d\n", diagram1.overlap())

	diagram2 := initDiagram(data, max)
	diagram2.move(true)
	fmt.Printf("part2: %d\n", diagram2.overlap())

}
