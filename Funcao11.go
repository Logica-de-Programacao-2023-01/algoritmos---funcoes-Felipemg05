package main

import (
	"errors"
	"fmt"
	"math"
)

func isPrime(n int) (bool, error) {
	if n < 0 {
		return false, errors.New("Número negativo")
	}

	if n <= 1 {
		return false, nil
	}

	if n == 2 {
		return true, nil
	}

	if n%2 == 0 {
		return false, nil
	}

	sqrt := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrt; i += 2 {
		if n%i == 0 {
			return false, nil
		}
	}

	return true, nil
}

func main() {
	number := 17
	isPrime, err := isPrime(number)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(isPrime)
}
