package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func unpackString(s string) (string, error) {
	// Check if string is empty
	if s == "" {
		return "", errors.New("empty string")
	}

	var res strings.Builder
	// Convert string to slice of runes
	runes := []rune(s)

	// Iterate over runes
	for i := 0; i < len(runes); i++ {
		if unicode.IsDigit(runes[i]) {
			// If first character is digit return error
			if i == 0 {
				return "", errors.New("invalid string")
			}

			// Convert digit to integer
			count, err := strconv.Atoi(string(runes[i]))
			if err != nil {
				return "", err
			}

			// Write previous character count times
			for j := 0; j < count; j++ {
				res.WriteRune(runes[i-1])
			}
			continue
		}
		// If not digit just append to result
		res.WriteRune(runes[i])
	}
	return res.String(), nil
}

func main() {
	fmt.Println(unpackString("a4bc2d5e"))
}
