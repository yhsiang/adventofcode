package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func count(str string) int {
	var questions = make(map[string]int)
	people := strings.Split(str, "\n")
	for _, person := range people {
		ans := strings.Split(person, "")
		for _, b := range ans {
			if _, ok := questions[b]; !ok {
				questions[b] = 0
			}
			questions[b] += 1
		}
	}

	if len(people) == 1 {
		return len(questions)
	}

	var n = 0
	for _, q := range questions {
		if q == len(people) {
			n += 1
		}
	}

	return n
}

func main() {
	dat, err := ioutil.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	answers := strings.Split(string(dat), "\n\n")

	var i int
	for _, answer := range answers {
		a := count(answer)
		// fmt.Println(a)
		i += a
	}
	fmt.Println(i)
}
