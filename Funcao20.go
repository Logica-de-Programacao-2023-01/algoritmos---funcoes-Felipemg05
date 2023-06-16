package main

import (
	"errors"
	"fmt"
)

func filterStrings(slice []string) ([]string, error) {
	if len(slice) == 0 {
		return nil, errors.New("Slice vazio")
	}

	var filteredSlice []string
	for _, str := range slice {
		if len(str) > 5 {
			filteredSlice = append(filteredSlice, str)
		}
	}

	return filteredSlice, nil
}

func main() {
	slice := []string{"apple", "banana", "orange", "grapefruit"}
	filteredSlice, err := filterStrings(slice)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(filteredSlice)
}
