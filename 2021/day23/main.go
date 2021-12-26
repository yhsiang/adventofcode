package main

import (
	"container/heap"
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

// learn from https://www.reddit.com/r/adventofcode/comments/rmnozs/comment/hpzkhv2/?utm_source=reddit&utm_medium=web2x&context=3
type state struct {
	grid [][]string
	cost int
}

func (s *state) Value() int {
	return s.cost
}

func initState(input string) *state {
	grid := [][]string{}
	lines := util.ByLine(input)
	for _, line := range lines {
		grid = append(grid, strings.Split(line, ""))
	}

	return &state{
		grid: grid,
	}
}

var part1 = map[[2]int]string{
	{2, 3}: "A", {3, 3}: "A",
	{2, 5}: "B", {3, 5}: "B",
	{2, 7}: "C", {3, 7}: "C",
	{2, 9}: "D", {3, 9}: "D",
}

var part2 = map[[2]int]string{
	{2, 3}: "A", {3, 3}: "A", {4, 3}: "A", {5, 3}: "A",
	{2, 5}: "B", {3, 5}: "B", {4, 5}: "B", {5, 5}: "B",
	{2, 7}: "C", {3, 7}: "C", {4, 7}: "C", {5, 7}: "C",
	{2, 9}: "D", {3, 9}: "D", {4, 9}: "D", {5, 9}: "D",
}

func (s *state) finish(result map[[2]int]string) bool {
	for coord, amphipod := range result {
		if s.grid[coord[0]][coord[1]] != amphipod {
			return false
		}
	}
	return true
}

func (s *state) getUnsettledCoords(result map[[2]int]string) [][2]int {
	var unsettled [][2]int
	// check entire hallway
	for col := 1; col < len(s.grid[0]); col++ {
		if strings.Contains("ABCD", s.grid[1][col]) {
			unsettled = append(unsettled, [2]int{1, col})
		}
	}

	for _, col := range []int{3, 5, 7, 9} {
		roomFullFromBack := true
		for row := len(s.grid) - 2; row >= 2; row-- {
			coord := [2]int{row, col}
			wantChar := result[coord]
			gotChar := s.grid[row][col]
			if gotChar != "." {
				if gotChar != wantChar {
					roomFullFromBack = false
					unsettled = append(unsettled, coord)
				} else if gotChar == wantChar && !roomFullFromBack {
					// need to get out of the way of someone in the wrong room
					unsettled = append(unsettled, coord)
				}
			}
		}
	}
	return unsettled
}

func isInHallway(coord [2]int) bool {
	return coord[0] == 1
}

var cannotStop = map[[2]int]bool{
	{1, 3}: true,
	{1, 5}: true,
	{1, 7}: true,
	{1, 9}: true,
}

func (s *state) getNextPossibleMoves(unsettledCoord [2]int, roomCoordToWantChar map[[2]int]string) [][2]int {
	// get all the eligible locations for this coord to go to
	unsettledChar := s.grid[unsettledCoord[0]][unsettledCoord[1]]

	if !strings.Contains("ABCD", unsettledChar) {
		panic("unexpected character to get next moves for " + unsettledChar)
	}

	var possible [][2]int

	startedInHallway := isInHallway(unsettledCoord)

	queue := [][2]int{unsettledCoord}
	seen := map[[2]int]bool{}
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]

		if seen[front] {
			continue
		}
		seen[front] = true

		if front != unsettledCoord {
			// is not a coord in front of a room
			if !cannotStop[front] {
				wantChar, isRoomCoord := roomCoordToWantChar[front]
				// if NOT in a room, append it
				if !isRoomCoord {
					// ONLY add a hallway if it started in a room bc of rule 3
					if !startedInHallway {
						possible = append(possible, front)
					}
				} else if wantChar == unsettledChar {
					// found the correct room
					// check if there is a deeper part of the room (aka lower)

					// if there is a "stuck" amphipod deeper in the room, cannot stop here
					// if not deepest empty coord, cannot stop here
					// in both cases walking further is handles all cases, whether that's
					//   to walk further in or out of the room
					isStuckAmphipod := false
					roomHasDeeperOpenSpaces := false
					for r := front[0] + 1; r < len(s.grid)-1; r++ {
						char := s.grid[r][front[1]]
						if char == "." {
							roomHasDeeperOpenSpaces = true
						}
						if char != "." && char != unsettledChar {
							isStuckAmphipod = true
							break
						}
					}

					if !roomHasDeeperOpenSpaces && !isStuckAmphipod {
						possible = append(possible, front)
					}
				}
			}
		}

		for _, d := range [][2]int{
			// up down left right
			{-1, 0},
			{1, 0},
			{0, -1},
			{0, 1},
		} {
			// do not need to check in range because the entire walkable area is surrounded by walls
			next := [2]int{front[0] + d[0], front[1] + d[1]}
			if s.grid[next[0]][next[1]] == "." {
				// add to queue to keep walking regardless of whether or not it gets added to the possible slice
				queue = append(queue, next)
			}
		}
	}

	return possible
}

