package util

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func ByLine(input string) []string {
	data := strings.Split(string(input), "\n")
	return data
}

func ToInt64(data []string) []int64 {
	var nums []int64
	for _, d := range data {
		i, err := strconv.ParseInt(d, 10, 64)
		if err != nil {
			panic(err)
		}
		nums = append(nums, i)
	}
	return nums
}

func ToInt(data []string) []int {
	var nums []int
	for _, d := range data {
		i, err := strconv.ParseInt(d, 10, 64)
		if err != nil {
			panic(err)
		}
		nums = append(nums, int(i))
	}
	return nums
}

func Int64(data string) (num int64, err error) {
	num, err = strconv.ParseInt(data, 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func Int(data string) (int, error) {
	num, err := strconv.ParseInt(data, 10, 64)
	return int(num), err
}

func BinToInt(data string) (int, error) {
	num, err := strconv.ParseInt(data, 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	return int(num), err
}

func BinToInt64(data string) (num int64, err error) {
	num, err = strconv.ParseInt(data, 2, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

// TODO: generic
func SumInt64(data []int64) int64 {
	var sum int64 = 0
	for _, d := range data {
		sum += d
	}
	return sum
}

func SumInt(data []int) int {
	var sum int = 0
	for _, d := range data {
		sum += d
	}
	return sum
}

func MultiplyInt(data []int) int {
	var sum int = 1
	for _, d := range data {
		sum *= d
	}
	return sum
}

func Abs(x int) int {
	return int(math.Abs(float64(x)))
}

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func Coord(input string) (x int, y int) {
	strs := strings.Split(input, ",")
	a, _ := Int64(strs[0])
	b, _ := Int64(strs[1])
	x = int(a)
	y = int(b)
	return
}

func Coord3d(input string) (x int, y int, z int) {
	strs := strings.Split(input, ",")
	a, _ := Int64(strs[0])
	b, _ := Int64(strs[1])
	c, _ := Int64(strs[2])
	x = int(a)
	y = int(b)
	z = int(c)
	return
}

func Coord3dSum(a string, b string) string {
	x1, y1, z1 := Coord3d(a)
	x2, y2, z2 := Coord3d(b)

	return fmt.Sprintf("%d,%d,%d", x1+x2, y1+y2, z1+z2)
}

func Coord3dSub(a string, b string) string {
	x1, y1, z1 := Coord3d(a)
	x2, y2, z2 := Coord3d(b)

	return fmt.Sprintf("%d,%d,%d", x1-x2, y1-y2, z1-z2)
}

func ManhattanDist(a, b string) int {
	x1, y1, z1 := Coord3d(a)
	x2, y2, z2 := Coord3d(b)
	return Abs(x1-x2) + Abs(y1-y2) + Abs(z1-z2)
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}

func ReverseStringSlice(strs []string) []string {
	var newStrs = make([]string, len(strs))
	copy(newStrs, strs)
	for i, j := 0, len(newStrs)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		newStrs[i], newStrs[j] = newStrs[j], newStrs[i]
	}

	return newStrs
}

func GCD(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int64, integers ...int64) int64 {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
