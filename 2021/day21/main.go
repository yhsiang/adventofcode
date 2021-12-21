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

func parse(input string) int {
	s := strings.Split(input, "starting position: ")
	n, _ := util.Int(s[1])
	return n
}

func play(positions []int) int {
	var scores = make([]int, len(positions))
	var times = 0
	var dice = 0
	var points = make([]int, len(positions))
	copy(points, positions)
	var winner int
loop:
	for {
		for i := 0; i < len(points); i++ {
			var nums []int
			for k := 0; k < 3; k++ {
				dice += 1
				if dice > 100 {
					dice -= 100
				}
				nums = append(nums, dice)
			}
			times += 3
			sum := util.SumInt(nums)
			newPos := points[i] + sum

			for newPos > 10 {
				newPos %= 10
				if newPos == 0 {
					newPos = 10
				}
			}

			points[i] = newPos
			scores[i] += newPos
			if scores[i] >= 1000 {
				winner = i
				break loop
			}
		}
	}

	return scores[1-winner] * times
}

type player struct {
	space  int
	points int
}

func (p player) move(steps int) player {
	p.space = (p.space + steps) % 10
	if p.space == 0 {
		p.space = 10
	}
	p.points += p.space
	return p
}

// learn from https://github.com/jdrst/adventofgo/blob/main/2021/21/main.go
var quantumDiceOutcomes map[int]int = map[int]int{3: 1, 4: 3, 5: 6, 6: 7, 7: 6, 8: 3, 9: 1}

func playDirac(current, other player, p1Turn bool, universes int) (p1win int, p2win int) {
	if other.points > 20 {
		if p1Turn {
			return 0, universes
		}
		return universes, 0
	}

	for roll, additional := range quantumDiceOutcomes {
		p1wins, p2wins := playDirac(other, current.move(roll), !p1Turn, universes*additional)
		p1win += p1wins
		p2win += p2wins
	}

	return p1win, p2win
}

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}

	lines := util.ByLine(file)
	var positions []int
	for _, line := range lines {
		positions = append(positions, parse(line))
	}
	fmt.Printf("part1: %d\n", play(positions))

	// var outcomes = make(map[int]int)
	// for i := 1; i < 4; i++ {
	// 	for j := 1; j < 4; j++ {
	// 		for k := 1; k < 4; k++ {
	// 			outcomes[i+j+k] += 1
	// 		}
	// 	}
	// }
	// fmt.Println(outcomes)
	playerOne := player{positions[0], 0}
	playerTwo := player{positions[1], 0}
	p1, p2 := playDirac(playerOne, playerTwo, true, 1)
	fmt.Printf("part2: %d\n", util.Max(p1, p2))

}
