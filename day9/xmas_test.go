package main

import (
	"fmt"
	"testing"
)

func TestXmasAdd(t *testing.T) {
	x := NewXmas([]int64{35, 20, 15, 25, 47})

	// x.Add()

	fmt.Println(x.Sums)
}
