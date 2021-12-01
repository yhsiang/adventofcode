package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Map struct {
	values [][]string
	RowNum int
	ColNum int
}

func NewMap(lines []string) *Map {
	var values [][]string
	for _, line := range lines {
		values = append(values, strings.Split(line, ""))
	}

	return &Map{
		values: values,
		RowNum: len(values),
		ColNum: len(values[0]),
	}
}

func (m *Map) Set(i, j int, value string) {
	m.values[i][j] = value
}

func (m *Map) Search(y, x int, direction string) string {
	var horizon, vertical int
	switch direction {
	case "e":
		horizon = 1
		for {
			x += horizon
			if x > m.ColNum-1 {
				break
			}
			if m.values[y][x] != "." {
				return m.values[y][x]
			}

		}
	case "w":
		horizon = -1
		for {
			x += horizon
			if x < 0 {
				break
			}
			if m.values[y][x] != "." {
				return m.values[y][x]
			}

		}
	case "s":
		vertical = 1
		for {
			y += vertical
			if y > m.RowNum-1 {
				break
			}
			if m.values[y][x] != "." {
				return m.values[y][x]
			}
		}
	case "n":
		vertical = -1
		for {
			y += vertical
			if y < 0 {
				break
			}
			if m.values[y][x] != "." {
				return m.values[y][x]
			}

		}
	case "nw", "wn":
		horizon = -1
		vertical = -1

		for {
			x += horizon
			y += vertical
			if y < 0 || x < 0 {
				break
			}
			if m.values[y][x] != "." {
				return m.values[y][x]
			}

		}
	case "ne", "en":
		horizon = 1
		vertical = -1
		for {
			x += horizon
			y += vertical
			if y < 0 || x > m.ColNum-1 {
				break
			}
			if m.values[y][x] != "." {
				return m.values[y][x]
			}

		}
	case "sw", "ws":
		horizon = -1
		vertical = 1
		for {
			x += horizon
			y += vertical
			if y > m.RowNum-1 || x < 0 {
				break
			}

			if m.values[y][x] != "." {
				return m.values[y][x]
			}

		}
	case "se", "es":
		horizon = 1
		vertical = 1

		for {
			x += horizon
			y += vertical
			if y > m.RowNum-1 || x > m.ColNum-1 {
				break
			}
			if m.values[y][x] != "." {
				return m.values[y][x]
			}

		}
	}

	return "."
}

func (m *Map) isEmpty2(str string) bool {
	if str == "L" || str == "." {
		return true
	}
	return false
}

func (m *Map) isOccupied2(str string) int {
	if str == "#" {
		return 1
	}
	return 0
}

func (m *Map) isEmpty(i, j int) bool {
	if m.values[i][j] == "L" || m.values[i][j] == "." {
		return true
	}
	return false
}

func (m *Map) isOccupied(i, j int) int {
	if m.values[i][j] == "#" {
		return 1
	}
	return 0
}

func (m Map) Print() {
	for _, row := range m.values {
		fmt.Println(row)
	}
}

func (m Map) CountSeats(str string) int {
	var n int
	for _, mm := range m.values {
		for _, i := range mm {
			if i == str {
				n++
			}
		}
	}
	return n
}

func copyMap(m Map) Map {
	n := Map{}
	n.values = make([][]string, len(m.values))
	for i, mm := range m.values {
		n.values[i] = make([]string, len(mm))
		copy(n.values[i], mm)
	}
	n.RowNum = len(n.values)
	n.ColNum = len(n.values[0])
	return n
}

