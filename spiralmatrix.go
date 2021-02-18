package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
  top, left := 0, 0
  total := len(matrix[0])*len(matrix)
  right := len(matrix[0]) - 1
  bottom := len(matrix) - 1
  arr := []int{}
  j := 0

  for j < total {

  	for i := left; i <= right; i++ {
  		arr = append(arr, matrix[top][i])
  		j++
  	}

  	top++

  	for i := top; i <= bottom; i++ {
  		arr = append(arr, matrix[i][right])
  		j++
  	}

  	right--

  	for i := right; i >= left; i-- {
  		arr = append(arr, matrix[bottom][i])
  		j++
  	}

  	bottom--

  	for i := bottom; i >= top; i-- {
  		arr = append(arr, matrix[i][left])
  		j++
  	}

  	left++
  }
  
    return arr[:total]
}

func main() {
	a := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12}}

	fmt.Println(spiralOrder(a))
}
