package main

import "fmt"

const TargetNumber = 600851475143

var primes map[int]interface{}

func main() {
	factorization := getFactorization(TargetNumber)

	biggestPrime := 0
	for k := range factorization {
		if k > biggestPrime {
			biggestPrime = k
		}
	}

	fmt.Println(biggestPrime)
}

func getFactorization(n int) map[int]int {
	primesInFactorization := make(map[int]int)

	i := 2
	for {
		if isPrime(i) {
			quotient, divisor := factorization(n)
			if divisor != 0 {
				if v, ok := primesInFactorization[divisor]; ok {
					primesInFactorization[divisor] = v + 1
				} else {
					primesInFactorization[divisor] = 1
				}
			}
			if quotient == 1 {
				break
			}
			n = quotient
		}
		i++
	}

	return primesInFactorization
}

func factorization(n int) (int, int) {
	for p := range primes {
		if n%p == 0 {
			return n / p, p
		}
	}

	return n, 0
}

func isPrime(n int) bool {
	for v := range primes {
		if n%v == 0 {
			return false
		}
	}

	if primes == nil {
		primes = make(map[int]interface{})
	}

	primes[n] = struct{}{}
	return true
}
