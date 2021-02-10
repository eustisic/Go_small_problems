package main

import (
	"fmt"
	"strings"
	"sort"
)

func uniq(array []string) []string {
	m := map[string]int
	r := []string{}

	for _, s := range array {
		m[s] = 0
	}

	for k, _ := range m {
		r = append(r, k)
	}

	return r
}

func InArray(arr1 []string, arr2 []string) []string {
	r := []string{}

	for _, s1 := range arr1 {
		for _, s2 := range arr2 {
			if strings.Contains(s2, s1) {
				r = append(r, s1)
				break
			}
		}
	}

	sort.Strings(r)
	return uniq(r)
}

func main() {
	var a1 = []string{"live", "arp", "strong"}
  var a2 = []string{"lively", "alive", "harp", "sharp", "armstrong"}
  fmt.Println(InArray(a1, a2))
}