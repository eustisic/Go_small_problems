package main

import (
	"fmt"
	"strings"
)

func ReverseWordsInString(s string) string {
	arr := strings.Fields(s)
	l := len(arr) - 1

	for i := 0; i < len(arr) / 2; i++ {
		arr[i], arr[l - i] = arr[l - i], arr[i]
	}

	return strings.Join(arr, " ")
}


func main() {
  fmt.Println(ReverseWordsInString("the sky is blue"))
  fmt.Println(ReverseWordsInString("  hello world  "))
}