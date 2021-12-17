package main

import (
	_ "embed"
	"fmt"
	"os"
	"regexp"
	"strings"
)

//go:embed example
var example string

//go:embed input
var input string

var (
	zero = regexp.MustCompile(`0`)
	one  = regexp.MustCompile(`1`)
	two  = regexp.MustCompile(`2`)
)

type Image struct {
	Layers []string
	Width  int
	Height int
}

func newImage(data string, width, height int) *Image {
	var layers []string
	var size = width * height
	var layerNum = len(data) / size
	var s int = 0
	var e int = size
	for i := 1; i <= layerNum; i++ {
		// layers[i] = data[s:e]
		layers = append(layers, data[s:e])
		s += size
		e += size

	}

	return &Image{
		Layers: layers,
		Width:  width,
		Height: height,
	}
}

func (i *Image) zeroLen() int {
	var layers = make(map[int]int)
	for k, v := range i.Layers {
		s := zero.FindAll([]byte(v), -1)
		layers[k] = len(s)
	}

	var j = 1
	var less = layers[j]
	delete(layers, j)
	for k, v := range layers {
		if v < less {
			less = v
			j = k
		}
	}

	return j
}

func (i *Image) sumOneTwo(layer int) int {
	value := i.Layers[layer]
	s := one.FindAll([]byte(value), -1)
	b := two.FindAll([]byte(value), -1)
	return len(s) * len(b)
}

func (i *Image) decode() string {
	var message []string
	var size = i.Width * i.Height
	// fmt.Println(i.Layers)
	for j := 0; j < size; j++ {
		var pixel string
		for _, v := range i.Layers {
			if string(v[j]) == "2" {
				continue
			}
			if pixel == "" {
				pixel = string(v[j])
			}
		}
		message = append(message, strings.Replace(pixel, "0", " ", 1))
	}

	var a int
	for k := 0; k < i.Height; k++ {
		for j := 0; j < i.Width; j++ {
			fmt.Printf("%s", message[a])
			a++
		}
		fmt.Println()
	}
	return strings.Join(message, "")
}

func main() {
	var file = example
	var w = 3
	var h = 2
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
		w = 25
		h = 6
	}

	l := newImage(file, w, h)
	fmt.Printf("part1: %d\n", l.sumOneTwo(l.zeroLen()))
	fmt.Printf("part2:\n")
	l.decode()
}
