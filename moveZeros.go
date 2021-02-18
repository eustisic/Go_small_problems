package main

import "fmt"

// func moveZeros(arr []int) []int{
// 	zeros := []int{}
// 	notZeros := []int{}

// 	for _, n := range arr {
// 		if n == 0 {
// 			zeros = append(zeros, n)
// 		} else {
// 			notZeros = append(notZeros, n)
// 		}
// 	}

// 	return append(notZeros, zeros...)
// }

func moveZeros(arr []int) []int {
	reader, writer := 0, 0

	for reader < len(arr) {
		if arr[reader] != 0 {
			arr[writer], arr[reader] = arr[reader], arr[writer]
			writer++
		}
		reader++
	}
	return arr 
}

func main() {
	fmt.Println(moveZeros([]int{0, 1, 2, 3, 0, 12}))
}

// if a zero is found swap index and index + 1 until the end is reached