package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

func concatenateStrings(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("Slice vazio")
	}

	sort.Strings(slice)
	result := strings.Join(slice, "")

	return result, nil
}

func main() {
	slice := []string{"banana", "abacaxi", "laranja", "uva"}
	concatenatedStr, err := concatenateStrings(slice)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(concatenatedStr)
}
