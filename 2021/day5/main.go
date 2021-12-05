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
		// fmt.Printf("%d", len(rows))
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

func (d *Diagram) move(isPart2 bool) {
	for _, line := range d.Lines {
		if line.From.X == line.To.X {
			if line.To.Y > line.From.Y {
				for i := int(line.From.Y); i <= int(line.To.Y); i++ {
					d.Map[i][int(line.From.X)] += 1
				}
			} else {
				for i := int(line.To.Y); i <= int(line.From.Y); i++ {
					d.Map[i][int(line.To.X)] += 1
				}
			}
		}
		if line.From.Y == line.To.Y {
			if line.To.X > line.From.X {
				for i := int(line.From.X); i <= int(line.To.X); i++ {
					d.Map[int(line.From.Y)][i] += 1
				}
			} else {
				for i := int(line.To.X); i <= int(line.From.X); i++ {
					d.Map[int(line.To.Y)][i] += 1
				}
			}
		}
		if isPart2 {
			if line.From.X != line.To.X && line.From.Y != line.To.Y {
				if line.From.X > line.To.X && line.From.Y < line.To.Y {
					// fmt.Printf("%+v\n", line)
					i := int(line.From.X)
					j := int(line.From.Y)
					for i >= int(line.To.X) && j <= int(line.To.Y) {
						d.Map[j][i] += 1
						i--
						j++
					}
				}

				if line.From.X < line.To.X && line.From.Y < line.To.Y {
					// fmt.Printf("%+v\n", line)
					i := int(line.From.X)
					j := int(line.From.Y)
					for i <= int(line.To.X) && j <= int(line.To.Y) {
						d.Map[j][i] += 1
						i++
						j++
					}
				}

				// 6,4 -> 2,0
				// 6,4 5,3 4,2 3,1 2,0
				if line.From.X > line.To.X && line.From.Y > line.To.Y {
					i := int(line.From.X)
					j := int(line.From.Y)
					for i >= int(line.To.X) && j >= int(line.To.Y) {
						d.Map[j][i] += 1
						i--
						j--
					}
				}
				// 5,5 -> 8,2
				// 5,5 6,4 7,3, 8,2
				if line.From.X < line.To.X && line.From.Y > line.To.Y {
					// fmt.Printf("%+v\n", line)
					i := int(line.From.X)
					j := int(line.From.Y)
					for i <= int(line.To.X) && j >= int(line.To.Y) {
						d.Map[j][i] += 1
						i++
						j--
					}
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
	// diagram2.print()
	fmt.Printf("part1: %d\n", diagram2.overlap())

}
