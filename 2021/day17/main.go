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

type Probe struct {
	PX       int
	PY       int
	VX       int
	VY       int
	HighestY int
	// target
	XRange []int
	YRange []int
}

func initProbe(xrange, yrange string) *Probe {
	xs := strings.Split(xrange, "..")
	ys := strings.Split(yrange, "..")
	fromX, _ := util.Int(xs[0])
	toX, _ := util.Int(xs[1])
	fromY, _ := util.Int(ys[0])
	toY, _ := util.Int(ys[1])

	return &Probe{
		XRange: []int{
			fromX,
			toX,
		},
		YRange: []int{
			toY,
			fromY,
		},
	}
}

func (p *Probe) move() {
	p.PX += p.VX
	p.PY += p.VY

	if p.VX > 0 {
		p.VX -= 1
	} else if p.VX < 0 {
		p.VX += 1
	}

	if p.PY > p.HighestY {
		p.HighestY = p.PY
	}

	p.VY -= 1
}

func (p *Probe) isTarget() bool {
	return p.PX >= p.XRange[0] &&
		p.PX <= p.XRange[1] &&
		p.PY <= p.YRange[0] &&
		p.PY >= p.YRange[1]
}

func (p *Probe) isOver() bool {
	if p.PX > p.XRange[1] {
		return true
	}

	if p.PY < p.YRange[1] {
		return true
	}

	return false
}

func (p *Probe) seekTrajectory() (int, []string) {
	i := 0
	var y int
	var velocities []string
	// 0 .. max x
	for i <= p.XRange[1] {
		k := p.YRange[1]
		// max y .. -1*max y
		for k <= -1*p.YRange[1] {
			p.PX = 0
			p.PY = 0
			p.VX = i
			p.VY = k
			p.HighestY = 0
			startVx := i
			startVy := k
			for {
				p.move()
				if p.isTarget() {
					break
				}
				if p.isOver() {
					break
				}
			}
			if p.isTarget() {
				if p.HighestY > y {
					y = p.HighestY
				}
				velocities = append(velocities, fmt.Sprintf("%d,%d", startVx, startVy))
			}
			k++
		}
		i++
	}

	return y, velocities
}

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}

	area := strings.Split(file, "target area: ")
	positions := strings.Split(area[1], ", ")
	xrange := strings.Split(positions[0], "=")
	yrange := strings.Split(positions[1], "=")
	p := initProbe(xrange[1], yrange[1])
	y, vs := p.seekTrajectory()
	fmt.Printf("part1: %d\n", y)
	fmt.Printf("part2: %d\n", len(vs))

}
