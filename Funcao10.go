package main

import (
	"errors"
	"fmt"
	"sort"
)

func ordenarSlice(slice []int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("O slice está vazio")
	}

	novoSlice := make([]int, len(slice))
	copy(novoSlice, slice)

	sort.Ints(novoSlice)

	return novoSlice, nil
}

func main() {
	slice := []int{4, 2, 6, 1, 3, 5}

	novoSlice, err := ordenarSlice(slice)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println("Slice original:", slice)
	fmt.Println("Slice ordenado:", novoSlice)
}
