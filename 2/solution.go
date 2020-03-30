package main

import (
	"fmt"
)

const UpperLimit = 4000000

func main() {
	n := 0
	evens := []int{}
	for {
		result := fibo(n)
		if result > UpperLimit {
			break
		}

		if result%2 == 0 {
			evens = append(evens, result)
		}
		n++
	}

	var sum int
	for _, v := range evens {
		sum = sum + v
	}

	fmt.Println(sum)
}

func fibo(n int) int {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return 2
	}

	return fibo(n-1) + fibo(n-2)
}
