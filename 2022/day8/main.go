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

type Tree struct {
	Visible bool
	Height  int
}

type TreeMap struct {
	Trees  map[string]*Tree
	RowLen int
	ColLen int
}

func NewMap(input string) *TreeMap {
	var trees = make(map[string]*Tree)
	rows := strings.Split(input, "\n")
	rowLen := len(rows)
	colLen := len(strings.Split(rows[0], ""))
	for i, s := range rows {

		for j, h := range strings.Split(s, "") {
			height, _ := util.Int(h)
			coord := fmt.Sprintf("%d,%d", i, j)
			visible := false
			if i == 0 || j == 0 || i == rowLen-1 || j == rowLen-1 {
				visible = true
			}

			trees[coord] = &Tree{
				Visible: visible,
				Height:  height,
			}
		}
	}

	return &TreeMap{
		Trees:  trees,
		RowLen: rowLen,
		ColLen: colLen,
	}
}

func (m *TreeMap) print() {
	var i = 0
	for _, tree := range m.Trees {
		fmt.Printf("%d", tree.Height)
		if i%m.RowLen == 4 {
			fmt.Println()
		}
		i++
	}
}

func (m *TreeMap) scan() {

	for i := 1; i < m.RowLen-1; i++ {
		for j := 1; j < m.ColLen-1; j++ {
			coord := fmt.Sprintf("%d,%d", i, j)
			tree := m.Trees[coord]
			// fmt.Printf("%s %d\n", coord, tree.Height)
			if tree.Visible {
				continue
			}

			up := i - 1
			upCount := 0
			for up >= 0 {
				upCoord := fmt.Sprintf("%d,%d", up, j)
				upTree := m.Trees[upCoord]
				// fmt.Printf("  %s\n",  )
				if upTree.Height >= tree.Height {
					break
				}
				upCount += 1
				up -= 1
			}

			left := j - 1
			leftCount := 0
			for left >= 0 {
				leftCoord := fmt.Sprintf("%d,%d", i, left)
				leftTree := m.Trees[leftCoord]
				// fmt.Printf("  %s\n", upTree)
				if leftTree.Height >= tree.Height {
					break
				}
				leftCount += 1
				left -= 1
			}

			down := i + 1
			downCount := 0
			for down < m.RowLen {
				downCoord := fmt.Sprintf("%d,%d", down, j)
				downTree := m.Trees[downCoord]
				// fmt.Printf("  %s %d\n", downCoord, downTree.Height)
				if downTree.Height >= tree.Height {
					break
				}
				downCount += 1
				down += 1
			}

			right := j + 1
			rightCount := 0
			for right < m.ColLen {
				rightCoord := fmt.Sprintf("%d,%d", i, right)
				rightTree := m.Trees[rightCoord]
				if rightTree.Height >= tree.Height {
					break
				}
				rightCount += 1
				right += 1
			}

			// fmt.Printf("%d, %d, %d, %d\n", leftCount, rightCount, upCount, downCount)
			// fmt.Printf("%d, %d, %d, %d\n", i, m.ColLen-1-i, j, m.RowLen-1-j)
			if leftCount == j || rightCount == m.RowLen-1-j || upCount == i || downCount == m.ColLen-1-i {
				tree.Visible = true
			}

		}
	}
}

func (m *TreeMap) scan2() int {
	var scores []int

	for i := 1; i < m.RowLen-1; i++ {
		for j := 1; j < m.ColLen-1; j++ {
			coord := fmt.Sprintf("%d,%d", i, j)
			tree := m.Trees[coord]
			fmt.Printf("%s %d\n", coord, tree.Height)
			if tree.Visible {
				continue
			}

			up := i - 1
			upCount := 0
			for up >= 0 {
				upCount += 1
				upCoord := fmt.Sprintf("%d,%d", up, j)
				upTree := m.Trees[upCoord]
				// fmt.Printf("  %s\n",  )
				if upTree.Height >= tree.Height {
					break
				}

				up -= 1
			}

			left := j - 1
			leftCount := 0
			for left >= 0 {
				leftCount += 1
				leftCoord := fmt.Sprintf("%d,%d", i, left)
				leftTree := m.Trees[leftCoord]
				// fmt.Printf("  %s\n", upTree)
				if leftTree.Height >= tree.Height {
					break
				}

				left -= 1
			}

			down := i + 1
			downCount := 0
			for down < m.RowLen {
				downCount += 1
				downCoord := fmt.Sprintf("%d,%d", down, j)
				downTree := m.Trees[downCoord]
				// fmt.Printf("  %s %d\n", downCoord, downTree.Height)
				if downTree.Height >= tree.Height {
					break
				}

				down += 1
			}

			right := j + 1
			rightCount := 0
			for right < m.ColLen {
				rightCount += 1
				rightCoord := fmt.Sprintf("%d,%d", i, right)
				rightTree := m.Trees[rightCoord]
				if rightTree.Height >= tree.Height {
					break
				}

				right += 1
			}
			fmt.Printf("%d, %d, %d, %d\n", leftCount, rightCount, upCount, downCount)
			fmt.Printf("%d, %d, %d, %d\n", j, m.RowLen-1-j, i, m.ColLen-1-i)
			score := leftCount * rightCount * upCount * downCount
			scores = append(scores, score)
		}
	}

	max := 0
	for _, score := range scores {
		if score > max {
			max = score
		}
	}
	return max
}

func (m *TreeMap) cal() int {
	count := 0
	for _, tree := range m.Trees {
		if tree.Visible {
			count += 1
		}
	}
	return count
}

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}

	// input := strings.Split(file, "\n")

	m := NewMap(file)
	// m.print()
	m.scan()
	// first puzzle
	fmt.Printf("%d\n", m.cal())
	m2 := NewMap(file)
	fmt.Printf("%d\n", m2.scan2())

}
