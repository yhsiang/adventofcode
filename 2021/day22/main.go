package main

import (
	_ "embed"
	"fmt"
	"os"
	"regexp"
	"strings"

	util "github.com/yhsiang/adventofcode"
)

//go:embed example
var example string

//go:embed input
var input string

var re = regexp.MustCompile(`x=(-?\d+)\.\.(-?\d+),y=(-?\d+)\.\.(-?\d+),z=(-?\d+)\.\.(-?\d+)`)

func getRange(input string) ([]int, []int, []int) {
	matched := re.FindSubmatch([]byte(input))
	x0, _ := util.Int(string(matched[1]))
	x1, _ := util.Int(string(matched[2]))
	y0, _ := util.Int(string(matched[3]))
	y1, _ := util.Int(string(matched[4]))
	z0, _ := util.Int(string(matched[5]))
	z1, _ := util.Int(string(matched[6]))

	return []int{x0, x1 + 1}, []int{y0, y1 + 1}, []int{z0, z1 + 1}
}

func part1(file string) int {
	lines := util.ByLine(file)
	var cubes = make(map[string]struct{})
	for _, line := range lines {
		s := strings.Split(line, " ")
		action := s[0]
		xr, yr, zr := getRange(s[1])
		switch action {
		case "on":
			for x := xr[0]; x < xr[1]; x++ {
				if x < -50 || x > 50 {
					continue
				}
				for y := yr[0]; y < yr[1]; y++ {
					if y < -50 || y > 50 {
						continue
					}
					for z := zr[0]; z < zr[1]; z++ {
						if z < -50 || z > 50 {
							continue
						}
						s := fmt.Sprintf("%d,%d,%d", x, y, z)
						cubes[s] = struct{}{}
					}
				}
			}
		case "off":
			for x := xr[0]; x < xr[1]; x++ {
				if x < -50 || x > 50 {
					continue
				}
				for y := yr[0]; y < yr[1]; y++ {
					if y < -50 || y > 50 {
						continue
					}
					for z := zr[0]; z < zr[1]; z++ {
						if z < -50 || z > 50 {
							continue
						}
						s := fmt.Sprintf("%d,%d,%d", x, y, z)
						delete(cubes, s)
					}
				}
			}
		}
	}
	return len(cubes)
}

type Cube struct {
	xStart int
	xEnd   int
	yStart int
	yEnd   int
	zStart int
	zEnd   int
	on     bool
}

func newCube(x []int, y []int, z []int, on bool) *Cube {
	return &Cube{
		xStart: x[0],
		xEnd:   x[1],
		yStart: y[0],
		yEnd:   y[1],
		zStart: z[0],
		zEnd:   z[1],
		on:     on,
	}
}

func part2(file string) int {
	lines := util.ByLine(file)
	var cubes []*Cube

	for _, line := range lines {
		s := strings.Split(line, " ")
		action := s[0]
		xr, yr, zr := getRange(s[1])
		cube := newCube(xr, yr, zr, action == "on")
		var newCubes []*Cube
		for _, c := range cubes {
			xOverlap := cube.xEnd > c.xStart && cube.xStart < c.xEnd
			yOverlap := cube.yEnd > c.yStart && cube.yStart < c.yEnd
			zOverlap := cube.zEnd > c.zStart && cube.zStart < c.zEnd
			if xOverlap && yOverlap && zOverlap {
				if c.xStart < cube.xStart {
					newCube := &Cube{
						xStart: c.xStart,
						xEnd:   cube.xStart,
						yStart: c.yStart,
						yEnd:   c.yEnd,
						zStart: c.zStart,
						zEnd:   c.zEnd,
						on:     c.on,
					}
					c.xStart = cube.xStart
					newCubes = append(newCubes, newCube)
				}

				if c.xEnd > cube.xEnd {
					newCube := &Cube{
						xStart: cube.xEnd,
						xEnd:   c.xEnd,
						yStart: c.yStart,
						yEnd:   c.yEnd,
						zStart: c.zStart,
						zEnd:   c.zEnd,
						on:     c.on,
					}
					c.xEnd = cube.xEnd
					newCubes = append(newCubes, newCube)
				}

				if c.yStart < cube.yStart {
					newCube := &Cube{
						xStart: c.xStart,
						xEnd:   c.xEnd,
						yStart: c.yStart,
						yEnd:   cube.yStart,
						zStart: c.zStart,
						zEnd:   c.zEnd,
						on:     c.on,
					}
					c.yStart = cube.yStart
					newCubes = append(newCubes, newCube)
				}

				if c.yEnd > cube.yEnd {
					newCube := &Cube{
						xStart: c.xStart,
						xEnd:   c.xEnd,
						yStart: cube.yEnd,
						yEnd:   c.yEnd,
						zStart: c.zStart,
						zEnd:   c.zEnd,
						on:     c.on,
					}
					c.yEnd = cube.yEnd
					newCubes = append(newCubes, newCube)
				}

				if c.zStart < cube.zStart {
					newCube := &Cube{
						xStart: c.xStart,
						xEnd:   c.xEnd,
						yStart: c.yStart,
						yEnd:   c.yEnd,
						zStart: c.zStart,
						zEnd:   cube.zStart,
						on:     c.on,
					}
					c.zStart = cube.zStart
					newCubes = append(newCubes, newCube)
				}

				if c.zEnd > cube.zEnd {
					newCube := &Cube{
						xStart: c.xStart,
						xEnd:   c.xEnd,
						yStart: c.yStart,
						yEnd:   c.yEnd,
						zStart: cube.zEnd,
						zEnd:   c.zEnd,
						on:     c.on,
					}
					c.zEnd = cube.zEnd
					newCubes = append(newCubes, newCube)
				}

			} else {
				newCubes = append(newCubes, c)
			}
		}
		newCubes = append(newCubes, cube)
		cubes = newCubes
	}

	var sum int
	for _, cube := range cubes {
		if cube.on {
			sum += (cube.xEnd - cube.xStart) * (cube.yEnd - cube.yStart) * (cube.zEnd - cube.zStart)
		}
	}
	return sum
}

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}

	fmt.Printf("part1: %d\n", part1(file))
	fmt.Printf("part2: %d\n", part2(file))

}
