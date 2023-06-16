package main

import (
	"errors"
	"fmt"
	"math"
)

func isPrime(num int) bool {
	if num < 2 {
		return false
	}

	sqrt := int(math.Sqrt(float64(num)))
	for i := 2; i <= sqrt; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func getPrimes(n int) ([]int, error) {
	if n < 0 {
		return nil, errors.New("Número negativo")
	}

	var primes []int
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}

	return primes, nil
}

func main() {
	number := 20
	primeSlice, err := getPrimes(number)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(primeSlice)
}
