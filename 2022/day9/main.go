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

func move(direction string, hx, hy int) (int, int) {
	switch direction {
	case "R":
		hx += 1
	case "L":
		hx -= 1
	case "U":
		hy += 1
	case "D":
		hy -= 1
	}
	return hx, hy
}

func follow(hx, hy int, tx, ty int) (int, int) {
	// tx - hx = 1, hy - ty = 2 H H   H H  hx-tx = 1, hy-ty = 2
	// tx - hx = 2, hy - ty = 1 H x   x H  hx-tx = 2, hy-ty = 1
	//                              T
	// tx - hx = 2, ty - hy = 1 H x   x H  hx-tx = 2, ty-hy = 1
	// tx - hx = 1, ty - hy = 2 H H   H H  hx-tx = 1, ty-hy = 2

	if (hx-tx == 2 && hy-ty == 1) || (hx-tx == 1 && hy-ty == 2) || (hx-tx == 2 && hy-ty == 2) {
		return tx + 1, ty + 1
	}

	if (hx-tx == 2 && ty-hy == 1) || (hx-tx == 1 && ty-hy == 2) || (hx-tx == 2 && ty-hy == 2) {
		return tx + 1, ty - 1
	}

	if (tx-hx == 1 && hy-ty == 2) || (tx-hx == 2 && hy-ty == 1) || (tx-hx == 2 && hy-ty == 2) {
		return tx - 1, ty + 1
	}

	if (tx-hx == 2 && ty-hy == 1) || (tx-hx == 1 && ty-hy == 2) || (tx-hx == 2 && ty-hy == 2) {
		return tx - 1, ty - 1
	}

	//     H
	// 	   x
	// H x T x H
	// 	   x
	// 	   H
	if hx-tx == 2 {
		return tx + 1, ty
	}

	if tx-hx == 2 {
		return tx - 1, ty
	}

	if hy-ty == 2 {
		return tx, ty + 1
	}

	if ty-hy == 2 {
		return tx, ty - 1
	}

	return tx, ty
}
func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}

	steps := strings.Split(file, "\n")

	states := map[string][]int{
		"H": {0, 0},
		"T": {0, 0},
	}

	visited := map[string]struct{}{}

	for _, step := range steps {
		s, _ := util.Int(string(step[2:]))
		direction := string(step[0])
		for i := 1; i <= s; i++ {
			hx, hy := states["H"][0], states["H"][1]
			tx, ty := states["T"][0], states["T"][1]
			hx, hy = move(direction, hx, hy)

			tx, ty = follow(hx, hy, tx, ty)

			coord := fmt.Sprintf("%d,%d", tx, ty)
			if _, ok := visited[coord]; !ok {
				visited[coord] = struct{}{}
			}
			// fmt.Printf("%d, %d | %d, %d\n", hx, hy, tx, ty)
			states["H"] = []int{hx, hy}
			states["T"] = []int{tx, ty}
		}

	}

	// first puzzle
	fmt.Printf("%d\n", len(visited))

	states = map[string][]int{
		"H": {0, 0},
		"1": {0, 0},
		"2": {0, 0},
		"3": {0, 0},
		"4": {0, 0},
		"5": {0, 0},
		"6": {0, 0},
		"7": {0, 0},
		"8": {0, 0},
		"9": {0, 0},
	}
	visited = map[string]struct{}{}
	knots := []string{"H", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	for _, step := range steps {
		s, _ := util.Int(string(step[2:]))
		direction := string(step[0])
		for i := 1; i <= s; i++ {
			for j, knot := range knots[:len(knots)-1] {
				next := knots[j+1]
				hx, hy := states[knot][0], states[knot][1]
				tx, ty := states[next][0], states[next][1]

				if knot == "H" {
					hx, hy = move(direction, hx, hy)
				}

				tx, ty = follow(hx, hy, tx, ty)

				if next == "9" {
					coord := fmt.Sprintf("%d,%d", tx, ty)
					if _, ok := visited[coord]; !ok {
						visited[coord] = struct{}{}
					}
				}
				states[knot] = []int{hx, hy}
				states[next] = []int{tx, ty}
			}
		}
	}

	fmt.Printf("%d\n", len(visited))
}
