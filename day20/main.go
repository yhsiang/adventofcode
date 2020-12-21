package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var reTile = regexp.MustCompile(`^Tile (\d+):$`)

type tile struct {
	id   int64
	data [][]string

	topTile    int64
	bottomTile int64
	leftTile   int64
	rightTile  int64

	locked  bool
	visited bool
}

func NewTile(str string) *tile {
	tile := &tile{locked: false}
	lines := strings.Split(str, "\n")
	if id := reTile.FindSubmatch([]byte(lines[0])); len(id) == 2 {
		i, _ := strconv.ParseInt(string(id[1]), 10, 64)
		tile.id = i
	}
	for _, line := range lines[1:] {
		tile.data = append(tile.data, strings.Split(line, ""))
	}
	return tile
}

func (t *tile) rotate() {
	var data [][]string
	var last = len(t.data) - 1
	for i, m := range t.data {
		var row []string
		for j := range m {
			// fmt.Println(j, last-i)
			row = append(row, t.data[j][last-i])
		}
		data = append(data, row)
	}
	t.data = data
}

func (t *tile) flip() {
	var data [][]string
	var last = len(t.data) - 1
	for i, m := range t.data {
		var row []string
		for j := range m {
			// fmt.Println(j, last-i)
			row = append(row, t.data[i][last-j])
		}
		data = append(data, row)
	}
	t.data = data
}

func (t tile) top() []string {
	return t.data[0]
}

func (t tile) bottom() []string {
	return t.data[len(t.data)-1]
}

func (t tile) left() []string {
	var d []string
	for _, m := range t.data {
		d = append(d, m[0])
	}
	return d
}

func (t tile) right() []string {
	var d []string
	for _, m := range t.data {
		d = append(d, m[len(m)-1])
	}
	return d
}

func match(a []string, b []string) bool {
	for i, j := range a {
		if b[i] != j {
			return false
		}
	}

	return true
}

func (t *tile) matchSide(o *tile) bool {
	if match(t.top(), o.bottom()) {
		t.topTile = o.id
		o.bottomTile = t.id
		return true
	}

	if match(t.bottom(), o.top()) {
		t.bottomTile = o.id
		o.topTile = t.id
		return true
	}
	if match(t.left(), o.right()) {
		t.leftTile = o.id
		o.rightTile = t.id
		return true
	}
	if match(t.right(), o.left()) {
		t.rightTile = o.id
		o.leftTile = t.id
		return true
	}

	return false

}

func (t *tile) match(t2 *tile) {
	var found = false
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if t.matchSide(t2) {
				found = true
				break
			}

			if t.locked {
				break
			}

			if i&1 == 0 {
				t.flip()
			} else {
				t.flip()
				t.rotate()
			}

		}

		if t2.locked {
			break
		}

		if found {
			t.locked = true
			t2.locked = true
			break
		}

		if i&1 == 0 {
			t2.flip()
		} else {
			t2.flip()
			t2.rotate()
		}
	}

}

func (t *tile) stripe() {
	width := len(t.data[0]) - 2

	var data = make([][]string, width)
	for i := 0; i < width; i++ {
		data[i] = make([]string, width)
		for j := 0; j < width; j++ {
			data[i][j] = t.data[i+1][j+1]
			// fmt.Printf("(%d, %d)", i+1, j+1)
		}
		// fmt.Println()
	}
	t.data = data
}

func (t tile) print() {
	for _, m := range t.data {
		fmt.Println(strings.Join(m, ""))
	}
	fmt.Println()
}

func (t tile) printNeighbors() {
	fmt.Printf("%d t:%4d b:%4d l:%4d r:%4d\n", t.id, t.topTile, t.bottomTile, t.leftTile, t.rightTile)
}

func exist(id int64, ts []*tile) bool {
	for _, t := range ts {
		if t.id == id {
			return true
		}
	}
	return false
}

func (t tile) neighbors(tiles map[int64]*tile) []*tile {
	var ts []*tile
	if t.topTile != 0 {
		if tt := tiles[t.topTile]; !tt.visited {
			ts = append(ts, tiles[t.topTile])
		}
	}
	if t.bottomTile != 0 {
		if tt := tiles[t.bottomTile]; !tt.visited {
			ts = append(ts, tiles[t.bottomTile])
		}
	}
	if t.leftTile != 0 {
		if tt := tiles[t.leftTile]; !tt.visited {
			ts = append(ts, tiles[t.leftTile])
		}
	}
	if t.rightTile != 0 {
		if tt := tiles[t.rightTile]; !tt.visited {
			ts = append(ts, tiles[t.rightTile])
		}
	}
	return ts
}

type image struct {
	data [][]string
}

