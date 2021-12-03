package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"

	util "github.com/yhsiang/adventofcode"
)

//go:embed example
var example string

//go:embed input
var input string

func genRates(data []string) []map[string]int {
	var rates = []map[string]int{}
	// fmt.Printf("%+v\n", data)
	var l = len(data[0])
	for i := 0; i < l; i++ {
		rates = append(rates, map[string]int{
			"0": 0,
			"1": 0,
		})
	}
	for _, d := range data {
		for i, s := range strings.Split(d, "") {
			rates[i][s] += 1
		}
	}
	return rates
}

func mostCommon(rates []map[string]int) []string {
	var bits []string
	for _, d := range rates {
		if d["0"] > d["1"] {
			bits = append(bits, "0")
		} else {
			bits = append(bits, "1")
		}
	}
	return bits
}

func mostCommonIndex(rates []map[string]int, index int) string {
	var bits []string
	for _, d := range rates {
		if d["0"] > d["1"] {
			bits = append(bits, "0")
		} else {
			bits = append(bits, "1")
		}
	}
	return bits[index]
}

func leastCommon(rates []map[string]int) []string {
	var bits []string
	for _, d := range rates {
		if d["0"] < d["1"] || d["0"] == d["1"] {
			bits = append(bits, "0")
		} else {
			bits = append(bits, "1")
		}
	}
	return bits
}

func leastCommonIndex(rates []map[string]int, index int) string {
	var bits []string
	for _, d := range rates {
		if d["0"] < d["1"] || d["0"] == d["1"] {
			bits = append(bits, "0")
		} else {
			bits = append(bits, "1")
		}
	}
	return bits[index]
}

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}

	data := util.Read(file)
	var rates = genRates(data)

	var gamma = mostCommon(rates)

	var epsilon = leastCommon(rates)
	// reverse
	// for _, d := range gamma {
	// 	switch d {
	// 	case "0":
	// 		epsilon = append(epsilon, "1")
	// 	case "1":
	// 		epsilon = append(epsilon, "0")
	// 	}
	// }
	gammaNum, _ := util.BinToInt64(strings.Join(gamma, ""))
	epsilonNum, _ := util.BinToInt64(strings.Join(epsilon, ""))
	fmt.Printf("part1: %d\n", gammaNum*epsilonNum)

	target := data
	var result []string
	for i := 0; i < len(data[0]); i++ {
		if len(target) == 1 {
			break
		}
		result = []string{}
		o := mostCommonIndex(genRates(target), i)

		for _, d := range target {
			if string(d[i]) == o {
				result = append(result, d)
			}
		}
		target = result
	}

	oxygenNum, _ := util.BinToInt64(strings.Join(result, ""))
	target = data
	for i := 0; i < len(data[0]); i++ {
		if len(target) == 1 {
			break
		}
		result = []string{}
		o := leastCommonIndex(genRates(target), i)

		for _, d := range target {
			if string(d[i]) == o {
				result = append(result, d)
			}
		}
		target = result
	}

	co2Num, _ := util.BinToInt64(strings.Join(result, ""))
	fmt.Printf("part2: %d\n", oxygenNum*co2Num)
}
