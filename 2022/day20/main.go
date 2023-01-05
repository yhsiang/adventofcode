package main

import (
	_ "embed"
	"fmt"
	"math"
	"os"
	"strings"

	util "github.com/yhsiang/adventofcode"
)

//go:embed example
var example string

//go:embed input
var input string

type Node struct {
	Value int
	Prev  *Node
	Next  *Node
}

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}

	strs := strings.Split(file, "\n")
	numbers := util.ToInt(strs)
	head := &Node{
		Value: numbers[0],
	}
	// var head *Node
	var prev *Node
	for i := 1; i < len(numbers); i++ {
		node := &Node{
			Value: numbers[i],
		}
		if i == 1 {
			node.Prev = head
			head.Next = node
			prev = node
		} else {
			prev.Next = node
			node.Prev = prev
			prev = node
		}
	}
	head.Prev = prev
	prev.Next = head

	tmp := head
	for i := 1; i <= 1000; i++ {
		idx := (i - 1) % len(numbers)
		// fmt.Printf("%+v\n", numbers[idx])
		for {
			// fmt.Printf("%+v\n", tmp.Value)
			if numbers[idx] == tmp.Value {
				break
			}
			tmp = tmp.Next
		}
		v := int(math.Abs(float64(tmp.Value)))
		h := tmp.Value
		for j := 0; j < v; j++ {
			// fmt.Printf("%d\n", j)
			if h > 0 {
				// 4, 1, 2, 3
				// 4, 2, 1, 3
				s := tmp.Next
				s.Prev = tmp.Prev
				s.Next.Prev = tmp
				tmp.Next = s.Next
				tmp.Prev.Next = s
				tmp.Prev = s
				s.Next = tmp
			} else {
				// 4, 2, -1, 3
				// 4, -1, 2, 3
				s := tmp.Prev
				tmp.Prev = s.Prev
				s.Prev.Next = tmp
				s.Prev = tmp
				s.Next = tmp.Next
				tmp.Next.Prev = s
				tmp.Next = s
			}
		}
		// fmt.Printf("%d\n", i)
		fmt.Printf("%d %d %d\n", tmp.Prev.Value, tmp.Value, tmp.Next.Value)
		// fmt.Printf("%+v", tmp)
		// fmt.Printf("%+v", tmp.Next)
		// fmt.Printf("%+v\n", tmp.Next.Next)
		// fmt.Printf("%+v\n", tmp.Next.Next.Next)
		// fmt.Printf("%+v\n", tmp.Next.Next.Next.Next)
		// fmt.Printf("%+v\n", tmp.Next.Next.Next.Next.Next)
		// fmt.Printf("%+v\n", tmp.Next.Next.Next.Next.Next.Next)
		// fmt.Printf("%+v\n", tmp.Next.Next.Next.Next.Next.Next.Next)
		if i == 1000 {

			s := tmp
			for {
				if s.Value == 0 {
					break
				}
				s = s.Next
			}
			fmt.Printf("%+v\n", s.Next)
		}

		if i == 2000 {

			s := tmp
			for {
				if s.Value == 0 {
					break
				}
				s = s.Next
			}
			fmt.Printf("%+v\n", s.Next)
		}

		if i == 3000 {

			s := tmp
			for {
				if s.Value == 0 {
					break
				}
				s = s.Next
			}
			fmt.Printf("%+v\n", s.Next)
		}

	}

}
