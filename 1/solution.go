package main

import (
	"fmt"
)

func main() {
	multiples := []int{}
	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			multiples = append(multiples, i)
		}
	}

	sum := 0
	for _, v := range multiples {
		sum = sum + v
	}

	fmt.Println(sum)
}
