package main

import "fmt"

func findPeakIndex(num []int) int {
  _, idx := findMax(num, 0, len(num) - 1)
  return idx
}

func findMax(nums []int, left int, right int) (int, int) {

  // base case 
  if left == right {
    return nums[left], left
  }

  // divide
  mid := (left + right)/2
  leftMax, leftIdx := findMax(nums, left, mid) // recurse left
  rightMax, rightIdx := findMax(nums, mid+1, right) // recurse right

  // combine
  if leftMax > rightMax {
    return leftMax, leftIdx
  } else {
    return rightMax, rightIdx
  }
}


func main() {
  fmt.Println(findPeakIndex([]int{-2, 1, -3, 4, 0}))
}

