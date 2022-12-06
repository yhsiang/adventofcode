package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
)

//go:embed example
var example string

//go:embed input
var input string

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}

	strs := strings.Split(file, "\n")

	for _, str := range strs {
		var start = 0
		var end = 4 // market
		cs := strings.Split(str, "")
		// fmt.Printf("%+v", cs[start:end])
		for {
			var maps = make(map[string]int)
			for _, c := range cs[start:end] {
				if _, ok := maps[c]; !ok {
					maps[c] = 0
				}
			}
			if len(maps) == 4 {
				break
			}
			start += 1
			end += 1
		}
		// first puzzle
		fmt.Printf("%d\n", end)

	}

	for _, str := range strs {
		var start = 0
		var end = 14 // market
		cs := strings.Split(str, "")
		// fmt.Printf("%+v", cs[start:end])
		for {
			var maps = make(map[string]int)
			for _, c := range cs[start:end] {
				if _, ok := maps[c]; !ok {
					maps[c] = 0
				}
			}
			if len(maps) == 14 {
				break
			}
			start += 1
			end += 1
		}
		fmt.Printf("%d\n", end)
	}

}
