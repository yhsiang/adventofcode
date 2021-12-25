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

type cucumber struct {
	loc map[string]string
}

func newCucumber(input string) (*cucumber, int, int) {
	lines := util.ByLine(input)
	height := len(lines)
	width := len(lines[0])
	var loc = make(map[string]string)
	for i, line := range lines {
		for j, c := range strings.Split(line, "") {
			s := fmt.Sprintf("%d,%d", i, j)
			if c == ">" || c == "v" {
				loc[s] = c
			}
		}
	}

	return &cucumber{
		loc: loc,
	}, width, height
}

func move(c *cucumber, width, height int) int {
	var i = 1
	var temp = &cucumber{
		loc: c.loc,
	}
	for {
		var newLoc = make(map[string]string)
		for k, v := range temp.loc {
			if v != ">" {
				newLoc[k] = v
				continue
			}
			x, y := util.Coord(k)
			y += 1
			if y == width {
				y = 0
			}
			s := fmt.Sprintf("%d,%d", x, y)
			if _, ok := temp.loc[s]; !ok {
				newLoc[s] = v
				// n++
			} else {
				newLoc[k] = v
			}
		}

		var newLoc2 = make(map[string]string)
		for k, v := range newLoc {
			if v != "v" {
				newLoc2[k] = v
				continue
			}
			x, y := util.Coord(k)
			x += 1
			if x == height {
				x = 0
			}
			s := fmt.Sprintf("%d,%d", x, y)
			if _, ok := newLoc[s]; !ok {
				newLoc2[s] = v
				// n++
			} else {
				newLoc2[k] = v
			}
		}

		next := &cucumber{
			loc: newLoc2,
		}

		if temp.equal(next) {
			break
		}
		temp = next
		i++
	}

	return i
}

func (c *cucumber) equal(b *cucumber) bool {
	for k, v := range c.loc {
		vv, ok := b.loc[k]
		if !ok {
			return false
		}
		if v != vv {
			return false
		}
	}

	return true
}

func (c cucumber) print(height, width int) {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			s := fmt.Sprintf("%d,%d", i, j)
			v, ok := c.loc[s]
			if ok {
				fmt.Printf("%s", v)
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}

	c, width, height := newCucumber(file)
	a := move(c, width, height)
	fmt.Printf("part1: %d\n", a)
}
