package main

import "fmt"

const TargetNumber = 100

func main() {
	fmt.Println(squareOfSum(TargetNumber) - sumOfSquares(TargetNumber))
}

func sumOfSquares(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum = sum + i*i
	}

	return sum
}

func squareOfSum(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum = sum + i
	}

	return sum * sum
}
