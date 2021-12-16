package main

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"
)

//go:embed example
var example string

//go:embed input
var input string

func convert(hex string) string {
	var output string
	for _, f := range hex {
		i, _ := strconv.ParseUint(string(f), 16, 32)
		output += fmt.Sprintf("%04b", i)
	}
	return output
}

func parse(binary string) int {
	i, _ := strconv.ParseInt(binary, 2, 64)
	return int(i)
}

type Packet struct {
	Input   string
	Version int
	Type_   int
	Value   int
}

func decodePacket(packet *Packet) (*Packet, *Packet) {
	var next = &Packet{}
	packet.Version = parse(packet.Input[0:3])
	packet.Type_ = parse(packet.Input[3:6])
	var remain string
	switch packet.Type_ {
	case 4: // literal value
		x := 6
		var output string
		for {
			output += packet.Input[x+1 : x+5]
			if string(packet.Input[x]) == "0" {
				break
			}
			x += 5

		}
		remain = packet.Input[x+5:]
		packet.Value = parse(output)
	default: // operator
		i := string(packet.Input[6])
		switch i {
		case "0":
			x := 7
			y := x + 15
			remain = packet.Input[y:]

		case "1":
			x := 7
			y := x + 11
			remain = packet.Input[y:]
		}
	}

	next.Input = remain

	return packet, next

}

func decode(data string, n int) (int, int) {
	_ = data[n : n+3]
	type_ := parse(data[n+3 : n+6])
	i := string(data[n+6])

	if type_ == 4 { // literal
		x := n + 6
		var output string
		for string(data[x]) == "1" {
			output += data[x+1 : x+5]
			x += 5
		}
		output += data[x+1 : x+5]
		x += 5
		return x, parse(output)
	}

	if i == "0" {
		var vals []int
		var s int
		start := n + 22
		len := parse(data[n+7 : n+22])
		end := start + len
		for start < end {
			start, s = decode(data, start)
			vals = append(vals, s)
		}
		return start, operate(vals, type_)
	}

	var vals []int
	var s int
	packs := parse(data[n+7 : n+18])
	start := n + 18
	for k := 0; k < packs; k++ {
		start, s = decode(data, start)
		vals = append(vals, s)
	}
	return start, operate(vals, type_)
}

func sum(vals []int) int {
	var num int
	for _, d := range vals {
		num += d
	}
	return num
}

func mul(vals []int) int {
	var num int = 1
	for _, d := range vals {
		num *= d
	}
	return num
}

func min(vals []int) int {
	var num int = vals[0]
	for _, d := range vals {
		if d < num {
			num = d
		}
	}
	return num
}

func max(vals []int) int {
	var num int = vals[0]
	for _, d := range vals {
		if d > num {
			num = d
		}
	}
	return num
}

func operate(vals []int, type_ int) (returnValue int) {
	if type_ == 0 {
		returnValue = sum(vals)
	}

	if type_ == 1 {
		returnValue = mul(vals)
	}

	if type_ == 2 {
		returnValue = min(vals)
	}
	if type_ == 3 {
		returnValue = max(vals)
	}
	if type_ == 5 {
		if vals[0] > vals[1] {
			returnValue = 1
		} else {
			returnValue = 0
		}
	}
	if type_ == 6 {
		if vals[0] < vals[1] {
			returnValue = 1
		} else {
			returnValue = 0
		}
	}

	if type_ == 7 {
		if vals[0] == vals[1] {
			returnValue = 1
		} else {
			returnValue = 0
		}
	}
	return
}

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}

	bs := convert(file)
	p := &Packet{
		Input: bs,
	}
	temp := p
	var sum int
	var cur *Packet
	for {
		cur, temp = decodePacket(temp)
		sum += cur.Version
		if len(temp.Input) < 11 {
			break
		}
	}
	fmt.Printf("part1: %d\n", sum)
	_, value := decode(bs, 0)
	fmt.Printf("part2: %d\n", value)
}
