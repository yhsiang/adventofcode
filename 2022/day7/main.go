package main

import (
	_ "embed"
	"fmt"
	"os"
	"sort"
	"strings"

	util "github.com/yhsiang/adventofcode"
)

//go:embed example
var example string

//go:embed input
var input string

type Node struct {
	Size     int
	Parent   *Node
	Children []*Node
}

type Nodes []*Node

func (v Nodes) Len() int           { return len(v) }
func (v Nodes) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }
func (v Nodes) Less(i, j int) bool { return v[i].Size < v[j].Size }

func traverse(current *Node, nodes *Nodes) {
	// fmt.Printf("%+v\n", current)
	*nodes = append(*nodes, current)
	for _, child := range current.Children {
		traverse(child, nodes)
	}
}

const TOTAL_SPACE = 70000000
const REQUIRED_SPACE = 30000000

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}

	inputs := strings.Split(file, "\n")

	root := &Node{
		Size:     0,
		Parent:   nil,
		Children: make([]*Node, 0),
	}
	current := root

	// trap：duplicated dir name
	for _, input := range inputs {
		if input[0] == '$' {
			switch input[2:4] {
			case "cd":
				if input[5:] == "/" {
					continue
				}
				// fmt.Printf("%s %+v\n", input, current)
				if input[5:] != ".." {
					child := &Node{
						Size:     0,
						Children: make([]*Node, 0),
						Parent:   current,
					}
					current.Children = append(current.Children, child)
					current = child
				} else {
					current = current.Parent
				}
			case "ls":
			}
		} else {
			if input[0:3] != "dir" {
				file := strings.Split(input, " ")
				size, _ := util.Int(file[0])
				current.Size += size
				parent := current.Parent
				for parent != nil {
					parent.Size += size
					parent = parent.Parent
				}
			}
		}
	}

	var nodes Nodes
	traverse(root, &nodes)

	sum := 0
	for _, node := range nodes {
		if node.Size < 100000 {
			sum += node.Size
		}
	}
	// first puzzle
	fmt.Printf("%d\n", sum)

	unused := TOTAL_SPACE - root.Size
	delete := REQUIRED_SPACE - unused
	// fmt.Printf("target space: %d\n", delete)
	sort.Sort(nodes)

	for _, node := range nodes {
		if node.Size >= delete {
			fmt.Printf("%d\n", node.Size)
			break
		}
	}
}
