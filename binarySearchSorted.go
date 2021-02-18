package main

import "fmt"

func bsearch(arr []int, val int) bool {
	if len(arr) == 0 {
		return false
	}

	fmt.Println(arr)
	mid := len(arr) / 2
	fmt.Println(mid)

	if arr[mid] == val {
		return true
	} else if arr[mid] > val {
		return bsearch(arr[:mid], val)
	} else {
		return bsearch(arr[mid:], val)
	}
}

func main() {
	fmt.Println(bsearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5))
}