func main() {
	dat, err := ioutil.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")

	m := NewMap(lines)

	// fmt.Println(m.Search(1, 1, "nw"))
	// fmt.Println(m.Search(1, 1, "nw"))

	// fmt.Println(m.Search(4, 3, "sw"))

	// Part 1
	// newMap := copyMap(*m)

	// for {
	// 	var changed int
	// 	for i, row := range m.values {
	// 		for j, col := range row {
	// 			switch col {
	// 			case "L":
	// 				if i == 0 && j == 0 {
	// 					if m.isEmpty(i+1, j) &&
	// 						m.isEmpty(i, j+1) &&
	// 						m.isEmpty(i+1, j+1) {
	// 						newMap.Set(i, j, "#")
	// 						changed++
	// 					}
	// 					continue
	// 				}
	// 				if i == 0 && j == len(row)-1 {
	// 					if m.isEmpty(i, j-1) &&
	// 						m.isEmpty(i+1, j) &&
	// 						m.isEmpty(i+1, j-1) {
	// 						newMap.Set(i, j, "#")
	// 						changed++
	// 					}
	// 					continue
	// 				}
	// 				if i == len(m.values)-1 && j == 0 {
	// 					if m.isEmpty(i-1, j) &&
	// 						m.isEmpty(i, j+1) &&
	// 						m.isEmpty(i-1, j+1) {
	// 						newMap.Set(i, j, "#")
	// 						changed++
	// 					}
	// 					continue
	// 				}
	// 				if i == len(m.values)-1 && j == len(row)-1 {
	// 					if m.isEmpty(i-1, j) &&
	// 						m.isEmpty(i, j-1) &&
	// 						m.isEmpty(i-1, j-1) {
	// 						newMap.Set(i, j, "#")
	// 						changed++
	// 					}
	// 					continue
	// 				}

	// 				if i == 0 {
	// 					if m.isEmpty(i, j-1) &&
	// 						m.isEmpty(i+1, j) &&
	// 						m.isEmpty(i, j+1) &&
	// 						m.isEmpty(i+1, j-1) &&
	// 						m.isEmpty(i+1, j+1) {
	// 						newMap.Set(i, j, "#")
	// 						changed++
	// 					}
	// 					continue
	// 				}

	// 				if i == len(m.values)-1 {
	// 					if m.isEmpty(i-1, j) &&
	// 						m.isEmpty(i, j-1) &&
	// 						m.isEmpty(i, j+1) &&
	// 						m.isEmpty(i-1, j-1) &&
	// 						m.isEmpty(i-1, j+1) {
	// 						newMap.Set(i, j, "#")
	// 						changed++
	// 					}
	// 					continue
	// 				}

	// 				if j == 0 {
	// 					if m.isEmpty(i-1, j) &&
	// 						m.isEmpty(i+1, j) &&
	// 						m.isEmpty(i, j+1) &&
	// 						m.isEmpty(i-1, j+1) &&
	// 						m.isEmpty(i+1, j+1) {
	// 						newMap.Set(i, j, "#")
	// 						changed++
	// 					}
	// 					continue
	// 				}

	// 				if j == len(row)-1 {
	// 					if m.isEmpty(i-1, j) &&
	// 						m.isEmpty(i, j-1) &&
	// 						m.isEmpty(i+1, j) &&
	// 						m.isEmpty(i-1, j-1) &&
	// 						m.isEmpty(i+1, j-1) {
	// 						newMap.Set(i, j, "#")
	// 						changed++
	// 					}
	// 					continue
	// 				}

	// 				if m.isEmpty(i-1, j) &&
	// 					m.isEmpty(i, j-1) &&
	// 					m.isEmpty(i+1, j) &&
	// 					m.isEmpty(i, j+1) &&
	// 					m.isEmpty(i-1, j-1) &&
	// 					m.isEmpty(i-1, j+1) &&
	// 					m.isEmpty(i+1, j-1) &&
	// 					m.isEmpty(i+1, j+1) {
	// 					newMap.Set(i, j, "#")
	// 					changed++
	// 				}

	// 			case "#":
	// 				var n = 0
	// 				if i == 0 && j == 0 {
	// 					continue
	// 				}
	// 				if i == 0 && j == len(row)-1 {
	// 					continue
	// 				}
	// 				if i == len(m.values)-1 && j == 0 {
	// 					continue
	// 				}
	// 				if i == len(m.values)-1 && j == len(row)-1 {
	// 					continue
	// 				}
	// 				if i == 0 {
	// 					n = m.isOccupied(i, j-1) +
	// 						m.isOccupied(i+1, j) +
	// 						m.isOccupied(i, j+1) +
	// 						m.isOccupied(i+1, j-1) +
	// 						m.isOccupied(i+1, j+1)
	// 					if n >= 4 {
	// 						newMap.Set(i, j, "L")
	// 						changed++
	// 					}
	// 					continue
	// 				}

	// 				if i == len(m.values)-1 {
	// 					n = m.isOccupied(i-1, j) +
	// 						m.isOccupied(i, j-1) +
	// 						m.isOccupied(i, j+1) +
	// 						m.isOccupied(i-1, j-1) +
	// 						m.isOccupied(i-1, j+1)
	// 					if n >= 4 {
	// 						newMap.Set(i, j, "L")
	// 						changed++
	// 					}
	// 					continue
	// 				}

	// 				if j == 0 {
	// 					n = m.isOccupied(i-1, j) +
	// 						m.isOccupied(i+1, j) +
	// 						m.isOccupied(i, j+1) +
	// 						m.isOccupied(i-1, j+1) +
	// 						m.isOccupied(i+1, j+1)
	// 					if n >= 4 {
	// 						newMap.Set(i, j, "L")
	// 						changed++
	// 					}
	// 					continue
	// 				}

	// 				if j == len(row)-1 {
	// 					n = m.isOccupied(i-1, j) +
	// 						m.isOccupied(i, j-1) +
	// 						m.isOccupied(i+1, j) +
	// 						m.isOccupied(i-1, j-1) +
	// 						m.isOccupied(i+1, j-1)
	// 					if n >= 4 {
	// 						newMap.Set(i, j, "L")
	// 						changed++
	// 					}
	// 					continue
	// 				}

	// 				n = m.isOccupied(i-1, j) +
	// 					m.isOccupied(i, j-1) +
	// 					m.isOccupied(i+1, j) +
	// 					m.isOccupied(i, j+1) +
	// 					m.isOccupied(i-1, j-1) +
	// 					m.isOccupied(i-1, j+1) +
	// 					m.isOccupied(i+1, j-1) +
	// 					m.isOccupied(i+1, j+1)
	// 				if n >= 4 {
	// 					newMap.Set(i, j, "L")
	// 					changed++
	// 				}
	// 			case ".":
	// 				continue
	// 			}
	// 		}
	// 	}
	// 	// newMap.Print()
	// 	// fmt.Println()
	// 	tmp := copyMap(newMap)
	// 	m = &tmp
	// 	if changed == 0 {
	// 		break
	// 	}
	// }
	// fmt.Println(m.CountSeats("#"))

	newMap := copyMap(*m)
	for {
		var changed int
		for i, row := range m.values {
			for j, col := range row {
				switch col {
				case "L":
					if i == 0 && j == 0 {
						if m.isEmpty2(m.Search(i, j, "e")) &&
							m.isEmpty2(m.Search(i, j, "se")) &&
							m.isEmpty2(m.Search(i, j, "s")) {
							newMap.Set(i, j, "#")
							changed++
						}
						continue
					}
					if i == 0 && j == len(row)-1 {
						if m.isEmpty2(m.Search(i, j, "w")) &&
							m.isEmpty2(m.Search(i, j, "sw")) &&
							m.isEmpty2(m.Search(i, j, "s")) {
							newMap.Set(i, j, "#")
							changed++
						}
						continue
					}
					if i == len(m.values)-1 && j == 0 {
						if m.isEmpty2(m.Search(i, j, "e")) &&
							m.isEmpty2(m.Search(i, j, "ne")) &&
							m.isEmpty2(m.Search(i, j, "n")) {
							newMap.Set(i, j, "#")
							changed++
						}
						continue
					}
					if i == len(m.values)-1 && j == len(row)-1 {
						if m.isEmpty2(m.Search(i, j, "w")) &&
							m.isEmpty2(m.Search(i, j, "nw")) &&
							m.isEmpty2(m.Search(i, j, "n")) {
							newMap.Set(i, j, "#")
							changed++
						}
						continue
					}

					if i == 0 {
						if m.isEmpty2(m.Search(i, j, "w")) &&
							m.isEmpty2(m.Search(i, j, "sw")) &&
							m.isEmpty2(m.Search(i, j, "s")) &&
							m.isEmpty2(m.Search(i, j, "se")) &&
							m.isEmpty2(m.Search(i, j, "e")) {
							newMap.Set(i, j, "#")
							changed++
						}
						continue
					}

					if i == len(m.values)-1 {
						if m.isEmpty2(m.Search(i, j, "w")) &&
							m.isEmpty2(m.Search(i, j, "nw")) &&
							m.isEmpty2(m.Search(i, j, "n")) &&
							m.isEmpty2(m.Search(i, j, "ne")) &&
							m.isEmpty2(m.Search(i, j, "e")) {
							newMap.Set(i, j, "#")
							changed++
						}
						continue
					}

					if j == 0 {
						if m.isEmpty2(m.Search(i, j, "n")) &&
							m.isEmpty2(m.Search(i, j, "ne")) &&
							m.isEmpty2(m.Search(i, j, "e")) &&
							m.isEmpty2(m.Search(i, j, "se")) &&
							m.isEmpty2(m.Search(i, j, "s")) {
							newMap.Set(i, j, "#")
							changed++
						}
						continue
					}

					if j == len(row)-1 {
						if m.isEmpty2(m.Search(i, j, "n")) &&
							m.isEmpty2(m.Search(i, j, "nw")) &&
							m.isEmpty2(m.Search(i, j, "w")) &&
							m.isEmpty2(m.Search(i, j, "sw")) &&
							m.isEmpty2(m.Search(i, j, "s")) {
							newMap.Set(i, j, "#")
							changed++
						}
						continue
					}

					if m.isEmpty2(m.Search(i, j, "w")) &&
						m.isEmpty2(m.Search(i, j, "nw")) &&
						m.isEmpty2(m.Search(i, j, "n")) &&
						m.isEmpty2(m.Search(i, j, "ne")) &&
						m.isEmpty2(m.Search(i, j, "e")) &&
						m.isEmpty2(m.Search(i, j, "se")) &&
						m.isEmpty2(m.Search(i, j, "s")) &&
						m.isEmpty2(m.Search(i, j, "sw")) {
						newMap.Set(i, j, "#")
						changed++
					}

				case "#":
					var n = 0
					if i == 0 && j == 0 {
						continue
					}
					if i == 0 && j == len(row)-1 {
						continue
					}
					if i == len(m.values)-1 && j == 0 {
						continue
					}
					if i == len(m.values)-1 && j == len(row)-1 {
						continue
					}
					if i == 0 {
						n = m.isOccupied2(m.Search(i, j, "w")) +
							m.isOccupied2(m.Search(i, j, "sw")) +
							m.isOccupied2(m.Search(i, j, "s")) +
							m.isOccupied2(m.Search(i, j, "se")) +
							m.isOccupied2(m.Search(i, j, "e"))
						if n >= 5 {
							newMap.Set(i, j, "L")
							changed++
						}
						continue
					}

					if i == len(m.values)-1 {
						n = m.isOccupied2(m.Search(i, j, "w")) +
							m.isOccupied2(m.Search(i, j, "nw")) +
							m.isOccupied2(m.Search(i, j, "n")) +
							m.isOccupied2(m.Search(i, j, "ne")) +
							m.isOccupied2(m.Search(i, j, "e"))
						if n >= 5 {
							newMap.Set(i, j, "L")
							changed++
						}
						continue
					}

					if j == 0 {
						n = m.isOccupied2(m.Search(i, j, "n")) +
							m.isOccupied2(m.Search(i, j, "ne")) +
							m.isOccupied2(m.Search(i, j, "e")) +
							m.isOccupied2(m.Search(i, j, "se")) +
							m.isOccupied2(m.Search(i, j, "s"))
						if n >= 5 {
							newMap.Set(i, j, "L")
							changed++
						}
						continue
					}

					if j == len(row)-1 {
						n = m.isOccupied2(m.Search(i, j, "n")) +
							m.isOccupied2(m.Search(i, j, "nw")) +
							m.isOccupied2(m.Search(i, j, "w")) +
							m.isOccupied2(m.Search(i, j, "sw")) +
							m.isOccupied2(m.Search(i, j, "s"))
						// if i == 1 {
						// 	fmt.Println(i, j, m.Search(i, j, "n"))
						// 	fmt.Println(i, j, m.Search(i, j, "nw"))
						// 	fmt.Println(i, j, m.Search(i, j, "w"))
						// 	fmt.Println(i, j, m.Search(i, j, "sw"))
						// 	fmt.Println(i, j, m.Search(i, j, "s"))
						// 	fmt.Println(n)
						// }

						if n >= 5 {
							newMap.Set(i, j, "L")
							changed++
						}
						continue
					}

					n = m.isOccupied2(m.Search(i, j, "w")) +
						m.isOccupied2(m.Search(i, j, "nw")) +
						m.isOccupied2(m.Search(i, j, "n")) +
						m.isOccupied2(m.Search(i, j, "ne")) +
						m.isOccupied2(m.Search(i, j, "e")) +
						m.isOccupied2(m.Search(i, j, "se")) +
						m.isOccupied2(m.Search(i, j, "s")) +
						m.isOccupied2(m.Search(i, j, "sw"))
					if n >= 5 {
						newMap.Set(i, j, "L")
						changed++
					}
				case ".":
					continue
				}
			}
		}
		// newMap.Print()
		// fmt.Println()
		tmp := copyMap(newMap)
		m = &tmp
		if changed == 0 {
			break
		}
	}
	fmt.Println(m.CountSeats("#"))
}
