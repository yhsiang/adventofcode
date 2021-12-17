package main

import (
	"testing"
)

// go test -bench=. ./2021/day16
// goos: darwin
// goarch: arm64
// pkg: github.com/yhsiang/adventofcode/2021/day16
// BenchmarkPrefixEval-8   	  166462	      7198 ns/op
// BenchmarkRecursion-8    	   36925	     32441 ns/op
// PASS
// ok  	github.com/yhsiang/adventofcode/2021/day16	4.007s
func BenchmarkPrefixEval(b *testing.B) {
	bs := convert(input)
	var temp = bs
	var cur *Packet
	var packets []*Packet
	for {
		cur, temp = decodePacket(temp)
		packets = append(packets, cur)
		if len(temp) < 11 {
			break
		}
	}
	for n := 0; n < b.N; n++ {
		prefixEval(packets)
	}
}

func BenchmarkRecursion(b *testing.B) {
	bs := convert(input)
	for n := 0; n < b.N; n++ {
		decode(bs, 0)
	}
}
