package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
)

var re = regexp.MustCompile(`^([\w+\s]+) \(contains ([\w+,\s]+)\)$`)

func exist(name string, ts []string) bool {
	for _, t := range ts {
		if t == name {
			return true
		}
	}
	return false
}

type food struct {
	label       int
	ingredients []string
	allergens   []string
}

type dangerous struct {
	allergen   string
	ingredient string
}

func main() {
	dat, err := ioutil.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")

	// var allergens = make(map[string][]string)
	var foods []food
	for i, line := range lines {
		if s := re.FindSubmatch([]byte(line)); len(s) == 3 {
			var ingredients []string
			var allergens []string

			for _, ss := range strings.Split(string(s[1]), " ") {
				ingredients = append(ingredients, strings.TrimSpace(ss))
			}
			for _, ss := range strings.Split(string(s[2]), ",") {
				allergens = append(allergens, strings.TrimSpace(ss))
			}

			foods = append(foods, food{
				label:       i,
				ingredients: ingredients,
				allergens:   allergens,
			})

		}
	}

	// Calculate frequency
	var allergens = make(map[string]map[string]int)
	for _, food := range foods {
		for _, a := range food.allergens {
			v, ok := allergens[a]
			if !ok {
				v = make(map[string]int)
			}
			for _, i := range food.ingredients {
				if _, ok := v[i]; !ok {
					v[i] = 0
				}
				v[i] += 1
			}
			allergens[a] = v
		}
	}
	fmt.Println(allergens)

	// convert to slice
	var allergenSlice []map[string]int
	for _, v := range allergens {
		allergenSlice = append(allergenSlice, v)
	}

	// sort by frequency
	sort.Slice(allergenSlice, func(i, j int) bool {
		return len(allergenSlice[i]) > len(allergenSlice[j])
	})
	// fmt.Println(allergenSlice)

	// find highest frequency and not in picked
	picked := make(map[string]string)
	for _, s := range allergenSlice {
		var times int
		var allergen string
		for k, v := range s {
			if _, ok := picked[k]; ok {
				continue
			}

			if v > times {
				times = v
				allergen = k
			}
		}
		if _, ok := picked[allergen]; !ok {
			for k, v := range allergens {
				if len(s) == len(v) {
					picked[allergen] = k
				}
			}
		}
	}
	fmt.Println(picked)

	// Find zero allergen of ingredient
	var zeroAllergens []string
	for _, f := range foods {
		for _, i := range f.ingredients {
			if _, ok := picked[i]; !ok && !exist(i, zeroAllergens) {
				zeroAllergens = append(zeroAllergens, i)
			}
		}
	}
	fmt.Println(zeroAllergens)

	// Part1 Count
	var count int
	for _, z := range zeroAllergens {
		for _, f := range foods {
			for _, i := range f.ingredients {
				if i == z {
					count += 1
				}
			}
		}
	}
	fmt.Println(count)

	// Part2
	var dangerousSlice []dangerous
	for k, v := range picked {
		dangerousSlice = append(dangerousSlice, dangerous{
			ingredient: k,
			allergen:   v,
		})
	}
	sort.Slice(dangerousSlice, func(i, j int) bool {
		// wrong here, should be alphabetically
		return dangerousSlice[i].allergen[0] < dangerousSlice[j].allergen[0]
	})
	fmt.Println(dangerousSlice)

	var str string
	for _, d := range dangerousSlice {
		str += d.ingredient + ","
	}
	fmt.Println(str)

}
