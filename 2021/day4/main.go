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

type Card struct {
	Numbers [][]string
	Marked  [][]string
	Order   int
}

func initCard(input string, order int) *Card {
	var numbers [][]string
	for _, s := range strings.Split(input, "\n") {
		var temp []string
		for _, c := range strings.Split(s, " ") {
			if c != "" {
				temp = append(temp, c)
			}
		}
		numbers = append(numbers, temp)
	}
	var marked [][]string
	for _, rows := range numbers {
		var temp = []string{}
		for range rows {
			temp = append(temp, "")
		}
		marked = append(marked, temp)
	}
	return &Card{
		Numbers: numbers,
		Marked:  marked,
		Order:   order,
	}
}

func reverse(input [][]string) (output [][]string) {
	for _, rows := range input {
		var temp = []string{}
		// fmt.Printf("%d", len(rows))
		for range rows {
			// fmt.Printf("%s\n", c)
			temp = append(temp, "")
		}
		output = append(output, temp)
	}
	for i, rows := range input {
		for j, c := range rows {
			output[j][i] = c
		}
	}

	return output
}

// debug
func print(input [][]string) {
	for _, rows := range input {
		for _, col := range rows {
			fmt.Printf("%s ", col)
		}
		fmt.Printf("\n")
	}
}

func (c *Card) mark(input string) {
	for i, rows := range c.Numbers {
		for j, col := range rows {
			if col == input {
				c.Marked[i][j] = "x"
			}
		}
	}
}

func (c *Card) sum() int64 {
	var total int64 = 0
	for i, rows := range c.Numbers {
		for j, col := range rows {
			if c.Marked[i][j] != "x" {
				n, _ := util.Int64(col)
				total += n
			}
		}
	}
	return total
}

func (c *Card) reset() {
	var marked [][]string
	for _, rows := range c.Marked {
		var temp = []string{}
		for range rows {
			temp = append(temp, "")
		}
		marked = append(marked, temp)
	}
	c.Marked = marked
}

func (c *Card) isBingo() bool {
	// horizontal
	for _, rows := range c.Marked {
		if strings.Join(rows, "") == "xxxxx" {
			return true //, fmt.Sprintf("row %d", i)
		}
	}
	// vertical
	vcard := reverse(c.Marked)
	for _, rows := range vcard {
		if strings.Join(rows, "") == "xxxxx" {
			return true //, fmt.Sprintf("col %d", i)
		}
	}
	// cross := [][]string{
	// 	{
	// 		c.Card[0][0],
	// 		c.Card[1][1],
	// 		c.Card[2][2],
	// 		c.Card[3][3],
	// 		c.Card[4][4],
	// 	},
	// 	{
	// 		c.Card[0][4],
	// 		c.Card[1][3],
	// 		c.Card[2][2],
	// 		c.Card[3][1],
	// 		c.Card[4][0],
	// 	},
	// }
	// for _, c := range cross {
	// 	if strings.Join(c, "") == "xxxxx" {
	// 		return true
	// 	}
	// }
	return false //, ""
}

func exist(nums []int, i int) bool {
	for _, c := range nums {
		if i == c {
			return true
		}
	}
	return false
}

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}

	d := strings.Split(file, "\n\n")
	numbers := strings.Split(d[0], ",")
	sample := d[1:]
	var cards []*Card
	for i, c := range sample {
		cards = append(cards, initCard(c, i))
	}

	var lastCard *Card
	var trigger string
	for _, num := range numbers {
		if lastCard != nil {
			break
		}
		for _, card := range cards {
			card.mark(num)
			if card.isBingo() {
				trigger = num
				lastCard = card
				break
			}
		}
	}

	n, _ := util.Int64(trigger)
	fmt.Printf("part1: %d\n", lastCard.sum()*n)

	for _, card := range cards {
		card.reset()
	}

	lastCard = nil
	var removed []int
	for _, num := range numbers {
		if len(removed) == len(cards) {
			break
		}
		for _, card := range cards {
			if exist(removed, card.Order) {
				continue
			}
			card.mark(num)
			if card.isBingo() {
				trigger = num
				lastCard = card
				removed = append(removed, card.Order)
			}
		}
	}
	n, _ = util.Int64(trigger)
	fmt.Printf("part2: %d\n", lastCard.sum()*n)
}
