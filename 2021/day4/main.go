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

type Tuple struct {
	X string
	Y string
}

type Board struct {
	Numbers []string
	Order   int
	Remove  bool
}

func initBoard(input string, order int) *Board {
	var nums []string
	for _, rows := range strings.Split(input, "\n") {
		for _, c := range strings.Split(rows, " ") {
			if c != "" {
				nums = append(nums, c)
			}
		}
	}
	return &Board{
		Numbers: nums,
		Order:   order,
		Remove:  false,
	}
}

func (b *Board) getIndex(num string) int {
	for i, d := range b.Numbers {
		if d == num {
			return i
		}
	}
	return -1
}

func (b *Board) marked(index int) {
	b.Numbers[index] = "x"
}

// only check two lines
func (b *Board) isBingo(index int) bool {
	base := index % 5
	times := index / 5
	// vertical
	var vertical = ""
	for i, j := base, 0; j < 5; i, j = i+5, j+1 {
		vertical += b.Numbers[i]
	}
	// horizontal
	var horizontal = ""
	for i := 5 * times; i < times*5+5; i++ {
		horizontal += b.Numbers[i]
	}

	return vertical == "xxxxx" || horizontal == "xxxxx"
}

func (b *Board) sum() int64 {
	var sum int64 = 0
	for _, num := range b.Numbers {
		if num != "x" {
			n, _ := util.Int64(num)
			sum += n
		}
	}
	return sum
}

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}

	d := strings.Split(file, "\n\n")
	numbers := strings.Split(d[0], ",")
	sample := d[1:]
	var boards []*Board
	for _, c := range sample {
		boards = append(boards, initBoard(c, 0))
	}

	var trigger string
	var lastBoard *Board
	for _, num := range numbers {
		if lastBoard != nil {
			break
		}
		for _, board := range boards {
			index := board.getIndex(num)
			if index == -1 {
				continue
			}
			board.marked(index)
			if board.isBingo(index) {
				trigger = num
				lastBoard = board
				break
			}

		}
	}
	n, _ := util.Int64(trigger)
	fmt.Printf("part1: %d\n", lastBoard.sum()*n)

	boards = []*Board{}
	for i, c := range sample {
		boards = append(boards, initBoard(c, i))
	}

	var count = 0
	for _, num := range numbers {
		if count == len(boards) {
			break
		}
		for _, board := range boards {
			if board.Remove {
				continue
			}
			index := board.getIndex(num)
			if index == -1 {
				continue
			}
			board.marked(index)

			if board.isBingo(index) {
				trigger = num
				lastBoard = board
				board.Remove = true
				count++
			}
		}
	}

	n, _ = util.Int64(trigger)
	fmt.Printf("part2: %d\n", lastBoard.sum()*n)
}
