package main

import (
	"errors"
	"fmt"
	"strconv"
)

func sumDigits(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("Número negativo")
	}

	str := strconv.Itoa(n)

	sum := 0
	for _, digit := range str {
		digitInt, _ := strconv.Atoi(string(digit))
		sum += digitInt
	}

	return sum, nil
}

func main() {
	number := 12345
	sum, err := sumDigits(number)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(sum)
}
