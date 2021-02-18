package main

import (
	"fmt"
	"math"
)

func LongestUniqueSubstring(str string) float64 {
	table := map[byte]bool{}
	var longest float64
	i := 0

	for j := 0; j < len(str); j++ {


		for table[str[j]] {
			table[str[i]] = false
			i++
		}

		table[str[j]] = true
		longest = math.Max(float64(j - i + 1), longest)
	}

	return longest
}

func SubWithTwoDistinct(str string) float64 {
	table := map[byte]int{}
	var longest float64
	i := 0

	for j := 0; j < len(str); j++ { // 
		table[str[j]] += 1

		for len(table) > 2 {
			if table[str[i]] == 1 {
				delete(table, str[i])
			} else {
				table[str[i]] -= 1
			}
			i++
		}

		longest = math.Max(float64(j - i + 1), longest) 
	}
	return longest
}

func main() {
	fmt.Println("Longest Unique Substring:")
	fmt.Println(LongestUniqueSubstring("abcadef")) // 6
	fmt.Println(LongestUniqueSubstring("abcabcdf")) // 5
	fmt.Println(LongestUniqueSubstring("bbbbb")) // 1
	fmt.Println(LongestUniqueSubstring("")) // 0
	fmt.Println(LongestUniqueSubstring("abc")) // 3

	fmt.Println()
	fmt.Println("Longest substring with two or less distinct characters")
	fmt.Println(SubWithTwoDistinct("eceba")) // 3

}