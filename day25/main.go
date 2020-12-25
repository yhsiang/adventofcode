package main

import "fmt"

func loopSize(n int) int {
	i := 0
	value := 1
	subject := 7
	for {
		multi := value * subject
		remainder := multi % 20201227
		if value == n {
			return i
		}
		i += 1
		value = remainder
	}
}

func encryption(key, loopSize int) int {
	value := 1
	subject := key
	for i := 0; i < loopSize; i++ {
		multi := value * subject
		remainder := multi % 20201227
		value = remainder
	}
	return value
}

func main() {
	// cardKey, doorKey := 5764801, 17807724
	// cardSize, doorSize := loopSize(cardKey), loopSize(doorKey)
	// fmt.Println(cardSize, doorSize)
	// fmt.Println(encryption(doorKey, cardSize), encryption(cardKey, doorSize))

	cardKey, doorKey := 5290733, 15231938
	cardSize, doorSize := loopSize(cardKey), loopSize(doorKey)
	fmt.Println(cardSize, doorSize)
	fmt.Println(encryption(doorKey, cardSize), encryption(cardKey, doorSize))
}
