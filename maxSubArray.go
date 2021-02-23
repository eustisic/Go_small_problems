package main

import "fmt"

func findMaxSubArray(nums []int, left int, right int) int {

	// base case 
	if left == right {
		return nums[left]
	}

	// divide
	mid := (left + right)/2
	leftMax := findMaxSubArray(nums, left, mid) // recurse left
	rightMax := findMaxSubArray(nums, mid+1, right) // recurse right
	crossMax := maxCrossing(nums, left, right, mid) // recurse center

	// combine
	return max(leftMax, rightMax, crossMax) // combining
}

func maxCrossing(nums []int, left int, right int, mid int) int {
	maxR, maxL := -99999, -99999

	sum := 0

	for i := mid; i >= left; i--  {
		sum += nums[i]
		maxL = max(maxL, sum)
	}

	sum = 0

	for i := mid + 1; i <= right; i++ {
		sum += nums[i]
		maxR = max(maxR, sum)
	}

	return maxR + maxL
}

func max(v ...int) int {
	max := v[0]

	for _, val := range v {
		if val > max {
			max = val
		}
	}
	return max
}



func main() {
	fmt.Println(findMaxSubArray([]int{-2, 1, -3, 4, 6}, 0, 4))
}

