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
	Version int
	Type_   int
	Value   int
	Length  int

	// operator
	LengthTypeId  int
	LengthTypeVal int
}

func decodePacket(data string) (*Packet, string) {
	var packet = &Packet{}
	packet.Version = parse(data[0:3])
	packet.Type_ = parse(data[3:6])
	var remain string
	switch packet.Type_ {
	case 4: // literal value
		x := 6
		var output string
		for {
			output += data[x+1 : x+5]
			if string(data[x]) == "0" {
				break
			}
			x += 5

		}
		remain = data[x+5:]
		packet.Value = parse(output)
		packet.Length = x + 5
	default: // operator
		i := string(data[6])
		switch i {
		case "0":
			x := 7
			y := x + 15
			remain = data[y:]
			packet.LengthTypeId = 1
			packet.Length = 22
			packet.LengthTypeVal = parse(data[7:22])
		case "1":
			x := 7
			y := x + 11
			remain = data[y:]
			packet.LengthTypeId = 2
			packet.Length = 18
			packet.LengthTypeVal = parse(data[7:18])
		}
	}

	return packet, remain

}

type Stack []*Packet

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) pop() (*Packet, bool) {
	if s.isEmpty() {
		return nil, false
	}
	elem := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return elem, true
}

func (s *Stack) push(i *Packet) {
	*s = append(*s, i)
}

func sumPackets(ps []*Packet) *Packet {
	var sum = 0
	var length = 0
	for _, p := range ps {
		length += p.Length
		sum += p.Value
	}

	return &Packet{
		Type_:  0,
		Length: length,
		Value:  sum,
	}
}

func prodPackets(ps []*Packet) *Packet {
	var prod = 1
	var length = 0
	for _, p := range ps {
		length += p.Length
		prod *= p.Value
	}

	return &Packet{
		Type_:  1,
		Length: length,
		Value:  prod,
	}
}

func minPackets(ps []*Packet) *Packet {
	var value = ps[0].Value
	var length = 0
	for _, p := range ps {
		length += p.Length
		if p.Value < value {
			value = p.Value
		}
	}

	return &Packet{
		Type_:  2,
		Length: length,
		Value:  value,
	}
}

func maxPackets(ps []*Packet) *Packet {
	var value = ps[0].Value
	var length = 0
	for _, p := range ps {
		length += p.Length
		if p.Value > value {
			value = p.Value
		}
	}

	return &Packet{
		Type_:  3,
		Length: length,
		Value:  value,
	}
}

// eq + 1 3 * 2 2
func count(input []*Packet) int {
	pointer := len(input) - 1
	var stack = &Stack{}
	for pointer >= 0 {
		p := input[pointer]
		// fmt.Printf("%d, %+v\n", pointer, p)
		// stack.print()
		// fmt.Println()
		if p.Type_ == 4 {
			stack.push(p)
			// stack = append(stack, p.Value)
		} else { // operator
			if p.Type_ == 0 {
				if p.LengthTypeId == 1 {
					var nums []*Packet
					var length int
					for length < p.LengthTypeVal {
						n, ok := stack.pop()
						if ok {
							nums = append(nums, n)
							length += n.Length
						}

					}
					sum := sumPackets(nums)
					sum.Length += p.Length
					stack.push(sum)
				}

				if p.LengthTypeId == 2 {
					var nums []*Packet
					for i := 0; i < p.LengthTypeVal; i++ {
						n, ok := stack.pop()
						if ok {
							nums = append(nums, n)
						}
					}
					sum := sumPackets(nums)
					sum.Length += p.Length
					stack.push(sum)
				}
			}

			if p.Type_ == 1 {
				if p.LengthTypeId == 1 {
					var nums []*Packet
					var length int
					for length < p.LengthTypeVal {
						n, ok := stack.pop()
						if ok {
							nums = append(nums, n)
							length += n.Length
						}
					}
					prod := prodPackets(nums)
					prod.Length += p.Length
					stack.push(prod)
				}

				if p.LengthTypeId == 2 {
					var nums []*Packet
					for i := 0; i < p.LengthTypeVal; i++ {
						n, ok := stack.pop()
						if ok {
							nums = append(nums, n)
						}
					}
					prod := prodPackets(nums)
					prod.Length += p.Length
					stack.push(prod)
				}
			}

			if p.Type_ == 2 {
				if p.LengthTypeId == 1 {
					var nums []*Packet
					var length int
					for length < p.LengthTypeVal {
						n, ok := stack.pop()
						if ok {
							nums = append(nums, n)
							length += n.Length
						}

					}
					min := minPackets(nums)
					min.Length += p.Length
					stack.push(min)
				}

				if p.LengthTypeId == 2 {
					var nums []*Packet
					for i := 0; i < p.LengthTypeVal; i++ {
						n, ok := stack.pop()
						if ok {
							nums = append(nums, n)
						}
					}
					min := minPackets(nums)
					min.Length += p.Length
					stack.push(min)
				}
			}

			if p.Type_ == 3 {
				if p.LengthTypeId == 1 {
					var nums []*Packet
					var length int
					for length < p.LengthTypeVal {
						n, ok := stack.pop()
						if ok {
							nums = append(nums, n)
							length += n.Length
						}
					}
					max := maxPackets(nums)
					max.Length += p.Length
					stack.push(max)
				}

				if p.LengthTypeId == 2 {
					var nums []*Packet
					for i := 0; i < p.LengthTypeVal; i++ {
						n, ok := stack.pop()
						if ok {
							nums = append(nums, n)
						}
					}
					max := maxPackets(nums)
					max.Length += p.Length
					stack.push(max)
				}
			}

			if p.Type_ == 5 {
				a, _ := stack.pop()
				b, _ := stack.pop()
				if a.Value > b.Value {
					stack.push(&Packet{
						Length: a.Length + b.Length + p.Length,
						Value:  1,
					})
				} else {
					stack.push(&Packet{
						Length: a.Length + b.Length + p.Length,
						Value:  0,
					})
				}
			}

			if p.Type_ == 6 {
				a, _ := stack.pop()
				b, _ := stack.pop()
				if a.Value < b.Value {
					stack.push(&Packet{
						Length: a.Length + b.Length + p.Length,
						Value:  1,
					})
				} else {
					stack.push(&Packet{
						Length: a.Length + b.Length + p.Length,
						Value:  0,
					})
				}
			}

			if p.Type_ == 7 {
				a, _ := stack.pop()
				b, _ := stack.pop()
				if a.Value == b.Value {
					stack.push(&Packet{
						Length: a.Length + b.Length + p.Length,
						Value:  1,
					})
				} else {
					stack.push(&Packet{
						Length: a.Length + b.Length + p.Length,
						Value:  0,
					})
				}
			}
		}
		pointer--
	}

	p, _ := stack.pop()
	return p.Value
	// stack.print()
}

func (s *Stack) print() {
	for _, p := range *s {
		fmt.Printf("%+v\n", p)
	}
}

func print(packets []*Packet) {
	for _, p := range packets {
		fmt.Printf("%+v\n", p)
	}
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
	var temp = bs
	var sum int
	var cur *Packet
	var packets []*Packet
	for {
		cur, temp = decodePacket(temp)
		sum += cur.Version
		packets = append(packets, cur)
		if len(temp) < 11 {
			break
		}
	}
	fmt.Printf("part1: %d\n", sum)
	fmt.Printf("part2: %d\n", count(packets))

	// _, value := decode(bs, 0)
	// fmt.Printf("part2: %d\n", value)
}
