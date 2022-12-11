package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
)

//go:embed example
var example string

//go:embed input
var input string

type Register struct {
	X        int
	Commands []string
	Cycle    int
	Sprites  string
}

func NewRegister(commands []string) *Register {
	return &Register{
		X:        1,
		Commands: commands,
		Cycle:    1,
		Sprites:  "###.....................................",
	}
}

func (r *Register) moveSprites() {
	sprites := ""
	for i := 0; i < 40; i++ {
		if i >= r.X-1 && i <= r.X+1 {
			sprites += "#"
		} else {
			sprites += "."
		}
	}
	r.Sprites = sprites
}

func (r *Register) run(cycle int) (int, string) {
	var beginFlag = false
	CRT := ""
	pos := 0
	for r.Cycle <= cycle {
		command := r.Commands[0]
		// fmt.Printf("current %d, register %d command %s\n", r.Cycle, r.X, command)
		CRT += string(r.Sprites[pos])
		var ins string
		var value int
		fmt.Sscanf(command, "%s %d", &ins, &value)
		// fmt.Printf("%d\n", r.lastCyle)
		switch ins {
		case "noop": // 1 cycle
			r.Commands = r.Commands[1:]
			r.Commands = append(r.Commands, command)
		case "addx": // 2 cycles
			if beginFlag {
				r.X += value
				r.moveSprites()
				r.Commands = r.Commands[1:]
				r.Commands = append(r.Commands, command)
				beginFlag = false
			} else {
				beginFlag = true

			}
		}

		pos += 1
		r.Cycle += 1
	}

	return r.X, CRT
}

func (r *Register) run2(cycle int) []string {
	var beginFlag = false
	var crts []string
	crt := ""
	pos := 0
	for r.Cycle <= cycle {
		command := r.Commands[0]
		// fmt.Printf("current %d, register %d command %s\n", r.Cycle, r.X, command)
		crt += string(r.Sprites[pos])
		var ins string
		var value int
		fmt.Sscanf(command, "%s %d", &ins, &value)
		// fmt.Printf("%d\n", r.lastCyle)
		switch ins {
		case "noop": // 1 cycle
			r.Commands = r.Commands[1:]
			r.Commands = append(r.Commands, command)
		case "addx": // 2 cycles
			if beginFlag {
				r.X += value
				r.moveSprites()
				r.Commands = r.Commands[1:]
				r.Commands = append(r.Commands, command)
				beginFlag = false
			} else {
				beginFlag = true

			}
		}

		if r.Cycle%40 == 0 {
			pos = 0
			crts = append(crts, crt)
			crt = ""
		} else {
			pos += 1
		}
		r.Cycle += 1
	}

	return crts
}

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}

	lines := strings.Split(file, "\n")

	x := NewRegister(lines)

	runs := []int{20, 60, 100, 140, 180, 220}
	sum := 0

	for _, run := range runs {
		register, _ := x.run(run)
		sum += run * register
	}
	// first puzzle
	fmt.Printf("%d\n", sum)

	y := NewRegister(lines)
	crts := y.run2(240)

	for _, crt := range crts {
		fmt.Printf("%s\n", crt)
	}

}
