package main

import (
	"fmt"
	"strconv"
	"strings"
)

type crab struct {
	cups []int
}

func NewCrab(str string) *crab {
	var cups []int
	for _, s := range strings.Split(str, "") {
		i, _ := strconv.ParseInt(s, 10, 64)
		cups = append(cups, int(i))
	}

	return &crab{cups: cups}
}

func exist(num int, numbers []int) bool {
	for _, n := range numbers {
		if n == num {
			return true
		}
	}
	return false
}

func find(num int, numbers []int) int {
	for i, n := range numbers {
		if n == num {
			return i
		}
	}
	return -1
}

func (c *crab) pop(idx int) int {
	a := c.cups[idx]
	var cups []int
	for _, r := range c.cups {
		if r != a {
			cups = append(cups, r)
		}
	}
	c.cups = cups
	return a
}

func (c *crab) append(n int) {
	c.cups = append(c.cups, n)
}

func (c *crab) insert(idx int, cup int) {
	var cups []int
	cups = append(cups, c.cups[:idx]...)
	cups = append(cups, cup)
	cups = append(cups, c.cups[idx:]...)
	c.cups = cups
}

func (c *crab) move(n int) {
	// fmt.Println(c.cups[0])
	for i := 0; i < n; i++ {
		var picked []int
		picked = append(picked, c.pop(1))
		picked = append(picked, c.pop(1))
		picked = append(picked, c.pop(1))
		// fmt.Println(picked)
		dest := c.cups[0] - 1
		for dest <= 0 || exist(dest, picked) {
			if dest <= 0 {
				dest = 9
			} else {
				dest -= 1
			}
		}
		// fmt.Println(dest)
		index := find(dest, c.cups)

		for j, cup := range picked {
			c.insert(index+j+1, cup)
		}
		c.append(c.pop(0))
	}
	// fmt.Println(c.cups)
}

type cup struct {
	value int
	next  *cup
}

func NewCup(value int) *cup {
	return &cup{value: value, next: nil}
}

func main() {
	// Part1
	// crab := NewCrab("784235916")
	// crab.move(100)
	// idx := find(1, crab.cups)
	// var cups []int
	// cups = append(cups, crab.cups[idx+1:]...)
	// cups = append(cups, crab.cups[:idx]...)
	// var s []string
	// for _, c := range cups {
	// 	s = append(s, fmt.Sprintf("%d", c))
	// }
	// fmt.Println(strings.Join(s, ""))

	cups := make(map[int]*cup)
	// init := []int{3, 8, 9, 1, 2, 5, 4, 6, 7}
	init := []int{7, 8, 4, 2, 3, 5, 9, 1, 6}
	prev := init[0]
	for _, c := range init {
		cups[c] = NewCup(c)
		cups[prev].next = cups[c]
		prev = c
	}

	for i := 10; i <= 1000000; i++ {
		cups[i] = NewCup(i)
		cups[prev].next = cups[i]
		prev = i
	}
	cups[1000000].next = cups[init[0]]

	current := cups[init[0]]
	for i := 0; i < 10000000; i++ {
		c1 := current.next
		c2 := c1.next
		c3 := c2.next

		dest := current.value - 1
		for dest <= 0 || exist(dest, []int{c1.value, c2.value, c3.value}) {
			if dest <= 0 {
				dest = 1000000
			} else {
				dest -= 1
			}
		}

		left := cups[dest]
		// fmt.Println(i, dest, left)
		right := left.next

		successor := c3.next
		current.next = successor
		left.next = c1
		c3.next = right

		current = successor
	}
	star := cups[1]
	fmt.Println(star.next.value * star.next.next.value)

}
