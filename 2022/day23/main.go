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

type Map struct {
	Values map[string]string
	Rows   int
	Cols   int
}

func NewMap(input string) *Map {
	rows := strings.Split(input, "\n")
	rowLen := len(rows)
	colLen := len(strings.Split(rows[0], ""))
	values := make(map[string]string)

	for x, row := range rows {
		for y, col := range strings.Split(row, "") {
			coord := fmt.Sprintf("%d,%d", x, y)
			values[coord] = col
		}
	}
	return &Map{
		Values: values,
		Rows:   rowLen,
		Cols:   colLen,
	}
}

func (m *Map) ShouldMove(i, j int) bool {

	nCoord := fmt.Sprintf("%d,%d", i-1, j)
	neCoord := fmt.Sprintf("%d,%d", i-1, j+1)
	nwCoord := fmt.Sprintf("%d,%d", i-1, j+1)
	sCoord := fmt.Sprintf("%d,%d", i+1, j)
	seCoord := fmt.Sprintf("%d,%d", i+1, j-1)
	swCoord := fmt.Sprintf("%d,%d", i+1, j+1)
	wCoord := fmt.Sprintf("%d,%d", i, j-1)
	eCoord := fmt.Sprintf("%d,%d", i, j+1)
	nv := m.Values[nCoord]
	nev := m.Values[neCoord]
	nwv := m.Values[nwCoord]
	sv := m.Values[sCoord]
	sev := m.Values[seCoord]
	swv := m.Values[swCoord]
	wv := m.Values[wCoord]
	ev := m.Values[eCoord]
	if nv == "." && nev == "." && nwv == "." && sv == "." && sev == "." && swv == "." && wv == "." && ev == "." {
		return false
	}
	return true
}

func (m *Map) Move(i, j int) string {
	coord := fmt.Sprintf("%d,%d", i, j)

	nCoord := fmt.Sprintf("%d,%d", i-1, j)
	neCoord := fmt.Sprintf("%d,%d", i-1, j+1)
	nwCoord := fmt.Sprintf("%d,%d", i-1, j+1)
	nv := m.Values[nCoord]
	nev := m.Values[neCoord]
	nwv := m.Values[nwCoord]
	if nv == "." && nev == "." && nwv == "." {
		return nCoord
	}

	sCoord := fmt.Sprintf("%d,%d", i+1, j)
	seCoord := fmt.Sprintf("%d,%d", i+1, j-1)
	swCoord := fmt.Sprintf("%d,%d", i+1, j+1)
	sv := m.Values[sCoord]
	sev := m.Values[seCoord]
	swv := m.Values[swCoord]
	if sv == "." && sev == "." && swv == "." {
		return sCoord
	}

	wCoord := fmt.Sprintf("%d,%d", i, j-1)
	// nwCoord := fmt.Sprintf("%d,%d", x-1, y-1)
	// swCoord := fmt.Sprintf("%d,%d", x-1, y+1)
	wv := m.Values[wCoord]
	// nwv := m.Values[nwCoord]
	// swv := m.Values[swCoord]
	if wv == "." && nwv == "." && swv == "." {

		return wCoord
	}

	eCoord := fmt.Sprintf("%d,%d", i, j+1)
	// nwCoord := fmt.Sprintf("%d,%d", x-1, y-1)
	// swCoord := fmt.Sprintf("%d,%d", x-1, y+1)
	ev := m.Values[eCoord]
	// nwv := m.Values[nwCoord]
	// swv := m.Values[swCoord]
	if ev == "." && nev == "." && sev == "." {
		return eCoord
	}

	return coord
}

func Scan(m *Map) *Map {
	temp := &Map{
		Values: make(map[string]string),
		Rows:   m.Rows,
		Cols:   m.Cols,
	}
	target := make(map[string]string)

	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			coord := fmt.Sprintf("%d,%d", i, j)
			temp.Values[coord] = "."
			if m.Values[coord] == "." {
				continue
			}

			if m.ShouldMove(i, j) {
				moveCoord := m.Move(i, j)
				if v, ok := target[moveCoord]; ok {
					target[v] = v
					target[coord] = coord
					delete(target, moveCoord)
				} else {
					target[moveCoord] = coord
				}
			}
		}
	}

	for k := range target {
		temp.Values[k] = "#"
	}
	return temp
}

func (m *Map) Print() {

	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			coord := fmt.Sprintf("%d,%d", i, j)
			fmt.Printf("%s", m.Values[coord])
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

	mp := NewMap(file)
	mp.Print()
	mp = Scan(mp)
	//
	mp.Print()
	mp = Scan(mp)
	//
	mp.Print()
	// strings.Split(file, "\n")
}