func sumCost(char string, start, end [2]int) int {
	// start with cols distance
	dist := util.Abs(end[1] - start[1])
	// add distance to hallway for start and end?
	dist += start[0] - 1
	dist += end[0] - 1

	energyPerType := map[string]int{
		"A": 1,
		"B": 10,
		"C": 100,
		"D": 1000,
	}

	if _, ok := energyPerType[char]; !ok {
		panic(char + " should not call calcEnergy()")
	}
	return energyPerType[char] * dist
}

func (s *state) copy() *state {
	cp := state{
		grid: make([][]string, len(s.grid)),
		cost: s.cost,
	}

	// need to directly copy grid or else underlying arrays will be the same & interfere
	for i := range cp.grid {
		cp.grid[i] = make([]string, len(s.grid[i]))
		copy(cp.grid[i], s.grid[i])
	}

	return &cp
}

func run(input string, isPart2 bool) int {
	start := initState(input)
	result := part1

	if isPart2 {
		result = part2
		start.grid = append(start.grid, nil, nil)
		start.grid[6] = start.grid[4]
		start.grid[5] = start.grid[3]

		start.grid[3] = strings.Split("  #D#C#B#A#  ", "")
		start.grid[4] = strings.Split("  #D#B#A#C#  ", "")
	}

	minHeap := &util.MinHeap{}
	heap.Init(minHeap)
	heap.Push(minHeap, start)

	seen := make(map[string]bool)
	for minHeap.Len() > 0 {
		current := heap.Pop(minHeap).(*state)
		key := fmt.Sprint(current.grid)
		if seen[key] {
			continue
		}
		seen[key] = true

		if current.finish(result) {
			return current.cost
		}

		unsettledCoords := current.getUnsettledCoords(result)
		for _, unsettledCoord := range unsettledCoords {
			ur, uc := unsettledCoord[0], unsettledCoord[1]
			nextMoves := current.getNextPossibleMoves(unsettledCoord, result)
			for _, nextCoord := range nextMoves {
				nr, nc := nextCoord[0], nextCoord[1]
				if current.grid[nr][nc] != "." {
					panic(fmt.Sprintf("should only be moving to walkable spaces, got %q at %d,%d", current.grid[nr][nc], nr, nc))
				}

				cp := current.copy()
				// add the energy that will be used, swap the two coords
				cp.cost += sumCost(cp.grid[ur][uc], unsettledCoord, nextCoord)
				// cp.path += fmt.Sprintf("%s%v->%v{%d},", current.grid[ur][uc], unsettledCoord, nextCoord, cp.energyUsed)
				cp.grid[nr][nc], cp.grid[ur][uc] = cp.grid[ur][uc], cp.grid[nr][nc]

				// add it to the min heap
				minHeap.Push(cp)
			}
		}
	}

	panic("failed")
}

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}

	fmt.Printf("part1: %d\n", run(file, false))
	fmt.Printf("part2: %d\n", run(file, true))
}
