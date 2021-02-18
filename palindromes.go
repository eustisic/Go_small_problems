package main

import (
	"fmt"
	"unicode"
)

// loop through the string with two indexes

func Valid(r rune) bool {
	unicode.IsLetter(r) || unicode.IsNumber(r)
}

func palindromeString(s string) bool {
	i := 0
	j := len(s) - 1

	for i < j {
		for i < j && !Valid(rune(s[i])) { i++ }
		for i < j && !Valid(rune(s[j])) { j-- }

		if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])) {
			return false
		}

		i++
		j--
	}

	return true
}



func main() {
	fmt.Println(palindromeString("A man, a plan, a canal, panama"))
	fmt.Println(palindromeString("race a car"))
	fmt.Println(palindromeString(""))
}