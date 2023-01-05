package main

import (
	_ "embed"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

//go:embed example
var example string

//go:embed input
var input string

var mapping = map[string]int{
	"=": -2,
	"-": -1,
	"2": 2,
	"1": 1,
	"0": 0,
}
var revMapping = map[int]string{
	-1: "-",
	-2: "=",
	0:  "0",
}

func FromSnafu(snafu string) int {
	inputs := strings.Split(snafu, "")
	var output int
	for i, v := range inputs {
		place := len(inputs) - 1 - i
		p := int(math.Pow(5, float64(place)))

		output += p * mapping[v]
	}
	return output
}

func ToSnafu(num int) string {
	if num == 0 {
		return "0"
	}

	s := ""
	for q := num; q > 0; q = q / 5 {
		m := q % 5
		k := strconv.Itoa(m)
		if m == 3 {
			q += 2
			k = "="
		} else if m == 4 {
			q += 1
			k = "-"
		}

		s = fmt.Sprintf("%v%v", k, s)
	}

	return s
}

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}

	inputs := strings.Split(file, "\n")
	var sum int
	for _, v := range inputs {
		n := FromSnafu(v)
		sum += n
	}
	fmt.Printf("%d\n", sum)
	fmt.Printf("%s\n", ToSnafu(sum))
}
