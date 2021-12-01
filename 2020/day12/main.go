package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type Ship struct {
	Face      int64
	PosX      int64
	PosY      int64
	WaypointX int64
	WaypointY int64
}

func (s *Ship) Move(action string) {
	tmp := strings.Split(action, "")
	step, _ := strconv.ParseInt(strings.Join(tmp[1:], ""), 10, 64)
	// fmt.Println(tmp[0], step)
	switch tmp[0] {
	case "N":
		s.PosY += step
	case "S":
		s.PosY -= step
	case "E":
		s.PosX += step
	case "W":
		s.PosX -= step
	case "L":
		s.Face += step
		if s.Face >= 180 {
			s.Face -= 360
		}
	case "R":
		s.Face -= step
		if s.Face <= -180 {
			s.Face += 360
		}
	case "F":
		switch s.Face {
		case 0: // E
			s.PosX += step
		case 90: // N
			s.PosY += step
		case -90: // S
			s.PosY -= step
		case 180, -180: // W
			s.PosX -= step
		}
	}
	fmt.Println(tmp[0], step, s)
}

func (s *Ship) MoveByWaypoint(action string) {
	tmp := strings.Split(action, "")
	num, _ := strconv.ParseInt(strings.Join(tmp[1:], ""), 10, 64)

	switch tmp[0] {
	case "F":
		s.PosX += s.WaypointX * num
		s.PosY += s.WaypointY * num
	case "N":
		s.WaypointY += num
	case "S":
		s.WaypointY -= num
	case "E":
		s.WaypointX += num
	case "W":
		s.WaypointX -= num
	case "L":
		// x := float64(s.WaypointX)*math.Cos(float64(num)*math.Pi/180) - float64(s.WaypointY)*math.Sin(float64(num)*math.Pi/180)
		// y := float64(s.WaypointY)*math.Cos(float64(num)*math.Pi/180) + float64(s.WaypointX)*math.Sin(float64(num)*math.Pi/180)
		// // fmt.Println(math.Round(x), math.Round(y))
		// s.WaypointX = int64(math.Round(x))
		// s.WaypointY = int64(math.Round(y))
		originX := s.WaypointX
		originY := s.WaypointY
		if num == 90 {
			s.WaypointX = -1 * originY
			s.WaypointY = originX
		}
		if num == 180 {
			s.WaypointX = -1 * s.WaypointX
			s.WaypointY = -1 * s.WaypointY
		}
		if num == 270 {
			s.WaypointX = originY
			s.WaypointY = -1 * originX
		}
	case "R":
		// 	x := float64(s.WaypointX)*math.Cos(float64(num)*math.Pi/180) + float64(s.WaypointY)*math.Sin(float64(num)*math.Pi/180)
		// 	y := float64(s.WaypointY)*math.Cos(float64(num)*math.Pi/180) - float64(s.WaypointX)*math.Sin(float64(num)*math.Pi/180)
		// 	// fmt.Println(math.Round(x), math.Round(y))
		// 	s.WaypointX = int64(math.Round(x))
		// 	s.WaypointY = int64(math.Round(y))
		originX := s.WaypointX
		originY := s.WaypointY
		if num == 90 {
			s.WaypointX = originY
			s.WaypointY = -1 * originX
		}
		if num == 180 {
			s.WaypointX = -1 * s.WaypointX
			s.WaypointY = -1 * s.WaypointY
		}

		if num == 270 {
			s.WaypointX = -1 * originY
			s.WaypointY = originX
		}
	}

	// fmt.Println(tmp[0], num, s)
}

func main() {
	dat, err := ioutil.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")

	// ship := Ship{
	// 	Face: 0, // east
	// 	PosX: 0,
	// 	PosY: 0,
	// }

	// Part 1
	// for _, action := range lines {
	// 	ship.Move(action)
	// }
	// fmt.Println(ship)

	ship := Ship{
		Face:      0, // east
		PosX:      0,
		PosY:      0,
		WaypointX: 10,
		WaypointY: 1,
	}
	for _, action := range lines {
		ship.MoveByWaypoint(action)
	}
	fmt.Println(math.Abs(float64(ship.PosX)) + math.Abs(float64(ship.PosY)))
}
