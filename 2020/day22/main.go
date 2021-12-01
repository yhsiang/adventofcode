package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type deck struct {
	cards []int
}

func (d *deck) copy(i int) *deck {
	var cards = make([]int, i)
	copy(cards, d.cards)
	return &deck{
		cards: cards,
	}
}

func (d *deck) pop() int {
	i := d.cards[0]
	d.cards = d.cards[1:]
	return i
}

func (d *deck) push(i ...int) {
	d.cards = append(d.cards, i...)
}

func (d *deck) isEmpty() bool {
	return len(d.cards) == 0
}

func (d *deck) String() string {
	var s []string
	for _, c := range d.cards {
		s = append(s, fmt.Sprintf("%d", c))
	}

	return strings.Join(s, ",")
}

func NewDeck(str string) *deck {
	data := strings.Split(str, "\n")
	var cards []int
	for _, d := range data[1:] {
		i, _ := strconv.ParseInt(d, 10, 64)
		cards = append(cards, int(i))
	}

	return &deck{
		cards: cards,
	}
}

func part1(p1, p2 *deck) (*deck, *deck) {
	for {
		if p1.isEmpty() || p2.isEmpty() {
			break
		}

		a := p1.pop()
		b := p2.pop()
		if a > b {
			p1.push(a, b)
		} else if b > a {
			p2.push(b, a)
		}
	}

	return p1, p2

}

func part2(p1, p2 *deck) (*deck, *deck, int) {
	var previous = make(map[string]struct{})
	// for {
	// for i := 0; i < 9; i++ {
	for len(p1.cards) > 0 && len(p2.cards) > 0 {
		key := fmt.Sprintf("%s %s", p1.String(), p2.String())
		if _, ok := previous[key]; ok {
			return p1, p2, 1
		}
		previous[key] = struct{}{}

		a := p1.pop()
		b := p2.pop()
		winner := 0
		if a > b {
			winner = 1
		} else if b > a {
			winner = 2
		}
		if len(p1.cards) >= a && len(p2.cards) >= b {
			_, _, winner = part2(p1.copy(a), p2.copy(b))
			// fmt.Println("winner", winner)
		}

		if winner == 1 {
			p1.push(a, b)
		} else if winner == 2 {
			p2.push(b, a)
		}

		// fmt.Println("p1", p1)
		// fmt.Println("p2", p2)
	}

	if p1.isEmpty() {
		return p1, p2, 2
	}

	return p1, p2, 1

}

func point(p1, p2 *deck) int {
	p := p1.cards
	if p1.isEmpty() {
		p = p2.cards
	}

	var sum int
	for i, v := range p {
		sum += v * (len(p) - i)
	}
	return sum
}

func main() {
	dat, err := ioutil.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n\n")
	p1, p2 := NewDeck(lines[0]), NewDeck(lines[1])

	// Part1
	// p1, p2, _ = part1(p1, p2, false)
	// fmt.Println(point(p1, p2))

	// Part2
	p1, p2, _ = part2(p1, p2)
	fmt.Println(point(p1, p2))
}
