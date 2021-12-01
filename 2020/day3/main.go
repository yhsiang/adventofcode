package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Map struct {
	Right  int
	Down   int
	maps   [][]string
	rowLen int
}

func NewMap(data []byte, r, d int) *Map {
	lines := strings.Split(string(data), "\n")

	var theMap [][]string
	for _, line := range lines {
		chars := strings.Split(line, "")
		theMap = append(theMap, chars)
	}

	return &Map{
		Right:  r,
		Down:   d,
		maps:   theMap,
		rowLen: len(theMap[0]),
	}
}

func (m *Map) CountTrees() int {
	var r = 0
	var treeNum = 0
	for i := 0; i < len(m.maps); i += m.Down {
		if i == 0 {
			continue
		}
		r += m.Right
		if r >= m.rowLen {
			r -= m.rowLen
		}
		if m.maps[i][r] == "#" {
			treeNum++
		}
	}

	return treeNum
}

func main() {
	dat, err := ioutil.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	m1 := NewMap(dat, 1, 1)
	v1 := m1.CountTrees()
	fmt.Println(v1)
	m2 := NewMap(dat, 3, 1)
	v2 := m2.CountTrees()
	fmt.Println(v2)
	m3 := NewMap(dat, 5, 1)
	v3 := m3.CountTrees()
	fmt.Println(v3)
	m4 := NewMap(dat, 7, 1)
	v4 := m4.CountTrees()
	fmt.Println(v4)
	m5 := NewMap(dat, 1, 2)
	v5 := m5.CountTrees()
	fmt.Println(v5)
	fmt.Println(v1 * v2 * v3 * v4 * v5)
}