func NewImage(tiles [][]*tile) *image {
	var dLen = len(tiles[0][0].data[0])
	var width = len(tiles[0]) * dLen
	var data = make([][]string, width)
	for i := 0; i < width; i++ {
		data[i] = make([]string, width)
	}

	for i, row := range tiles {
		for j, col := range row {
			for k, d := range col.data {
				for m, c := range d {
					data[i*dLen+k][j*dLen+m] = c
				}
			}
		}
	}
	return &image{
		data: data,
	}
}

func (im image) findMonster(start []int, mons monster) bool {
	x, y := start[0], start[1]

	for _, m := range mons {
		delx, dely := m[0], m[1]
		if im.data[x+delx][y+dely] != "#" {
			return false
		}
	}
	return true
}

func (t *image) rotate() {
	var data [][]string
	var last = len(t.data) - 1
	for i, m := range t.data {
		var row []string
		for j := range m {
			// fmt.Println(j, last-i)
			row = append(row, t.data[j][last-i])
		}
		data = append(data, row)
	}
	t.data = data
}

func (t *image) flip() {
	var data [][]string
	var last = len(t.data) - 1
	for i, m := range t.data {
		var row []string
		for j := range m {
			// fmt.Println(j, last-i)
			row = append(row, t.data[i][last-j])
		}
		data = append(data, row)
	}
	t.data = data
}

func (t image) print() {
	for _, m := range t.data {
		fmt.Println(strings.Join(m, ""))
	}
	fmt.Println()
}

func (t image) countSymbol() (sum int) {
	for _, m := range t.data {
		for _, r := range m {
			if r == "#" {
				sum += 1
			}
		}
	}
	return
}

type monster [][]int

func parseSeaMonter(str string) monster {
	var data [][]int
	for i, r := range strings.Split(str, "\n") {
		for j, b := range strings.Split(r, "") {
			if b == "#" {
				data = append(data, []int{i, j})
			}
		}
	}

	return monster(data)
}

func main() {
	dat, err := ioutil.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n\n")
	var tiles []*tile
	var visited = make(map[int64]*tile)
	for _, line := range lines {
		tile := NewTile(line)
		tiles = append(tiles, tile)
		visited[tile.id] = tile
	}

	var imageWidth = int(math.Sqrt(float64(len(tiles))))
	var images = make([][]*tile, imageWidth)
	for i := 0; i < imageWidth; i++ {
		images[i] = make([]*tile, imageWidth)
	}

	// traverse tiles
	var nextRound = []*tile{tiles[0]}
	for len(nextRound) > 0 {
		// for i := 0; i < 2; i++ {
		var tmp []*tile
		for _, t1 := range nextRound {
			t1.visited = true
			for _, t2 := range tiles {
				if t1.id == t2.id {
					continue
				}
				t1.match(t2)
			}

			if t1.topTile == 0 && t1.leftTile == 0 {
				images[0][0] = t1
			}

			// part1
			// count t1.neighbors == 2

			// t1.printNeighbors()
			for _, k := range t1.neighbors(visited) {
				if !exist(k.id, tmp) {
					tmp = append(tmp, k)
				}
			}
		}
		// fmt.Println(tmp)
		nextRound = tmp
	}

	// Assemble image
	for i := 0; i < imageWidth; i++ {
		if i > 0 {
			id := images[i-1][0].bottomTile
			images[i][0] = visited[id]
		}
		for j := 0; j < imageWidth; j++ {
			if j == 0 {
				continue
			}
			id := images[i][j-1].rightTile
			images[i][j] = visited[id]
		}
	}

	for _, s := range images {
		for _, t := range s {
			fmt.Printf("%d ", t.id)
		}
		fmt.Println()
	}

	// stripe tiles
	for _, t := range tiles {
		t.stripe()
	}

	im := NewImage(images)
	im.print()
	// fmt.Println(im.countSymbol())

	seaMonster := "                  # \n" + "#    ##    ##    ###\n" + " #  #  #  #  #  #   "

	monster := parseSeaMonter(seaMonster)

	mwidth := 20
	mheight := 3
	var n = 0
	for i := 0; i < 8; i++ {
		for i := 0; i < (imageWidth*8 - mheight); i++ {
			for j := 0; j < (imageWidth*8 - mwidth); j++ {
				if im.findMonster([]int{i, j}, monster) {
					fmt.Println("yes", i, j)
					n += 1
				}

			}
		}

		if i&1 == 0 {
			im.flip()
		} else {
			im.flip()
			im.rotate()
		}
	}

	fmt.Println(im.countSymbol() - 15*n)
	// (0,0) ....(0,19)
	// (1,0)
	// (2,0)  ...

}
