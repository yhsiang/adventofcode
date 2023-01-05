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

func draw(from, to string, maps map[string]string) {
	fx, fy := util.Coord(from)
	tx, ty := util.Coord(to)
	// fmt.Printf("from %s to %s\n", from, to)
	if fx == tx {
		if ty > fy {
			for y := fy; y <= ty; y++ {
				c := fmt.Sprintf("%d,%d", fx, y)
				maps[c] = "#"
			}
		} else {
			for y := ty; y <= fy; y++ {
				c := fmt.Sprintf("%d,%d", fx, y)
				maps[c] = "#"
			}
		}

	} else if fy == ty {
		if tx > fx {
			for x := fx; x <= tx; x++ {
				c := fmt.Sprintf("%d,%d", x, fy)
				maps[c] = "#"
			}
		} else {
			for x := tx; x <= fx; x++ {
				c := fmt.Sprintf("%d,%d", x, fy)
				maps[c] = "#"
			}
		}
	}
}

func fall(sand string, maps map[string]string) (string, bool) {
	x, y := util.Coord(sand)
	d := fmt.Sprintf("%d,%d", x, y+1)
	_, okD := maps[d]
	// try left
	l := fmt.Sprintf("%d,%d", x-1, y+1)
	_, okL := maps[l]
	// try right
	r := fmt.Sprintf("%d,%d", x+1, y+1)
	_, okR := maps[r]

	if okD && !okL {
		return l, false
	}

	if okD && okL && !okR {
		return r, false
	}

	if okD && okL && okR {
		maps[sand] = "o"
		return sand, true
	}

	return d, false
}

func fall2(sand string, maxY int, maps map[string]string) (string, bool) {
	x, y := util.Coord(sand)
	d := fmt.Sprintf("%d,%d", x, y+1)
	_, okD := maps[d]
	// try left
	l := fmt.Sprintf("%d,%d", x-1, y+1)
	_, okL := maps[l]
	// try right
	r := fmt.Sprintf("%d,%d", x+1, y+1)
	_, okR := maps[r]

	if okD && !okL {
		return l, false
	}

	if okD && okL && !okR {
		return r, false
	}

	if okD && okL && okR {
		maps[sand] = "o"
		return sand, true
	}

	if y+1 == maxY {
		maps[sand] = "o"
		return sand, true
	}

	return d, false
}

func NewMap(input string) (map[string]string, int) {
	lines := strings.Split(input, "\n")
	from := ""
	to := ""
	maps := make(map[string]string)
	maxY := 0
	for _, line := range lines {
		coords := strings.Split(line, "->")
		from = strings.TrimSpace(coords[0])
		for i := 1; i < len(coords); i++ {
			to = strings.TrimSpace(coords[i])
			_, y := util.Coord(to)
			if y > maxY {
				maxY = y
			}
			draw(from, to, maps)
			from = to
		}
	}

	return maps, maxY
}

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}

	maps, maxY := NewMap(file)

	unit := 1
	loc := "500,0"
	for {
		next, rest := fall(loc, maps)
		if rest {
			next = "500,0"
			unit++
		} else {
			_, y := util.Coord(next)
			if y > maxY {
				unit -= 1
				break
			}
		}
		loc = next
	}

	fmt.Printf("%d\n", unit)

	maps, maxY = NewMap(file)
	unit = 1
	loc = "500,0"
	for {
		next, rest := fall2(loc, maxY+2, maps)
		if rest {
			if loc == next && next == "500,0" {
				break
			} else {
				next = "500,0"
				unit++
			}
		}
		loc = next
	}
	fmt.Printf("%d\n", unit)
}
