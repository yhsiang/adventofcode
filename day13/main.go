package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Schedule struct {
	bus []int64
	num int
}

func NewSchedule(str string) *Schedule {
	busSlice := strings.Split(str, ",")
	var busIDs []int64
	var num int
	for _, bus := range busSlice {
		if bus == "x" {
			busIDs = append(busIDs, 0)
		} else {
			b, _ := strconv.ParseInt(bus, 10, 64)
			busIDs = append(busIDs, b)
			num++
		}
	}

	return &Schedule{
		bus: busIDs,
		num: num,
	}
}

// func (s Schedule) findEarliestTime() int64 {
// 	var t int64 = 100000000000003 //s.bus[0]
// 	var find int
// 	for {
// 		fmt.Println(t)
// 		for i := range s.bus {
// 			if s.bus[i] == 0 {
// 				continue
// 			}
// 			sum := t + int64(i)
// 			if sum%s.bus[i] == 0 {
// 				find++
// 			}
// 		}

// 		if find == s.num {
// 			break
// 		}
// 		find = 0
// 		t += s.bus[0]
// 	}
// 	return t
// }

func solve(a, b, r int64) int64 {
	var x int64 = 1
	if r < 0 {
		r += b
	}
	for {
		if x*a%b == r {
			break
		}
		x++
	}
	return x
}

// crt is for Chinese Reminder Theorem
func crt(mods []int64, reminders []int64) int64 {
	var prod int64 = 1
	var sum int64 = 0
	for _, n := range mods {
		prod *= n
	}
	for i, n := range mods {
		var p = prod / n
		s := solve(p, n, reminders[i])
		// fmt.Println(s)
		sum += s * p
		// sum = sum + reminders[i] + mulInv(p, n)
	}

	for {
		if sum-prod > 0 {
			sum -= prod
		} else {
			break
		}
	}

	return sum
}

func (s Schedule) findEarliestTime() int64 {
	var mods []int64
	var reminders []int64
	for i := range s.bus {
		if s.bus[i] == 0 {
			continue
		}
		mods = append(mods, s.bus[i])
		reminders = append(reminders, -1*int64(i)%s.bus[i])
	}

	// fmt.Println(mods, reminders)
	return crt(mods, reminders)
}

func main() {
	// dat, err := ioutil.ReadFile("./input")
	// if err != nil {
	// 	panic(err)
	// }

	// lines := strings.Split(string(dat), "\n")
	// Part1
	// timestamp, _ := strconv.ParseInt(lines[0], 10, 64)
	// fmt.Println(timestamp)
	// busSlice := strings.Split(lines[1], ",")
	// var busIDs []int64
	// for _, bus := range busSlice {
	// 	if bus != "x" {
	// 		b, _ := strconv.ParseInt(bus, 10, 64)
	// 		busIDs = append(busIDs, b)
	// 	}
	// }
	// var start = timestamp
	// var find bool
	// var id int64
	// for {
	// 	for _, bus := range busIDs {
	// 		if start%bus == 0 {
	// 			id = bus
	// 			find = true
	// 			break
	// 		}
	// 	}
	// 	if find {
	// 		break
	// 	}
	// 	start++
	// }
	// fmt.Println((start - timestamp) * id)

	s := NewSchedule("19,x,x,x,x,x,x,x,x,41,x,x,x,x,x,x,x,x,x,523,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,17,13,x,x,x,x,x,x,x,x,x,x,29,x,853,x,x,x,x,x,37,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,23")

	fmt.Println(s.findEarliestTime())

}
