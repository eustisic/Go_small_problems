package main

import "fmt"

func isSubstring(s, t string) bool {
	sPoint := 0

	for tPoint := 0; tPoint < len(t); tPoint++ {
		if sPoint < len(s) && s[sPoint] == t[tPoint] {
			sPoint++
		}
	}

	return sPoint == len(s)
}

func main() {
	fmt.Println(isSubstring("a", "abc")) // true
	fmt.Println(isSubstring("abc", "asdfdbsdfsdc")) // true
	fmt.Println(isSubstring("acb", "asdfsdbsdfdsc")) // false
}