package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"

	util "github.com/yhsiang/adventofcode"
)

//go:embed example
var example string

//go:embed input
var input string

// learn from https://topaz.github.io/paste/#XQAAAQAqBgAAAAAAAAA4GEiZzRd1JAgz+whYRQxSFI7XvmlfhtGDvX76cUN2pCKaM7BkVT4Yiw81FIfZgj+EGPRB+F5p3n/rfVkRLpibYXSgAMVzbV/QizeBnmyzL7vZRyWhh/ETJpNyY1ITugV6XYmtQpRW/gfmxSw7KGvS1iNMSpEf+TuYY3E5yvESN8kUZMQ9z1T7BV7TalnB58063YoM+NMLpo9b/H4AxlCj6Y/ILj6qcUV0gc+gyqyMnTE8sX5xUTODr2sv/5rzxdUVXL0HtcH87LPblYwsOdgzhgc2EUGdG96l6moq4QhhEnnhdThPMPhYMbJk9fAxBXUKLx8BK6D5wQxsKgQiiMAhP7VQn9f1+tq86TRDwKgRLalu3GDSqpmpCf3Zy0NGFooZMzTTvp+W4j6+nIjpojqS7CrKzMsbbKGV2r4eSLkJ8Sk+DMQKNahd1oojGtj3BKgE/4X3ehXvdlzGpwKyyyhhTQ/L693YAjtWUn90lysilCxX/3ThBapO/hP9OExIihOINShgezr2baF1CURJKDkM3C/y6Aac13Zw8P2IxZdtSJSe24icHqNACIAmuA56TcEcJoXHtkLiqZMtYv1IlIEUwhUlYR+yXfZcQoy1yzPRYIJufOeGYZzGcCs1HPzES1qLVGvdpnKzxPUGDE0osq1kSVQ/53u03grry7pA/zxGpnSOap6Rfy/N6sLkBWUKK+Fqgyi4WbrrWeVK6YYxsLDNxaw9NswJdeaRky6xFTzRfHE9KL3/MePiQaaOk8ET9Bi2VNXWY0ULZ3iJ9Kts95NE0GslLMrXEBMqLJDWyqfLASyVQcMW5Wai4bm8oLKtWvFkzZKjstwUBWBdyEvB00wfKwkM2Ois9j465tIuqBAITcGlDzezsV6jRDRPbIq/A7aPKvqpu1b9CNvBgvue/6BKQmY=

func compare(left, right any) int {
	l, okL := left.(float64)
	r, okR := right.(float64)
	if okL && okR {
		return int(l) - int(r)
	}

	var leftList, rightList []any
	switch left.(type) {
	case []any, []float64:
		leftList = left.([]any)
	case float64:
		leftList = []any{left}
	}
	switch right.(type) {
	case []any, []float64:
		rightList = right.([]any)
	case float64:
		rightList = []any{right}
	}

	for i := range leftList {
		if len(rightList) <= i {
			return 1
		}

		if r := compare(leftList[i], rightList[i]); r != 0 {
			return r
		}
	}

	if len(rightList) == len(leftList) {
		return 0
	}

	return -1
}

func main() {
	var file = example
	if len(os.Args) == 2 && os.Args[1] == "input" {
		file = input
	}

	pairs := strings.Split(file, "\n\n")

	indices := []int{}
	var all []any
	for i, pair := range pairs {
		p := strings.Split(pair, "\n")
		var left, right any
		json.Unmarshal([]byte(p[0]), &left)
		json.Unmarshal([]byte(p[1]), &right)
		all = append(all, left, right)
		if compare(left, right) <= 0 {
			indices = append(indices, i+1)
		}
	}
	// fmt.Printf("%+v\n", indices)
	// first puzzle
	fmt.Printf("%d\n", util.SumInt(indices))

	var one, two any
	json.Unmarshal([]byte("[[2]]"), &one)
	json.Unmarshal([]byte("[[6]]"), &two)
	all = append(all, one, two)
	sort.Slice(all, func(i, j int) bool {
		return compare(all[i], all[j]) < 0
	})

	var r int = 1
	for k, v := range all {
		str, _ := json.Marshal(v)
		if string(str) == "[[2]]" || string(str) == "[[6]]" {
			r *= k + 1
		}
	}
	fmt.Printf("%d\n", r)
}
