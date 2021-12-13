package main

import (
	_ "embed"
	"fmt"
	"os"
	"sort"
	"strings"

	util "github.com/yhsiang/adventofcode"
)

//go:embed example
var example string

//go:embed input
var input string

func command(path string) (string, int) {
	n, _ := util.Int64(string(path[1:]))
	return string(path[0]), int(n)
}

func paths(lines []string) (int, [][]int) {
	var maps = make(map[string]int)
	for num, line := range lines {
		var x, y int
		for _, path := range strings.Split(line, ",") {
			direction, iteration := command(path)
			switch direction {
			case "R":
				for i := 0; i < iteration; i++ {
					x += 1
					coord := fmt.Sprintf("%d,%d", x, y)
					v, ok := maps[coord]
					if !ok {
						maps[coord] = num
					} else if v != num {
						maps[coord] = -1
					}
				}
			case "U":
				for i := 0; i < iteration; i++ {
					y += 1
					coord := fmt.Sprintf("%d,%d", x, y)
					v, ok := maps[coord]
					if !ok {
						maps[coord] = num
					} else if v != num {
						maps[coord] = -1
					}
				}
			case "L":
				for i := 0; i < iteration; i++ {
					x -= 1
					coord := fmt.Sprintf("%d,%d", x, y)
					v, ok := maps[coord]
					if !ok {
						maps[coord] = num
					} else if v != num {
						maps[coord] = -1
					}
				}
			case "D":
				for i := 0; i < iteration; i++ {
					y -= 1
					coord := fmt.Sprintf("%d,%d", x, y)
					v, ok := maps[coord]
					if !ok {
						maps[coord] = num
					} else if v != num {
						maps[coord] = -1
					}
				}
			}
		}
	}

	var distances []int
	var points [][]int
	for k, v := range maps {
		if v == -1 {
			// fmt.Println(k)
			x, y := util.Coord(k)
			// fmt.Println(x, y)
			distances = append(distances, util.Abs(x)+util.Abs(y))
			points = append(points, []int{x, y})
		}
	}
	sort.Ints(distances)
	return distances[0], points
}

func match(x, y int, points [][]int) bool {
	for _, p := range points {
		if p[0] == x && p[1] == y {
			return true
		}
	}
	return false
}

func paths2(lines []string, points [][]int) int {
	var maps = make(map[string]int)
	var totalSteps = make(map[string]int)
	for num, line := range lines {
		var x, y int
		var steps int
		for _, path := range strings.Split(line, ",") {
			direction, iteration := command(path)
			switch direction {
			case "R":
				for i := 0; i < iteration; i++ {
					x += 1
					steps++
					coord := fmt.Sprintf("%d,%d", x, y)
					v, ok := maps[coord]
					if !ok {
						maps[coord] = num
					} else if v != num {
						maps[coord] = -1
					}
					if match(x, y, points) {
						totalSteps[fmt.Sprintf("%d,%d", x, y)] += steps
					}
				}
			case "U":
				for i := 0; i < iteration; i++ {
					y += 1
					steps++
					coord := fmt.Sprintf("%d,%d", x, y)
					v, ok := maps[coord]
					if !ok {
						maps[coord] = num
					} else if v != num {
						maps[coord] = -1
					}
					if match(x, y, points) {
						totalSteps[fmt.Sprintf("%d,%d", x, y)] += steps
					}
				}
			case "L":
				for i := 0; i < iteration; i++ {
					x -= 1
					steps++
					coord := fmt.Sprintf("%d,%d", x, y)
					v, ok := maps[coord]
					if !ok {
						maps[coord] = num
					} else if v != num {
						maps[coord] = -1
					}
					if match(x, y, points) {
						totalSteps[fmt.Sprintf("%d,%d", x, y)] += steps
					}
				}
			case "D":
				for i := 0; i < iteration; i++ {
					y -= 1
					steps++
					coord := fmt.Sprintf("%d,%d", x, y)
					v, ok := maps[coord]
					if !ok {
						maps[coord] = num
					} else if v != num {
						maps[coord] = -1
					}
					if match(x, y, points) {
						totalSteps[fmt.Sprintf("%d,%d", x, y)] += steps
					}
				}
			}
		}
	}

	// fmt.Println(totalSteps)
	var best int = 9999999999999
	for _, v := range totalSteps {
		if v < best {
			best = v
		}
	}
	return best
}

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}
	lines := util.ByLine(file)
	distance, points := paths(lines)
	fmt.Printf("part1: %d\n", distance)
	fmt.Printf("part2: %d\n", paths2(lines, points))

}
