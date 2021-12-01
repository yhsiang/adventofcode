package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Grid struct {
	Value [][][]string
}

func NewGrid(n int) *Grid {
	var value = make([][][]string, n)
	for i := 0; i < n; i++ {
		value[i] = make([][]string, n)
		for j := 0; j < n; j++ {
			value[i][j] = make([]string, n)
			for k := 0; k < n; k++ {
				value[i][j][k] = "."
			}
		}
	}
	return &Grid{
		Value: value,
	}
}

func (g *Grid) update(lines []string) {
	for i, line := range lines {
		for j, cube := range strings.Split(line, "") {
			g.Value[12][i+9][j+9] = cube
		}
	}
}

func (g *Grid) run(cycles int) {
	var n = len(g.Value)
	var products = product([]int{0, 1, -1}, 3)

	for i := 0; i < cycles; i++ {
		clone := g.Clone()
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				for k := 0; k < n; k++ {
					var active = 0
					for _, p := range products {
						if p[0] == 0 && p[1] == 0 && p[2] == 0 {
							continue
						}
						var x = i + p[0]
						var y = j + p[1]
						var z = k + p[2]
						if (x >= 0 && x <= 23) &&
							(y >= 0 && y <= 23) &&
							(z >= 0 && z <= 23) &&
							g.Value[x][y][z] == "#" {
							active += 1
						}
					}

					if g.Value[i][j][k] == "#" && !(active >= 2 && active <= 3) {
						clone[i][j][k] = "."
					}
					if g.Value[i][j][k] == "." && (active == 3) {
						clone[i][j][k] = "#"
					}

				}
			}
		}

		g.Value = clone
	}
}

func (g *Grid) Clone() [][][]string {
	var n = len(g.Value)
	var a = make([][][]string, n)
	for i := 0; i < n; i++ {
		a[i] = make([][]string, n)
		for j := 0; j < n; j++ {
			a[i][j] = make([]string, n)
			for k := 0; k < n; k++ {
				a[i][j][k] = g.Value[i][j][k]
			}
		}
	}
	return a
}

func (g *Grid) count() int {
	var n = len(g.Value)
	var sum int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if g.Value[i][j][k] == "#" {
					sum += 1
				}
			}
		}
	}
	return sum
}

func product(a []int, repeat int) (r [][]int) {
	np := nextProduct(a, repeat)
	for {
		p := np()
		if len(p) == 0 {
			break
		}
		var c []int
		c = append(c, p...)
		r = append(r, c)

	}
	return
}

func nextProduct(a []int, r int) func() []int {
	p := make([]int, r)
	x := make([]int, len(p))
	return func() []int {
		p := p[:len(x)]
		for i, xi := range x {
			p[i] = a[xi]
		}
		for i := len(x) - 1; i >= 0; i-- {
			x[i]++
			if x[i] < len(a) {
				break
			}
			x[i] = 0
			if i <= 0 {
				x = x[0:0]
				break
			}
		}
		return p
	}
}

type HyperGrid [][][][]string

func NewHyperGrid(n int) HyperGrid {
	var value = make([][][][]string, n)
	for i := 0; i < n; i++ {
		value[i] = make([][][]string, n)
		for j := 0; j < n; j++ {
			value[i][j] = make([][]string, n)
			for k := 0; k < n; k++ {
				value[i][j][k] = make([]string, n)
				for w := 0; w < n; w++ {
					value[i][j][k][w] = "."
				}
			}
		}
	}
	return value
}

func (h HyperGrid) update(lines []string) {
	for i, line := range lines {
		for j, cube := range strings.Split(line, "") {
			h[12][12][i+9][j+9] = cube
		}
	}
}

func (h *HyperGrid) run(cycles int) {
	var n = len(*h)
	var products = product([]int{0, 1, -1}, 4)

	for i := 0; i < cycles; i++ {
		clone := h.clone()
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				for k := 0; k < n; k++ {
					for f := 0; f < n; f++ {
						var active = 0
						for _, p := range products {
							if p[0] == 0 && p[1] == 0 && p[2] == 0 && p[3] == 0 {
								continue
							}
							var x = i + p[0]
							var y = j + p[1]
							var z = k + p[2]
							var w = f + p[3]
							if (x >= 0 && x <= 23) &&
								(y >= 0 && y <= 23) &&
								(z >= 0 && z <= 23) &&
								(w >= 0 && w <= 23) &&
								(*h)[x][y][z][w] == "#" {
								active += 1
							}
						}

						if (*h)[i][j][k][f] == "#" && !(active >= 2 && active <= 3) {
							clone[i][j][k][f] = "."
						}
						if (*h)[i][j][k][f] == "." && (active == 3) {
							clone[i][j][k][f] = "#"
						}
					}
				}
			}
		}

		*h = clone
	}
}

func (h HyperGrid) clone() [][][][]string {
	var n = len(h)
	var a = make([][][][]string, n)
	for i := 0; i < n; i++ {
		a[i] = make([][][]string, n)
		for j := 0; j < n; j++ {
			a[i][j] = make([][]string, n)
			for k := 0; k < n; k++ {
				a[i][j][k] = make([]string, n)
				for w := 0; w < n; w++ {
					a[i][j][k][w] = h[i][j][k][w]
				}
			}
		}
	}
	return a
}

func (h HyperGrid) count() int {
	var n = len(h)
	var sum int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				for w := 0; w < n; w++ {
					if h[i][j][k][w] == "#" {
						sum += 1
					}
				}
			}
		}
	}
	return sum
}

func main() {
	dat, err := ioutil.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")
	// Part1
	// g := NewGrid(24)
	// g.update(lines)
	// g.run(6)
	// fmt.Println(g.count())

	g := NewHyperGrid(24)
	g.update(lines)
	g.run(6)
	fmt.Println(g.count())

}
