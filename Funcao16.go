package main

import (
	"errors"
	"fmt"
	"strings"
)

func replaceVowels(str string) (string, error) {
	if len(str) == 0 {
		return "", errors.New("String vazia")
	}

	vowels := "aeiouAEIOU"
	var result strings.Builder
	for _, char := range str {
		if strings.ContainsRune(vowels, char) {
			result.WriteRune('*')
		} else {
			result.WriteRune(char)
		}
	}

	return result.String(), nil
}

func main() {
	str := "Salve, rapaziada !"
	replacedStr, err := replaceVowels(str)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(replacedStr)
}
