package main

import (
	_ "embed"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"

	util "github.com/yhsiang/adventofcode"
)

//go:embed example
var example string

//go:embed input
var input string

type Monkey struct {
	Items     []int64
	Operation string
	Test      int64
	TestTrue  int
	TestFalse int
}

func NewMonkey(input string) *Monkey {
	lines := strings.Split(input, "\n")
	itemLine := strings.Split(lines[1], ": ")
	itemStrs := strings.Split(itemLine[1], ", ")
	items := util.ToInt64(itemStrs)

	opLine := strings.Split(lines[2], ": ")
	testLine := strings.Split(lines[3], ": ")

	var testNum int64
	fmt.Sscanf(testLine[1], "divisible by %d", &testNum)

	var testTrue, testFalse int
	fmt.Sscanf(lines[4], "    If true: throw to monkey %d", &testTrue)
	fmt.Sscanf(lines[5], "    If false: throw to monkey %d", &testFalse)
	return &Monkey{
		Items:     items,
		Operation: opLine[1],
		Test:      testNum,
		TestTrue:  testTrue,
		TestFalse: testFalse,
	}
}

func (m *Monkey) operate(item int64) int64 {
	var op, num string
	fmt.Sscanf(m.Operation, "new = old %s %s", &op, &num)
	var answer int64
	switch op {
	case "*":
		if num == "old" {
			answer = item * item
		} else {
			n, _ := util.Int64(num)
			answer = item * n
		}
	case "+":
		if num == "old" {
			answer = item + item
		} else {
			n, _ := util.Int64(num)
			answer = item + n
		}
	}

	return answer
}

func (m *Monkey) test(item int64) int {
	if item%int64(m.Test) == 0 {
		return m.TestTrue
	}
	return m.TestFalse
}

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}

	monkeyInputs := strings.Split(file, "\n\n")

	var monkeys []*Monkey
	for _, monkey := range monkeyInputs {
		monkeys = append(monkeys, NewMonkey(monkey))
	}

	var inspects []int
	for i := 0; i < len(monkeys); i++ {
		inspects = append(inspects, 0)
	}
	round := 20
	for i := 1; i <= round; i++ {
		for monkeyIndex, monkey := range monkeys {
			itemSize := len(monkey.Items)
			k := 0
			for k < itemSize {
				inspects[monkeyIndex] += 1
				item := monkey.Items[0]
				newItem := monkey.operate(item)
				bored := int64(math.Floor(float64(newItem) / 3))
				target := monkey.test(bored)
				// fmt.Printf("%d %d\n", bored, target)
				k++
				monkey.Items = monkey.Items[1:]
				monkeys[target].Items = append(monkeys[target].Items, bored)
			}
		}
		// fmt.Printf("round %d\n", i)
		// for _, monkey := range monkeys {
		// 	fmt.Printf("%+v\n", monkey.Items)
		// }
	}

	// for _, v := range inspects {
	// 	fmt.Printf("%d\n", v)
	// }
	sort.Sort(sort.Reverse(sort.IntSlice(inspects)))
	// fmt.Printf("%+v\n", inspects)
	fmt.Printf("%d\n", inspects[0]*inspects[1])

	monkeys = []*Monkey{}
	inspects = []int{}
	for _, monkey := range monkeyInputs {
		monkeys = append(monkeys, NewMonkey(monkey))
	}
	for i := 0; i < len(monkeys); i++ {
		inspects = append(inspects, 0)
	}
	var testNums []int64
	for _, monkey := range monkeys {
		testNums = append(testNums, monkey.Test)
	}
	lcm := util.LCM(testNums[0], testNums[1], testNums[2:]...)

	round = 10000
	for i := 1; i <= round; i++ {
		for monkeyIndex, monkey := range monkeys {
			itemSize := len(monkey.Items)
			k := 0
			for k < itemSize {
				inspects[monkeyIndex] += 1
				item := monkey.Items[0]
				newItem := monkey.operate(item)
				// fmt.Printf("%d\n", newItem)
				bored := newItem % lcm

				target := monkey.test(bored)
				// fmt.Printf("%d\n", newItem)
				k++
				monkey.Items = monkey.Items[1:]
				monkeys[target].Items = append(monkeys[target].Items, bored)
				// fmt.Printf("%+v\n", monkeys[target])
			}
		}
	}
	// fmt.Printf("%+v", inspects)

	sort.Sort(sort.Reverse(sort.IntSlice(inspects)))
	// // fmt.Printf("%+v\n", inspects)
	fmt.Printf("%d\n", inspects[0]*inspects[1])
}
