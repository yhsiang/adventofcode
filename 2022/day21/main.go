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

func yell(inputs []string, part2 bool) int {
	monkeys := make(map[string]int)
	for len(inputs) > 0 {
		monkey := inputs[0]
		inputs = inputs[1:]
		if strings.Contains(monkey, "*") ||
			strings.Contains(monkey, "-") ||
			strings.Contains(monkey, "/") ||
			strings.Contains(monkey, "+") {
			var m string
			var a string
			var b string
			var op string
			fmt.Sscanf(strings.Replace(monkey, ":", "", 1), "%s %s %s %s", &m, &a, &op, &b)
			an, aok := monkeys[a]
			bn, bok := monkeys[b]
			mn, mok := monkeys[m]
			if part2 {
				if m == "root" {
					if aok {
						monkeys[b] = an
						continue
					} else if bok {
						monkeys[a] = bn
						continue
					}
					// fmt.Printf("%+v\n", monkeys)

				}
				// fmt.Printf("%s %s %s %d %d %d %+v %+v %+v\n", m, a, b, mn, an, bn, mok, aok, bok)
				if mok && aok {
					switch op {
					case "+":
						monkeys[b] = mn - an
					case "-":
						monkeys[b] = an - mn
					case "/":
						monkeys[b] = an / mn
					case "*":
						monkeys[b] = mn / an
					}
					continue
				}
				if mok && bok {
					switch op {
					case "+":
						monkeys[a] = mn - bn
					case "-":
						monkeys[a] = mn + bn
					case "/":
						monkeys[a] = mn * bn
					case "*":
						monkeys[a] = mn / bn
					}
					continue
				}
			}

			//part 1
			if !aok || !bok {
				inputs = append(inputs, monkey)
				continue
			}

			if aok && bok {
				switch op {
				case "+":
					monkeys[m] = an + bn
				case "-":
					monkeys[m] = an - bn
				case "/":
					monkeys[m] = an / bn
				case "*":
					monkeys[m] = an * bn
				}
			}

		} else {
			var m string
			var value int
			fmt.Sscanf(strings.Replace(monkey, ":", "", 1), "%s %d", &m, &value)
			if part2 {
				if m != "humn" {
					monkeys[m] = value
				}
			} else {
				monkeys[m] = value
			}
		}
	}

	if part2 {
		return monkeys["humn"]
	}

	return monkeys["root"]
}
func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}

	inputs := strings.Split(file, "\n")

	// fmt.Printf("%+v\n", monkeys)
	fmt.Printf("%d\n", yell(inputs, false))
	fmt.Printf("%d\n", yell(inputs, true))
}
