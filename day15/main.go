package main

import "fmt"

type Game struct {
	startLen int
	numbers  []int
	spoken   map[int][]int
}

func NewGame(start []int) *Game {
	var numbers = []int{0}
	var spoken = make(map[int][]int)
	for i, n := range start {
		spoken[n] = []int{i + 1}
	}
	return &Game{
		startLen: len(start),
		numbers:  append(numbers, start...),
		spoken:   spoken,
	}
}
func (g *Game) turn(n int) int {
	if n <= g.startLen {
		return g.numbers[n]
	}
	// fmt.Println(g)
	mostRecently := g.numbers[n-1]
	v := g.spoken[mostRecently]
	value := 0
	if len(v) != 1 {
		value = v[len(v)-1] - v[len(v)-2]
	}
	g.spoken[value] = append(g.spoken[value], n)
	g.numbers = append(g.numbers, value)
	return value
}

func (g *Game) calTurns(n int) int {
	var last int
	for i := 1; i <= n; i++ {
		last = g.turn(i)
	}
	return last
}

func main() {
	// Part1
	// g := NewGame([]int{0, 3, 6})
	// fmt.Println(g.calTurns(2020))
	// g = NewGame([]int{1, 3, 2})
	// fmt.Println(g.calTurns(2020))
	// g = NewGame([]int{2, 1, 3})
	// fmt.Println(g.calTurns(2020))
	// g = NewGame([]int{1, 2, 3})
	// fmt.Println(g.calTurns(2020))
	// g = NewGame([]int{2, 3, 1})
	// fmt.Println(g.calTurns(2020))
	// g = NewGame([]int{3, 2, 1})
	// fmt.Println(g.calTurns(2020))
	// g = NewGame([]int{3, 1, 2})
	// fmt.Println(g.calTurns(2020))
	// g = NewGame([]int{1, 20, 11, 6, 12, 0})
	// fmt.Println(g.calTurns(2020))

	// Part2
	g := NewGame([]int{0, 3, 6})
	fmt.Println(g.calTurns(30000000))
	g = NewGame([]int{1, 3, 2})
	fmt.Println(g.calTurns(30000000))
	g = NewGame([]int{2, 1, 3})
	fmt.Println(g.calTurns(30000000))
	g = NewGame([]int{1, 2, 3})
	fmt.Println(g.calTurns(30000000))
	g = NewGame([]int{2, 3, 1})
	fmt.Println(g.calTurns(30000000))
	g = NewGame([]int{3, 2, 1})
	fmt.Println(g.calTurns(30000000))
	g = NewGame([]int{3, 1, 2})
	fmt.Println(g.calTurns(30000000))
	g = NewGame([]int{1, 20, 11, 6, 12, 0})
	fmt.Println(g.calTurns(30000000))
}
