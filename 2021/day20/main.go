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

type Image struct {
	dots   map[string]struct{}
	width  int
	height int
}

var neighbors = [][]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},

	{0, -1},
	{0, 0},
	{0, 1},

	{1, -1},
	{1, 0},
	{1, 1},
}

func (im *Image) minMax() (int, int, int, int) {
	var xs []int
	var ys []int
	for k := range im.dots {
		x, y := util.Coord(k)
		xs = append(xs, x)
		ys = append(ys, y)
	}
	minX, maxX := util.MinMax(xs)
	minY, maxY := util.MinMax(ys)
	return minX, minY, maxX, maxY
}

func (im *Image) scan(x, y int, boundary map[string]struct{}, outside bool) int {
	var output string
	for _, n := range neighbors {
		s := fmt.Sprintf("%d,%d", x+n[0], y+n[1])
		_, ok1 := im.dots[s]
		if outside {
			_, ok2 := boundary[s]
			if ok1 || !ok2 {
				output += "1"
			} else {
				output += "0"
			}
		} else {
			if ok1 {
				output += "1"
			} else {
				output += "0"
			}
		}

	}
	n, _ := util.BinToInt(output)
	return n
}

func (im *Image) enhance(algo string, steps int) {
	for i := 0; i < steps; i++ {
		var dots = make(map[string]struct{})
		minX, minY, maxX, maxY := im.minMax()
		boundary := make(map[string]struct{})
		for x := minX; x < maxX+1; x++ {
			for y := minY; y < maxY+1; y++ {
				s := fmt.Sprintf("%d,%d", x, y)
				boundary[s] = struct{}{}
			}
		}
		outside := i%2 == 1
		for k := minX - 1; k <= maxX+1; k++ {
			for j := minY - 1; j <= maxY+1; j++ {
				v := fmt.Sprintf("%d,%d", k, j)
				n := im.scan(k, j, boundary, outside)

				if algo[n] == '#' {
					dots[v] = struct{}{}
				}
			}
		}
		im.dots = dots
	}

}

func initMap(input string, algo string) *Image {
	var dots = make(map[string]struct{})
	lines := util.ByLine(input)
	height := len(lines)
	width := len(lines[0])
	for i, line := range lines {
		for j, d := range line {
			if d == '#' {
				s := fmt.Sprintf("%d,%d", i, j)
				dots[s] = struct{}{}
			}
		}
	}

	return &Image{
		dots:   dots,
		width:  width,
		height: height,
	}
}

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}

	data := strings.Split(file, "\n\n")
	algo := data[0]

	m := initMap(data[1], data[0])
	m.enhance(algo, 2)
	fmt.Printf("part1: %d\n", len(m.dots))
	m = initMap(data[1], data[0])
	m.enhance(algo, 50)
	fmt.Printf("part2: %d\n", len(m.dots))

}
