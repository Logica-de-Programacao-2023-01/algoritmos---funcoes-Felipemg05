package main

import (
	"errors"
	"fmt"
)

func applyFunction(slice []int, fn func(int) int) (int, error) {
	if len(slice) == 0 {
		return 0, errors.New("Slice vazio")
	}

	sum := 0
	for _, num := range slice {
		result := fn(num)
		sum += result
	}

	return sum, nil
}

func double(num int) int {
	return num * 2
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	result, err := applyFunction(slice, double)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
