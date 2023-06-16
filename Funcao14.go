package main

import (
	"errors"
	"fmt"
)

func intersectSlices(slice1, slice2 []int) ([]int, error) {
	if len(slice1) == 0 || len(slice2) == 0 {
		return nil, errors.New("Pelo menos um dos slices está vazio")
	}

	elements := make(map[int]bool)
	for _, num := range slice1 {
		elements[num] = true
	}

	var intersect []int
	for _, num := range slice2 {
		if elements[num] {
			intersect = append(intersect, num)
		}
	}

	return intersect, nil
}

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{4, 5, 6, 7, 8}
	intersection, err := intersectSlices(slice1, slice2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(intersection)
}
