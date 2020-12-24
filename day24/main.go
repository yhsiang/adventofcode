package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func parseLine(str string) []string {
	s := strings.Split(str, "")
	i := 0
	var directions []string
	for i < len(s) {
		switch s[i] {
		case "s", "n":
			directions = append(directions, fmt.Sprintf("%s%s", s[i], s[i+1]))
			i++
		case "e", "w":
			directions = append(directions, s[i])
		}
		i++
	}
	return directions
}

func traceTiles(directions []string) (float64, float64) {
	var x, y float64
	for _, d := range directions {
		switch d {
		case "e":
			x += 1
		case "w":
			x -= 1
		case "se":
			y -= 0.5
			x += 0.5
		case "sw":
			y -= 0.5
			x -= 0.5
		case "ne":
			y += 0.5
			x += 0.5
		case "nw":
			y += 0.5
			x -= 0.5
		}
	}

	return x, y
}

type hexagonal struct {
	x float64
	y float64
}

type state struct {
	times int
	side  string
}

func count(hexagonals map[hexagonal]state) (n int) {
	for _, hex := range hexagonals {
		if hex.side == "black" {
			n += 1
		}
	}
	return
}

func main() {
	dat, err := ioutil.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")

	var hexagonals = make(map[hexagonal]state)
	for _, line := range lines {
		ds := parseLine(line)
		x, y := traceTiles(ds)
		hex := hexagonal{x: x, y: y}
		s, ok := hexagonals[hex]
		if !ok {
			s = state{
				side:  "white",
				times: 0,
			}
		}
		s.times += 1
		if s.side == "white" {
			s.side = "black"
		} else {
			s.side = "white"
		}
		hexagonals[hex] = s
	}

	// part1
	// var n int
	// for _, hex := range hexagonals {
	// 	if hex.side == "black" {
	// 		n += 1
	// 	}
	// }
	// fmt.Println(n)
	// fmt.Println(hexagonals)
	var adjacencies = [][]float64{
		{1, 0},
		{-1, 0},
		{0.5, 0.5},
		{-0.5, 0.5},
		{0.5, -0.5},
		{-0.5, -0.5},
	}
	days := 100
	var tmp = make(map[hexagonal]state)
	for h, s := range hexagonals {
		if s.side == "black" {
			tmp[h] = s
		}
	}
	hexagonals = tmp

	for i := 0; i < days; i++ {
		var tmp = make(map[hexagonal]state)
		var whiteBlack = make(map[hexagonal]int)
		for hex, s := range hexagonals {
			var blackNeighbours int
			for _, adj := range adjacencies {
				h := hexagonal{x: hex.x + adj[0], y: hex.y + adj[1]}
				if _, ok := hexagonals[h]; ok {
					blackNeighbours += 1
				} else {
					if _, ok := whiteBlack[h]; !ok {
						whiteBlack[h] = 0
					}
					whiteBlack[h] += 1
				}
			}
			if blackNeighbours == 0 || blackNeighbours > 2 {
				continue
			} else {
				tmp[hex] = s
			}
		}

		for k, v := range whiteBlack {
			if v == 2 {
				tmp[k] = state{side: "black"}
			}
		}
		fmt.Println(count(tmp))
		hexagonals = tmp
	}

}
