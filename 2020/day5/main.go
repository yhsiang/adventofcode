package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type Seat struct {
	Column int
	Row    int
	ID     int
}

func NewSeat(str string) *Seat {
	rows := str[:len(str)-3]
	var start = 0
	var end = 127
	for _, row := range strings.Split(rows, "") {
		switch row {
		case "F":
			end = start + (end-start)/2
		case "B":
			start = end - (end-start)/2
		}
	}

	cols := str[len(str)-3:]
	var cStart = 0
	var cEnd = 7
	for _, col := range strings.Split(cols, "") {
		switch col {
		case "R":
			cStart = cEnd - (cEnd-cStart)/2
		case "L":
			cEnd = cStart + (cEnd-cStart)/2
		}
		// fmt.Println(cStart, cEnd)
	}

	return &Seat{
		Row:    start,
		Column: cStart,
		ID:     start*8 + cStart,
	}
}

func main() {
	dat, err := ioutil.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")

	var seats []Seat
	for _, line := range lines {
		seats = append(seats, *NewSeat(line))
	}

	// var s Seat
	// for _, seat := range seats {
	// 	fmt.Println(seat)
	// 	if seat.ID > s.ID {
	// 		s = seat
	// 	}
	// }

	sort.Slice(seats, func(i, j int) bool {
		return seats[i].ID > seats[j].ID
	})

	for _, seat := range seats {
		if seat.ID > 606 && seat.ID < 654 {
			fmt.Println(seat)
		}
	}
}
