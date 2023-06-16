package main

import (
	"errors"
	"fmt"
)

func containsNumber(number int, slice []int) (bool, error) {
	if len(slice) == 0 {
		return false, errors.New("Slice vazio")
	}

	for _, n := range slice {
		if n == number {
			return true, nil
		}
	}

	return false, nil
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	number := 3

	contains, err := containsNumber(number, slice)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(contains)
}
