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

func ManhattanDist(a, b string) int {
	x1, y1 := util.Coord(a)
	x2, y2 := util.Coord(b)
	return util.Abs(x1-x2) + util.Abs(y1-y2)
}

func main1() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}
	maps := make(map[string]string)
	closets := make(map[string]string)
	inputs := strings.Split(file, "\n")
	for _, input := range inputs {
		var sx, sy, bx, by int
		fmt.Sscanf(input, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		sensorCoord := fmt.Sprintf("%d,%d", sx, sy)
		beacnCoord := fmt.Sprintf("%d,%d", bx, by)
		maps[sensorCoord] = "S"
		maps[beacnCoord] = "B"
		// closets[beacnCoord] = sensorCoord
		closets[sensorCoord] = beacnCoord
	}

	targetY := 10
	if len(os.Args) == 2 && os.Args[1] == "input" {
		targetY = 2000000
	}

	pos := 0
	for k, v := range closets {
		x, y := util.Coord(k)
		fmt.Printf("%s\n", k)
		d := ManhattanDist(k, v)

		for i := x - d; i <= x+d; i++ {
			for j := y - d; j <= y+d; j++ {
				c := fmt.Sprintf("%d,%d", i, j)
				cd := ManhattanDist(c, k)
				_, ok := maps[c]
				if cd <= d && !ok {
					maps[c] = "#"
				}
			}
		}
	}

	for k, v := range maps {
		_, y := util.Coord(k)
		if y == targetY && v == "#" {
			pos += 1
		}
	}

	fmt.Printf("%d\n", pos)

}
