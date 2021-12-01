package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Bag struct {
	Color    string
	Contains map[string]int
}

func (b Bag) Find(str string) string {
	if _, ok := b.Contains[str]; ok {
		return b.Color
	}

	return ""
}

var re = regexp.MustCompile(`(.+) bags$`)

var re2 = regexp.MustCompile(`(\d+) (.+) bag`)

func NewBag(str string) *Bag {
	var contains = make(map[string]int)
	s := strings.Split(str, "contain")
	// fmt.Println(str, s)
	r := re.FindSubmatch([]byte(strings.TrimSpace(s[0])))
	// if strings.TrimSpace(s[1]) == "no other bags." {
	// 	contains["other"] = 0
	// } else {
	cc := strings.Split(strings.TrimSpace(s[1]), ",")
	for _, c := range cc {
		ss := re2.FindSubmatch([]byte(strings.TrimSpace(c)))
		// fmt.Printf("%q", ss)
		if len(ss) != 3 {
			continue
		}
		i, _ := strconv.ParseInt(string(ss[1]), 10, 64)
		contains[string(ss[2])] = int(i)
	}
	// }

	return &Bag{
		Color:    strings.TrimSpace(string(r[1])),
		Contains: contains,
	}
}

func sumBags(n int, key string, all map[string]Bag) int {
	bag, ok := all[key]
	if !ok {
		return n
	}

	var sum int
	for k, v := range bag.Contains {
		s := sumBags(v, k, all)
		sum += s
	}

	return n + n*sum
}

func main() {
	dat, err := ioutil.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")
	var bags []Bag
	for _, line := range lines {
		// fmt.Println(line)
		bags = append(bags, *NewBag(line))
	}
	// fmt.Println(bags)
	// var target = []string{"shiny gold"}
	// var results = make(map[string]struct{})
	// for {
	// 	var tmp []string
	// 	for _, t := range target {
	// 		for _, bag := range bags {
	// 			k := bag.Find(t)
	// 			if k != "" {
	// 				tmp = append(tmp, k)
	// 			}
	// 		}
	// 	}
	// 	target = tmp
	// 	for _, r := range tmp {
	// 		results[r] = struct{}{}
	// 	}
	// 	if len(tmp) == 0 {
	// 		break
	// 	}
	// }
	// PART 1
	// fmt.Println(len(results))

	var target = []string{"shiny gold"}
	var results = make(map[string]Bag)
	for {
		var tmp []Bag
		var tmp2 []string
		for _, t := range target {
			for _, bag := range bags {
				if bag.Color == t && len(bag.Contains) > 0 {
					tmp = append(tmp, bag)
					for k, _ := range bag.Contains {
						tmp2 = append(tmp2, k)
					}
				}
			}
		}

		target = tmp2
		// results = append(results, tmp...)
		for _, t := range tmp {
			results[t.Color] = t
		}
		if len(target) == 0 {
			break
		}
	}

	// fmt.Println(results)
	// fmt.Println(sumBags(2, "wavy silver", results))
	// fmt.Println(sumBags(1, "dark gray", results))
	// fmt.Println(sumBags(1, "dull purple", results))
	// fmt.Println(sumBags(5, "pale salmon", results))
	// fmt.Println(sumBags(2, "plaid blue", results))
	// fmt.Println(sumBags(2, "bright olive", results))
	fmt.Println(sumBags(5, "faded gold", results) + sumBags(5, "pale green", results))
}
