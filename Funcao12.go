package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func filterUpperCase(strings []string) (string, error) {
	if len(strings) == 0 {
		return "", errors.New("Slice vazio")
	}

	var result strings.Builder
	for _, str := range strings {
		if unicode.IsUpper(rune(str[0])) {
			result.WriteString(str)
			result.WriteString(" ")
		}
	}

	return result.String(), nil
}

func main() {
	strings := []string{"OpenAI", "Hello", "world", "ChatGPT", "Go"}
	filteredStr, err := filterUpperCase(strings)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(filteredStr)
}
