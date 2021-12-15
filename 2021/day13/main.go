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

type Paper struct {
	Dots map[string]string
}

func initPaper(data string) *Paper {
	var dots = make(map[string]string)
	for _, line := range util.ByLine(data) {
		dots[line] = "#"
	}

	return &Paper{
		Dots: dots,
	}
}

func (p *Paper) fold(action string) int {
	act := strings.Replace(action, "fold along ", "", 1)
	t := strings.Split(act, "=")
	c64, _ := util.Int64(t[1])
	c := int(c64)
	switch t[0] {
	case "y": // fold up
		for dot := range p.Dots {
			x, y := util.Coord(dot)
			if y < c {
				continue
			}
			y = c - (y - c)
			delete(p.Dots, dot)
			p.Dots[fmt.Sprintf("%d,%d", x, y)] = "#"
		}
	case "x": // fold left
		for dot := range p.Dots {
			x, y := util.Coord(dot)
			if x < c {
				continue
			}
			x = c - (x - c)
			delete(p.Dots, dot)
			p.Dots[fmt.Sprintf("%d,%d", x, y)] = "#"
		}
	}

	return len(p.Dots)
}

func (p *Paper) print(x, y int) {
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			d := fmt.Sprintf("%d,%d", j, i)
			v, ok := p.Dots[d]
			if ok {
				fmt.Printf("%s", v)
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}

	inputs := strings.Split(file, "\n\n")
	paper := initPaper(inputs[0])
	actions := util.ByLine(inputs[1])
	fmt.Printf("part1: %d\n", paper.fold(actions[0]))

	paper = initPaper(inputs[0])
	for _, act := range actions {
		paper.fold(act)
	}
	fmt.Println("part2:")
	paper.print(6, 39)
}
