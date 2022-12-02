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

/*
A X 3
  Y 6
  Z 0
B X 0
  Y 3
  Z 6
C X 6
  Y 0
  Z 3
*/

var scores = map[string]map[string]int{
	"A": {"X": 3, "Y": 6, "Z": 0},
	"B": {"X": 0, "Y": 3, "Z": 6},
	"C": {"X": 6, "Y": 0, "Z": 3},
}

var chosens = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
	"A": 1,
	"B": 2,
	"C": 3,
}

var strategies = map[string]int{
	"X": 0,
	"Y": 3,
	"Z": 6,
}

var strategyScores = map[string]map[int]string{
	"A": {0: "C", 3: "A", 6: "B"},
	"B": {0: "A", 3: "B", 6: "C"},
	"C": {0: "B", 3: "C", 6: "A"},
}

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}

	rounds := strings.Split(file, "\n")
	var total int
	var total2 int
	for _, round := range rounds {
		game := strings.Split(round, " ")
		score := chosens[game[1]] + scores[game[0]][game[1]]
		total += score
		strategy := strategies[game[1]]
		choose := strategyScores[game[0]][strategy]
		score2 := chosens[choose] + strategy
		total2 += score2
	}

	// first puzzle
	fmt.Printf("%d\n", total)
	// second puzzle
	fmt.Printf("%d\n", total2)
}
