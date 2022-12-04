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

// a:97 z: 122 -> (-96) -> 1:26
// A:65 Z:90 -> (-38) -> 27:52
func toCode(s string) int {
	var charCode = int([]rune(s)[0])
	if charCode >= 97 {
		return charCode - 96
	}
	return charCode - 38
}

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}

	rucksacks := strings.Split(file, "\n")

	var sum int
	for _, rucksack := range rucksacks {
		half := len(rucksack) / 2
		one := rucksack[:half]
		two := rucksack[half:]
		var both string
		for _, s := range strings.Split(one, "") {
			if strings.Contains(two, s) {
				both = s
				break
			}
		}
		sum += toCode(both)
	}

	fmt.Printf("%d\n", sum)

	var group []string
	var sum2 int
	for i, rucksack := range rucksacks {
		group = append(group, rucksack)
		if i%3 == 2 {
			var badge string
			for _, s := range strings.Split(group[0], "") {
				if strings.Contains(group[1], s) && strings.Contains(group[2], s) {
					badge = s
					break
				}
			}
			sum2 += toCode(badge)
			group = []string{}
		}
	}

	fmt.Printf("%d\n", sum2)
}
